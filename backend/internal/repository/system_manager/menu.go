package system_manager

import (
	"backend/internal/model/system_manager"

	"gorm.io/gorm"
)

// MenuRepository 菜单数据访问层
type MenuRepository struct {
	DB *gorm.DB
}

// NewMenuRepository 创建菜单仓库
func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{DB: db}
}

// Create 创建菜单
func (r *MenuRepository) Create(menu *system_manager.Menu) error {
	return r.DB.Create(menu).Error
}

// GetByID 根据ID获取菜单
// func (r *MenuRepository) GetByID(id uint) (*system_manager.Menu, error) {
// 	var menu system_manager.Menu
// 	err := r.DB.Preload("Children").First(&menu, id).Error
// 	return &menu, err
// }

// GetAll 递归加载所有子菜单（用于菜单管理）
func (r *MenuRepository) GetAll() ([]system_manager.Menu, error) {
	var allMenus []system_manager.Menu
	// 获取所有菜单
	err := r.DB.Where("parent_id=?", 0).Find(&allMenus).Error
	if err != nil {
		return nil, err
	}

	for i := range allMenus {
		var children []system_manager.Menu
		r.DB.Where("parent_id=?", allMenus[i].ID).Order("sort ASC").Find(&children)
		allMenus[i].Children = children

	}
	return allMenus, nil

}

// Update 更新菜单
func (r *MenuRepository) Update(menu *system_manager.Menu) error {
	return r.DB.Save(menu).Error
}

// Delete 删除菜单
func (r *MenuRepository) Delete(id uint) error {
	return r.DB.Delete(&system_manager.Menu{}, id).Error
}

// GetByUserID 根据用户ID获取菜单（通过角色关联）
func (r *MenuRepository) GetByUserID(userID uint) ([]system_manager.Menu, error) {
	var menus []system_manager.Menu

	// 获取用户有权限的菜单ID
	var menuIDs []uint
	err := r.DB.Table("role_menus").
		Select("menu_id").
		Where("role_id IN (SELECT role_id FROM users WHERE id = ?)", userID).
		Find(&menuIDs).Error

	if err != nil {
		return nil, err
	}

	if len(menuIDs) == 0 {
		return []system_manager.Menu{}, nil
	}

	// 获取父级菜单（parent_id = 0 且在用户有权限的菜单列表中）
	err = r.DB.Where("parent_id = ? AND status = ? AND id IN ?", 0, 1, menuIDs).
		Order("sort ASC").
		Find(&menus).Error

	if err != nil {
		return nil, err
	}

	//为每个父菜单加载子菜单
	for i, menu := range menus {
		var c []system_manager.Menu
		r.DB.Where("parent_id = ? AND status = ? AND id IN ?", menus[i].ID, 1, menuIDs).
			Order("sort ASC").
			Find(&c)
		menu.Children = c
		menus[i] = menu
	}

	return menus, nil
}
