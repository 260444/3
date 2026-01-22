package repository

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

// RoleMenuRepository 规则与菜单关系仓库
type RoleMenuRepository struct {
	DB *gorm.DB
}

// NewRoleMenuRepository 规则与菜单关系仓库
func NewRoleMenuRepository(db *gorm.DB) *RoleMenuRepository {
	return &RoleMenuRepository{DB: db}
}

// CreateRoleMenus 创建记录RoleMenu
func (r *RoleMenuRepository) CreateRoleMenus(roleMeans []model.RoleMenu) error {
	return r.DB.Create(&roleMeans).Error
}

// GetRoleMenus 获取记录RoleMenu
//func (r *RoleMenuRepository) GetRoleMenus(roleId uint) (roleMeans []*model.RoleMenuRequest, err error) {
//	err = r.DB.Raw(`
//		SELECT a.role_id, a.menu_id, b.title
//		FROM role_menus AS a
//		LEFT JOIN menus AS b ON a.menu_id = b.id
//		WHERE a.role_id = ?
//    `).Scan(&roleMeans).Error
//	return roleMeans, err
//}

// DeleteRoleMenus 删除记录RoleMenu
func (r *RoleMenuRepository) DeleteRoleMenus(roleId uint, roleMeans []uint) error {
	return r.DB.Where("role_id = ? AND menu_id IN ?", roleId, roleMeans).Delete(&model.RoleMenu{}).Error
}

// GetRoleMenuByID 根据ID获取记录RoleMenu
func (r *RoleMenuRepository) GetRoleMenuByID(roleId uint) (roleMeans []*model.RoleMenuRequest, err error) {
	err = r.DB.Raw(`
		SELECT a.role_id, a.menu_id, b.title 
		FROM role_menus AS a 
		LEFT JOIN menus AS b ON a.menu_id = b.id 
		WHERE a.role_id = ?
    `, roleId).Scan(&roleMeans).Error
	return roleMeans, err
}
