package system_manager

import (
	"backend/internal/model/system_manager"
	repository "backend/internal/repository/system_manager"
)

// MenuService 菜单服务
type MenuService struct {
	MenuRepo *repository.MenuRepository
}

// NewMenuService 创建菜单服务
func NewMenuService(menuRepo *repository.MenuRepository) *MenuService {
	return &MenuService{
		MenuRepo: menuRepo,
	}
}

// CreateMenu 创建菜单
func (s *MenuService) CreateMenu(menu *system_manager.Menu) error {
	return s.MenuRepo.Create(menu)
}

// GetMenuByID 根据ID获取菜单
// func (s *MenuService) GetMenuByID(id uint) (*system_manager.Menu, error) {
// 	return s.MenuRepo.GetByID(id)
// }

// GetMenuTree 获取菜单树
//func (s *MenuService) GetMenuTree(parentID *uint) ([]system_manager.Menu, error) {
//	return s.MenuRepo.GetByParentID(parentID)
//}

// GetAllMenus 获取所有菜单
func (s *MenuService) GetAllMenus() ([]system_manager.Menu, error) {
	return s.MenuRepo.GetAll()
}

// UpdateMenu 更新菜单
func (s *MenuService) UpdateMenu(menu *system_manager.Menu) error {
	return s.MenuRepo.Update(menu)
}

// DeleteMenu 删除菜单
func (s *MenuService) DeleteMenu(id uint) error {
	return s.MenuRepo.Delete(id)
}

// GetUserMenus 获取当前用户的菜单
func (s *MenuService) GetUserMenus(userID uint) ([]system_manager.Menu, error) {
	return s.MenuRepo.GetByUserID(userID)
}
