package asset_management

import (
	assModel "backend/internal/model/asset_management"
	assRepo "backend/internal/repository/asset_management"
	"errors"
)

type CredentialService struct {
	CredentialRepository *assRepo.CredentialRepository
}

func NewCredentialService(credentialRepository *assRepo.CredentialRepository) *CredentialService {
	return &CredentialService{
		CredentialRepository: credentialRepository,
	}
}

// CreateCredential 创建凭据
func (s *CredentialService) CreateCredential(req *assModel.CredentialCreateRequest, userID uint) (*assModel.Credential, error) {
	credential := &assModel.Credential{
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password, // 假设密码已加密
		Description: req.Description,
	}

	err := s.CredentialRepository.CreateCredential(credential)
	if err != nil {
		return nil, err
	}

	return credential, nil
}

// GetCredentialByID 根据ID获取凭据
func (s *CredentialService) GetCredentialByID(id uint) (*assModel.Credential, error) {
	credential, err := s.CredentialRepository.GetCredentialByID(id)
	if err != nil {
		return nil, err
	}
	return credential, nil
}

// ListCredentials 分页获取凭据列表
func (s *CredentialService) ListCredentials(page, pageSize int, name, username string) ([]assModel.Credential, int64, error) {
	credentials, total, err := s.CredentialRepository.ListCredentials(page, pageSize, name, username)
	if err != nil {
		return nil, 0, err
	}
	return credentials, total, nil
}

// UpdateCredential 更新凭据
func (s *CredentialService) UpdateCredential(id uint, req *assModel.CredentialUpdateRequest, userID uint) (*assModel.Credential, error) {
	credential := &assModel.Credential{}
	credential, err := s.CredentialRepository.GetCredentialByID(id)
	if err != nil {
		return nil, errors.New("凭据不存在")
	}

	// 更新字段
	if req.Name != "" {
		credential.Name = req.Name
	}
	if req.Username != "" {
		credential.Username = req.Username
	}
	if req.Password != "" {
		credential.Password = req.Password // 假设密码已加密
	}
	if req.Description != "" {
		credential.Description = req.Description
	}
	// credential.UpdatedBy = &userID

	err = s.CredentialRepository.UpdateCredential(credential)
	if err != nil {
		return nil, errors.New("凭据更新失败")
	}

	return credential, nil
}

// DeleteCredential 删除凭据
func (s *CredentialService) DeleteCredential(id uint) error {
	credential := &assModel.Credential{}
	credential, err := s.CredentialRepository.GetCredentialByID(id)
	if err != nil {
		return errors.New("凭据不存在")
	}

	err = s.CredentialRepository.DeleteCredential(credential)
	if err != nil {
		return err
	}
	return nil
}

// BatchDeleteCredentials 批量删除凭据
func (s *CredentialService) BatchDeleteCredentials(ids []uint) (int64, error) {
	if len(ids) == 0 {
		return 0, errors.New("ids不能为空")
	}

	result, err := s.CredentialRepository.BatchDeleteCredentials(ids)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// GetCredentialsByHost 获取主机关联的凭据
func (s *CredentialService) GetCredentialsByHost(hostID uint) ([]assModel.Credential, error) {
	credentials, err := s.CredentialRepository.GetCredentialsByHost(hostID)
	if err != nil {
		return nil, err
	}
	return credentials, nil
}
