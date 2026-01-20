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
func (s *PermissionService) GetPermissions(limit, offset int) ([]model.Permission, error) {
	return s.PermissionRepo.List(limit, offset)
}

// GetPermissionTotal 获取权限总数
func (s *PermissionService) GetPermissionTotal() (int64, error) {
	return s.PermissionRepo.GetTotal()
}

// UpdatePermission 更新权限
func (s *PermissionService) UpdatePermission(permission *model.Permission) error {
	return s.PermissionRepo.Update(permission)
}

// UpdatePermissionStatus 更新权限状态
func (s *PermissionService) UpdatePermissionStatus(id uint, status int8) error {
	return s.PermissionRepo.UpdateStatus(id, status)
}

// DeletePermission 删除权限
func (s *PermissionService) DeletePermission(id uint) error {
	return s.PermissionRepo.Delete(id)
}

// GetPermissionByPathAndMethod 根据路径和方法获取权限
func (s *PermissionService) GetPermissionByPathAndMethod(path, method string) (*model.Permission, error) {
	return s.PermissionRepo.GetByPathAndMethod(path, method)
}

// AssignMenuToRole 为角色分配菜单权限
func (s *PermissionService) AssignMenuToRole(roleID uint, menuIDs []uint) error {
	role, err := s.RoleRepo.GetByID(roleID)
	if err != nil {
		return err
	}

	// 获取所有菜单
	var menus []model.Menu
	for _, menuID := range menuIDs {
		menu, err := s.MenuRepo.GetByID(menuID)
		if err != nil {
			return err
		}
		menus = append(menus, *menu)
	}

	// 更新角色的菜单关联
	role.Menus = menus
	return s.RoleRepo.Update(role)
}

// GetRoleMenus 获取角色的菜单权限
func (s *PermissionService) GetRoleMenus(roleID uint) ([]model.Menu, error) {
	role, err := s.RoleRepo.GetByID(roleID)
	if err != nil {
		return nil, err
	}
	return role.Menus, nil
}

// GetUserMenus 获取用户的菜单权限（通过角色）
func (s *PermissionService) GetUserMenus(userID uint) ([]model.Menu, error) {
	// 这里需要从UserService获取用户信息，暂时简化处理
	// 实际应用中应该通过UserRepo获取用户，然后获取用户的角色，再获取角色的菜单
	return nil, nil
}

// RemoveMenuFromRole 移除角色的菜单权限
func (s *PermissionService) RemoveMenuFromRole(roleID uint, menuIDs []uint) error {
	role, err := s.RoleRepo.GetByID(roleID)
	if err != nil {
		return err
	}

	// 过滤掉要移除的菜单
	var remainingMenus []model.Menu
	menuIDMap := make(map[uint]bool)
	for _, menuID := range menuIDs {
		menuIDMap[menuID] = true
	}

	for _, menu := range role.Menus {
		if !menuIDMap[menu.ID] {
			remainingMenus = append(remainingMenus, menu)
		}
	}

	role.Menus = remainingMenus
	return s.RoleRepo.Update(role)
}

// AddPolicy 添加Casbin策略
func (s *PermissionService) AddPolicy(roleID uint, path, method string) error {
	sub := fmt.Sprintf("role_%d", roleID)
	policy, err := casbin.Enforcer.AddPolicy(sub, path, method)
	fmt.Println("policy:", policy)
	if err != nil {
		return err
	}
	return nil
}

// RemovePolicy 移除Casbin策略
func (s *PermissionService) RemovePolicy(roleID uint, path, method string) error {
	sub := fmt.Sprintf("role_%d", roleID)
	policy, err := casbin.Enforcer.RemovePolicy(sub, path, method)
	fmt.Println("policy:", policy)
	if err != nil {
		return err
	}
	return nil
}

// GetPolicies 获取角色的所有Casbin策略
func (s *PermissionService) GetPolicies(roleID uint) ([][]string, error) {
	sub := fmt.Sprintf("role_%d", roleID)
	return casbin.Enforcer.GetFilteredPolicy(0, sub)
}

// GetAllPolicies 获取所有Casbin策略
func (s *PermissionService) GetAllPolicies() ([][]string, error) {
	return casbin.Enforcer.GetPolicy()
}
