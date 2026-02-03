package system_manager

import (
	"errors"
)

// RoleService 角色服务
type RoleService struct {
	RoleRepo *sysRepository.RoleRepository
}

// NewRoleService 创建角色服务
func NewRoleService(roleRepo *sysRepository.RoleRepository) *RoleService {
	return &RoleService{
		RoleRepo: roleRepo,
	}
}

// CreateRole 创建角色
func (s *RoleService) CreateRole(role *sysModel.Role) error {
	_, err := s.RoleRepo.GetByName(role.Name)
	if err == nil {
		return errors.New("角色名称已存在")
	}

	return s.RoleRepo.Create(role)
}

// GetRoleByID 根据ID获取角色
func (s *RoleService) GetRoleByID(id uint) (*sysModel.Role, error) {
	return s.RoleRepo.GetByID(id)
}

// GetRoleByName 根据名称获取角色
func (s *RoleService) GetRoleByName(name string) (*sysModel.Role, error) {
	return s.RoleRepo.GetByName(name)
}

// UpdateRole 更新角色
func (s *RoleService) UpdateRole(role *sysModel.Role) error {
	return s.RoleRepo.Update(role)
}

// DeleteRole 删除角色
func (s *RoleService) DeleteRole(id uint) error {
	return s.RoleRepo.Delete(id)
}

// GetRoles 获取角色列表
func (s *RoleService) GetRoles(limit, offset int) ([]sysModel.Role, int64, error) {
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
