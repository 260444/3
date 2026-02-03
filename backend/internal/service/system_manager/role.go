package system_manager

import (
	"backend/internal/model/system_manager"
	repository "backend/internal/repository/system_manager"
	"errors"
)

// RoleService 角色服务
type RoleService struct {
	RoleRepo *repository.RoleRepository
}

// NewRoleService 创建角色服务
func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{
		RoleRepo: roleRepo,
	}
}

// CreateRole 创建角色
func (s *RoleService) CreateRole(role *system_manager.Role) error {
	_, err := s.RoleRepo.GetByName(role.Name)
	if err == nil {
		return errors.New("角色名称已存在")
	}

	return s.RoleRepo.Create(role)
}

// GetRoleByID 根据ID获取角色
func (s *RoleService) GetRoleByID(id uint) (*system_manager.Role, error) {
	return s.RoleRepo.GetByID(id)
}

// GetRoleByName 根据名称获取角色
func (s *RoleService) GetRoleByName(name string) (*system_manager.Role, error) {
	return s.RoleRepo.GetByName(name)
}

// UpdateRole 更新角色
func (s *RoleService) UpdateRole(role *system_manager.Role) error {
	return s.RoleRepo.Update(role)
}

// DeleteRole 删除角色
func (s *RoleService) DeleteRole(id uint) error {
	return s.RoleRepo.Delete(id)
}

// GetRoles 获取角色列表
func (s *RoleService) GetRoles(limit, offset int) ([]system_manager.Role, int64, error) {
	roles, err := s.RoleRepo.List(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.RoleRepo.GetTotal()
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}
