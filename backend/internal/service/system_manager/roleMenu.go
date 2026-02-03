package system_manager

import (
	"backend/internal/model/system_manager"
	repository "backend/internal/repository/system_manager"
)

// RoleMenuService 角色服务
type RoleMenuService struct {
	RoleMenuRepo *repository.RoleMenuRepository
}

// NewRoleMenuService 创建角色服务
func NewRoleMenuService(RoleMenuRepo *repository.RoleMenuRepository) *RoleMenuService {
	return &RoleMenuService{
		RoleMenuRepo: RoleMenuRepo,
	}
}

// AssignMenuToRole 为角色分配菜单权限
func (s *RoleMenuService) AssignMenuToRole(roleID uint, menuIDs []uint) error {
	//TODO 判断roleID是否存在
	var roleMeans []system_manager.RoleMenu
	for _, d := range menuIDs {
		roleMeans = append(roleMeans, system_manager.RoleMenu{
			RoleID: roleID,
			MenuID: d,
		})
	}
	return s.RoleMenuRepo.CreateRoleMenus(roleMeans)
}

// GetUserMenusByID 获取用户的菜单权限（通过角色）
func (s *RoleMenuService) GetUserMenusByID(roleID uint) ([]*system_manager.RoleMenuRequest, error) {
	//TODO 判断roleID是否存在
	return s.RoleMenuRepo.GetRoleMenuByID(roleID)
}

// RemoveMenuFromRole 移除角色的菜单权限
func (s *RoleMenuService) RemoveMenuFromRole(roleID uint, menuIDs []uint) error {
	//TODO 判断roleID是否存在
	return s.RoleMenuRepo.DeleteRoleMenus(roleID, menuIDs)
}
