package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/casbin"
	"fmt"
)

// PermissionService 权限管理服务
type PermissionService struct {
	RoleRepo       *repository.RoleRepository
	MenuRepo       *repository.MenuRepository
	PermissionRepo *repository.PermissionRepository
}

// NewPermissionService 创建权限管理服务
func NewPermissionService(roleRepo *repository.RoleRepository, menuRepo *repository.MenuRepository, permissionRepo *repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		RoleRepo:       roleRepo,
		MenuRepo:       menuRepo,
		PermissionRepo: permissionRepo,
	}
}

// CreatePermission 创建权限
func (s *PermissionService) CreatePermission(permission *model.Permission) error {
	return s.PermissionRepo.Create(permission)
}

// GetPermissionByID 根据ID获取权限
func (s *PermissionService) GetPermissionByID(id uint) (*model.Permission, error) {
	return s.PermissionRepo.GetByID(id)
}

// GetPermissions 获取权限列表
func (s *PermissionService) GetPermissions(limit, offset int, path, method string) ([]model.Permission, error) {
	return s.PermissionRepo.List(limit, offset, path, method)
}

// GetPermissionTotal 获取权限总数
func (s *PermissionService) GetPermissionTotal(path, method string) (int64, error) {
	return s.PermissionRepo.GetTotal(path, method)
}

// UpdatePermission 更新权限
func (s *PermissionService) UpdatePermission(permission *model.Permission) error {
	return s.PermissionRepo.Update(permission)
}

// UpdatePermissionStatus 更新权限状态
func (s *PermissionService) UpdatePermissionStatus(id uint, status int8) error {
	return s.PermissionRepo.UpdateStatus(id, status)
}

// GetAllPermissions 获取所有权限
func (s *PermissionService) GetAllPermissions(path, method string) ([]model.Permission, error) {
	return s.PermissionRepo.GetAll(path, method)
}

// DeletePermission 删除权限
func (s *PermissionService) DeletePermission(id uint) error {
	return s.PermissionRepo.Delete(id)
}

// // GetPermissionByPathAndMethod 根据路径和方法获取权限
// func (s *PermissionService) GetPermissionByPathAndMethod(path, method string) (*model.Permission, error) {
// 	return s.PermissionRepo.GetByPathAndMethod(path, method)
// }

// GetRoleMenus 获取角色的菜单权限
// func (s *PermissionService) GetRoleMenus(roleID uint) ([]model.Menu, error) {
// 	role, err := s.RoleRepo.GetByID(roleID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return role.Menus, nil
// }

// AddPolicy 添加Casbin策略
func (s *PermissionService) AddPolicy(roleID uint, path, method string) error {
	ident, err := s.RoleRepo.GetIdent(roleID)
	if err != nil {
		return err
	}
	if ident == "" {
		return fmt.Errorf("role identifier is missing")
	}
	sub := ident
	policy, err := casbin.Enforcer.AddPolicy(sub, path, method)
	fmt.Println("policy:", policy)
	if err != nil {
		return err
	}
	return nil
}

// RemovePolicy 移除Casbin策略
func (s *PermissionService) RemovePolicy(roleID uint, path, method string) error {
	ident, err := s.RoleRepo.GetIdent(roleID)
	if err != nil {
		return err
	}
	if ident == "" {
		return fmt.Errorf("role identifier is missing")
	}
	sub := ident
	policy, err := casbin.Enforcer.RemovePolicy(sub, path, method)
	fmt.Println("policy:", policy)
	if err != nil {
		return err
	}
	return nil
}

// GetPolicies 获取角色的所有Casbin策略
func (s *PermissionService) GetPolicies(roleID uint) ([][]string, error) {
	ident, err := s.RoleRepo.GetIdent(roleID)
	if err != nil {
		return nil, err
	}
	sub := ident
	return casbin.Enforcer.GetFilteredPolicy(0, sub)
}

// GetAllPolicies 获取所有Casbin策略
func (s *PermissionService) GetAllPolicies() ([][]string, error) {
	return casbin.Enforcer.GetPolicy()
}
