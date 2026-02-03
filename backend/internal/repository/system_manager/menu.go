package system_manager

import (
	sysModel "backend/internal/model/system_manager"
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
func (r *MenuRepository) Create(menu *sysModel.Menu) error {
	return r.DB.Create(menu).Error
}

// GetByID 根据ID获取菜单
// func (r *MenuRepository) GetByID(id uint) (*sysModel.Menu, error) {
// 	var menu sysModel.Menu
// 	err := r.DB.Preload("Children").First(&menu, id).Error
// 	return &menu, err
// }

// GetAll 递归加载所有子菜单（用于菜单管理）
func (r *MenuRepository) GetAll() ([]sysModel.Menu, error) {
	var allMenus []sysModel.Menu
	// 获取所有菜单
	err := r.DB.Where("parent_id=?", 0).Find(&allMenus).Error
	if err != nil {
		return nil, err
	}

	for i := range allMenus {
		var children []sysModel.Menu
		r.DB.Where("parent_id=?", allMenus[i].ID).Order("sort ASC").Find(&children)
		allMenus[i].Children = children

	}
	return allMenus, nil

}

// Update 更新菜单
func (r *MenuRepository) Update(menu *sysModel.Menu) error {
	return r.DB.Save(menu).Error
}

// Delete 删除菜单
func (r *MenuRepository) Delete(id uint) error {
	return r.DB.Delete(&sysModel.Menu{}, id).Error
}

// GetByUserID 根据用户ID获取菜单（通过角色关联）
func (r *MenuRepository) GetByUserID(userID uint) ([]sysModel.Menu, error) {
	var menus []sysModel.Menu

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
		return []sysModel.Menu{}, nil
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
		var c []sysModel.Menu
		r.DB.Where("parent_id = ? AND status = ? AND id IN ?", menus[i].ID, 1, menuIDs).
			Order("sort ASC").
			Find(&c)
		menu.Children = c
		menus[i] = menu
	}

	return menus, nil
}
