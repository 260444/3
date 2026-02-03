package system_manager

import (
	sysModel "backend/internal/model/system_manager"
	"gorm.io/gorm"
)

// PermissionRepository 权限数据访问层
type PermissionRepository struct {
	DB *gorm.DB
}

// NewPermissionRepository 创建权限仓库
func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}

// Create 创建权限
func (r *PermissionRepository) Create(permission *sysModel.Permission) error {
	return r.DB.Create(permission).Error
}

// GetByID 根据ID获取权限
func (r *PermissionRepository) GetByID(id uint) (*sysModel.Permission, error) {
	var permission sysModel.Permission
	err := r.DB.First(&permission, id).Error
	return &permission, err
}

// Update 更新权限
func (r *PermissionRepository) Update(permission *sysModel.Permission) error {
	return r.DB.Save(permission).Error
}

// Delete 删除权限
func (r *PermissionRepository) Delete(id uint) error {
	return r.DB.Delete(&sysModel.Permission{}, id).Error
}

// List 获取权限列表
func (r *PermissionRepository) List(limit, offset int, path, method string) ([]sysModel.Permission, error) {
	var permissions []sysModel.Permission
	query := r.DB.Model(&sysModel.Permission{})

	if path != "" {
		query = query.Where("path LIKE ?", "%"+path+"%")
	}
	if method != "" {
		query = query.Where("method = ?", method)
	}

	err := query.Offset(offset).Limit(limit).Find(&permissions).Error
	return permissions, err
}

// GetTotal 获取权限总数
func (r *PermissionRepository) GetTotal(path, method string) (int64, error) {
	var count int64
	query := r.DB.Model(&sysModel.Permission{})

	if path != "" {
		query = query.Where("path LIKE ?", "%"+path+"%")
	}
	if method != "" {
		query = query.Where("method = ?", method)
	}

	err := query.Count(&count).Error
	return count, err
}

// GetByPathAndMethod 根据路径和方法获取权限
func (r *PermissionRepository) GetByPathAndMethod(path, method string) (*sysModel.Permission, error) {
	var permission sysModel.Permission
	err := r.DB.Where("path = ? AND method = ?", path, method).First(&permission).Error
	return &permission, err
}

// GetAll 获取所有权限
func (r *PermissionRepository) GetAll(path, method string) ([]sysModel.Permission, error) {
	var permissions []sysModel.Permission
	query := r.DB.Model(&sysModel.Permission{})

	if path != "" {
		query = query.Where("path LIKE ?", "%"+path+"%")
	}
	if method != "" {
		query = query.Where("method = ?", method)
	}

	err := query.Find(&permissions).Error
	return permissions, err
}

// UpdateStatus 更新权限状态
func (r *PermissionRepository) UpdateStatus(id uint, status int8) error {
	return r.DB.Model(&sysModel.Permission{}).Where("id = ?", id).Update("status", status).Error
}
