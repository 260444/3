package asset_management

import (
	assModel "backend/internal/model/asset_management"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type CredentialRepository struct {
	DB *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) *CredentialRepository {
	return &CredentialRepository{DB: db}
}

// CreateCredential 创建凭据
func (s *CredentialRepository) CreateCredential(credential *assModel.Credential) (err error) {

	if err := s.DB.Create(credential).Error; err != nil {
		return errors.New("凭据创建失败")
	}
	return err
}

// GetCredentialByID 根据ID获取凭据
func (s *CredentialRepository) GetCredentialByID(id uint) (*assModel.Credential, error) {
	var credential assModel.Credential
	if err := s.DB.Preload("Hosts").First(&credential, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("凭据不存在")
		}
		return nil, err
	}
	return &credential, nil
}

// ListCredentials 分页获取凭据列表
func (s *CredentialRepository) ListCredentials(page, pageSize int, name, username string) ([]assModel.Credential, int64, error) {
	var credentials []assModel.Credential
	var total int64

	db := s.DB.Model(&assModel.Credential{})

	// 添加筛选条件
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	if username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", username))
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&credentials).Error; err != nil {
		return nil, 0, err
	}

	return credentials, total, nil
}

// UpdateCredential 更新凭据
func (s *CredentialRepository) UpdateCredential(credential *assModel.Credential) error {
	if err := s.DB.Save(credential).Error; err != nil {
		return errors.New("凭据更新失败")
	}
	return nil
}

// DeleteCredential 删除凭据
func (s *CredentialRepository) DeleteCredential(credential *assModel.Credential) error {
	// 开始事务
	tx := s.DB.Begin()
	defer func() {
		if r := tx.Rollback(); r != nil {
			fmt.Printf("事务回滚错误: %v\n", r)
		}
	}()

	// 从关联表中删除凭据与主机的关联
	if err := tx.Exec("DELETE FROM host_credentials WHERE credential_id = ?", credential.ID).Error; err != nil {
		tx.Rollback()
		return errors.New("删除凭据与主机的关联失败")
	}

	// 逻辑删除凭据
	if err := tx.Delete(credential).Error; err != nil {
		tx.Rollback()
		return errors.New("凭据删除失败")
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("操作提交失败")
	}

	return nil
}

// BatchDeleteCredentials 批量删除凭据
func (s *CredentialRepository) BatchDeleteCredentials(ids []uint) (int64, error) {
	// 开始事务
	tx := s.DB.Begin()
	defer func() {
		if r := tx.Rollback(); r != nil {
			fmt.Printf("事务回滚错误: %v\n", r)
		}
	}()

	// 从关联表中删除凭据与主机的关联
	if err := tx.Exec("DELETE FROM host_credentials WHERE credential_id IN ?", ids).Error; err != nil {
		tx.Rollback()
		return 0, errors.New("删除凭据与主机的关联失败")
	}

	// 逻辑删除凭据
	result := tx.Delete(&assModel.Credential{}, "id IN ?", ids)
	if result.Error != nil {
		tx.Rollback()
		return 0, errors.New("凭据批量删除失败")
	}

	if err := tx.Commit().Error; err != nil {
		return 0, errors.New("操作提交失败")
	}

	return result.RowsAffected, nil
}

// GetCredentialsByHost 获取主机关联的凭据
func (s *CredentialRepository) GetCredentialsByHost(hostID uint) ([]assModel.Credential, error) {
	var credentials []assModel.Credential
	if err := s.DB.
		Joins("JOIN host_credentials ON credentials.id = host_credentials.credential_id").
		Where("host_credentials.host_id = ?", hostID).
		Find(&credentials).Error; err != nil {
		return nil, errors.New("获取主机凭据失败")
	}
	return credentials, nil
}
