package system_manager

import (
	"backend/internal/model/system_manager"

	"gorm.io/gorm"
)

// RoleRepository 角色数据访问层
type RoleRepository struct {
	DB *gorm.DB
}

// NewRoleRepository 创建角色仓库
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

// Create 创建角色
func (r *RoleRepository) Create(role *system_manager.Role) error {
	return r.DB.Create(role).Error
}

// GetByID 根据ID获取角色
func (r *RoleRepository) GetByID(id uint) (*system_manager.Role, error) {
	var role system_manager.Role
	err := r.DB.Preload("Menus").First(&role, id).Error
	return &role, err
}

// GetByName 根据名称获取角色
func (r *RoleRepository) GetByName(name string) (*system_manager.Role, error) {
	var role system_manager.Role
	err := r.DB.Where("name = ?", name).First(&role).Error
	return &role, err
}

// Update 更新角色
func (r *RoleRepository) Update(role *system_manager.Role) error {
	return r.DB.Save(role).Error
}

// Delete 删除角色
func (r *RoleRepository) Delete(id uint) error {
	return r.DB.Delete(&system_manager.Role{}, id).Error
}

// List 获取角色列表
func (r *RoleRepository) List(limit, offset int) ([]system_manager.Role, error) {
	var roles []system_manager.Role
	err := r.DB.Offset(offset).Limit(limit).Find(&roles).Error
	return roles, err
}

// GetTotal 获取角色总数
func (r *RoleRepository) GetTotal() (int64, error) {
	var count int64
	err := r.DB.Model(&system_manager.Role{}).Count(&count).Error
	return count, err
}

// GetIdent 根据ID获取角色标识
func (r *RoleRepository) GetIdent(id uint) (string, error) {
	var role system_manager.Role
	err := r.DB.Select("ident").First(&role, id).Error
	return role.Ident, err
}

// GetByIdent 根据Ident获取角色
func (r *RoleRepository) GetByIdent(ident string) (*system_manager.Role, error) {
	var role system_manager.Role
	err := r.DB.Where("ident = ?", ident).First(&role).Error
	return &role, err
}
