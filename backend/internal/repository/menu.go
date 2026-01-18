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
func (r *MenuRepository) GetByParentID(parentID *uint) ([]model.Menu, error) {
	var menus []model.Menu
	query := r.DB.Where("status = ?", 1)
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}
	err := query.Order("sort ASC").Find(&menus).Error
	if err != nil {
		return nil, err
	}

	// 递归加载子菜单
	for i := range menus {
		r.loadAllChildren(&menus[i])
	}

	return menus, nil
}

// loadAllChildren 递归加载所有子菜单（用于菜单管理）
func (r *MenuRepository) loadAllChildren(menu *model.Menu) error {
	var children []model.Menu
	err := r.DB.Where("parent_id = ? AND status = ?", menu.ID, 1).
		Order("sort ASC").
		Find(&children).Error
	if err != nil {
		return err
	}

	menu.Children = children

	// 递归加载子菜单的子菜单
	for i := range children {
		r.loadAllChildren(&children[i])
	}

	return nil
}

// GetAll 获取所有菜单
func (r *MenuRepository) GetAll() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.DB.Preload("Children").Find(&menus).Error
	return menus, err
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
	
	// 通过用户 -> 角色 -> 菜单的关联查询
	err := r.DB.Table("menus").
		Select("menus.*, menus.id as id, menus.name as name, menus.title as title, menus.path as path, menus.component as component, menus.redirect as redirect, menus.parent_id as parent_id, menus.icon as icon, menus.sort as sort, menus.is_hidden as is_hidden, menus.is_link as is_link, menus.link_url as link_url, menus.status as status, menus.created_at as created_at, menus.updated_at as updated_at").
		Joins("INNER JOIN role_menus ON menus.id = role_menus.menu_id").
		Joins("INNER JOIN roles ON role_menus.role_id = roles.id").
		Joins("INNER JOIN users ON users.role_id = roles.id").
		Where("users.id = ? AND menus.status = 1 AND roles.status = 1", userID).
		Where("menus.parent_id IS NULL").
		Order("menus.sort ASC").
		Find(&menus).Error
	
	if err != nil {
		return nil, err
	}

	// 递归加载子菜单
	for i := range menus {
		r.loadChildren(&menus[i], userID)
	}

	return menus, nil
}

// loadChildren 递归加载子菜单
func (r *MenuRepository) loadChildren(menu *model.Menu, userID uint) error {
	var children []model.Menu
	
	err := r.DB.Table("menus").
		Select("menus.*, menus.id as id, menus.name as name, menus.title as title, menus.path as path, menus.component as component, menus.redirect as redirect, menus.parent_id as parent_id, menus.icon as icon, menus.sort as sort, menus.is_hidden as is_hidden, menus.is_link as is_link, menus.link_url as link_url, menus.status as status, menus.created_at as created_at, menus.updated_at as updated_at").
		Joins("INNER JOIN role_menus ON menus.id = role_menus.menu_id").
		Joins("INNER JOIN roles ON role_menus.role_id = roles.id").
		Joins("INNER JOIN users ON users.role_id = roles.id").
		Where("users.id = ? AND menus.status = 1 AND roles.status = 1", userID).
		Where("menus.parent_id = ?", menu.ID).
		Order("menus.sort ASC").
		Find(&children).Error
	
	if err != nil {
		return err
	}

	menu.Children = children

	// 递归加载子菜单的子菜单
	for i := range children {
		r.loadChildren(&children[i], userID)
	}

	return nil
}