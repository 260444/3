package repository

import (
	"backend/internal/model"
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
func (r *MenuRepository) Create(menu *model.Menu) error {
	return r.DB.Create(menu).Error
}

// GetByID 根据ID获取菜单
func (r *MenuRepository) GetByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	err := r.DB.Preload("Children").First(&menu, id).Error
	return &menu, err
}

// GetByParentID 根据父级ID获取菜单列表
//func (r *MenuRepository) GetByParentID(parentID *uint) ([]model.Menu, error) {
//	var menus []model.Menu
//	query := r.DB.Where("status = ?", 1)
//	if parentID != nil {
//		query = query.Where("parent_id = ?", *parentID)
//	} else {
//		query = query.Where("parent_id IS NULL")
//	}
//	err := query.Order("sort ASC").Find(&menus).Error
//	if err != nil {
//		return nil, err
//	}
//
//	// 递归加载子菜单
//	for i := range menus {
//		r.loadAllChildren(&menus[i])
//	}
//
//	return menus, nil
//}

// GetAll 递归加载所有子菜单（用于菜单管理）
func (r *MenuRepository) GetAll() ([]model.Menu, error) {
	var allMenus []model.Menu
	// 获取所有菜单
	err := r.DB.Where("parent_id=?", 0).Find(&allMenus).Error
	if err != nil {
		return nil, err
	}

	for i := range allMenus {
		var children []model.Menu
		r.DB.Where("parent_id=?", allMenus[i].ID).Order("sort ASC").Find(&children)
		allMenus[i].Children = append(children)
	}
	return allMenus, nil

}

// Update 更新菜单
func (r *MenuRepository) Update(menu *model.Menu) error {
	return r.DB.Save(menu).Error
}

// Delete 删除菜单
func (r *MenuRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Menu{}, id).Error
}

// GetByUserID 根据用户ID获取菜单（通过角色关联）
func (r *MenuRepository) GetByUserID(userID uint) ([]model.Menu, error) {
	var menus []model.Menu

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
		return []model.Menu{}, nil
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
		var c []model.Menu
		r.DB.Where("parent_id = ? AND status = ? AND id IN ?", menus[i].ID, 1, menuIDs).
			Order("sort ASC").
			Find(&c)
		menu.Children = c
		menus[i] = menu
	}

	return menus, nil
}
