package system_manager

import (
	"backend/internal/model/system_manager"
	"fmt"

	"gorm.io/gorm"
)

// UserRepository 用户数据访问层
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository 创建用户仓库
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create 创建用户 *
func (r *UserRepository) Create(user *system_manager.User) error {
	return r.DB.Create(user).Error
}

// GetByID 根据ID获取用户 *
func (r *UserRepository) GetByID(id uint) (*system_manager.User, error) {
	var user system_manager.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

// GetByUsername 根据用户名获取用户 *
func (r *UserRepository) GetByUsername(username string) (*system_manager.User, error) {
	var user system_manager.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// GetPermissionsByUsername 根据用户名获取用户和权限信息
func (r *UserRepository) UserWithRoleInfo(username string) (*system_manager.UserWithRoleInfo, error) {
	var userWithRole system_manager.UserWithRoleInfo
	err := r.DB.Table("users").
		Select("users.*, roles.ident as ident").
		Joins("INNER JOIN roles ON users.role_id = roles.id").
		Where("users.username = ?", username).
		Scan(&userWithRole).Error

	fmt.Println("userWithRole:", userWithRole.RoleIdent)
	return &userWithRole, err
}

// GetByEmail 根据邮箱获取用户
func (r *UserRepository) GetByEmail(email string) (*system_manager.User, error) {
	var user system_manager.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Update 更新用户 *
func (r *UserRepository) Update(user *system_manager.User) error {
	return r.DB.Save(user).Error
}

// Delete 删除用户 *
func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&system_manager.User{}, id).Error
}

// List 获取用户列表 *
func (r *UserRepository) List(limit, offset int) ([]system_manager.User, error) {
	var users []system_manager.User
	err := r.DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

// GetTotal 获取用户总数 *
func (r *UserRepository) GetTotal() (int64, error) {
	var count int64
	err := r.DB.Model(&system_manager.User{}).Count(&count).Error
	return count, err
}

// UpdatePassword 更新密码
func (r *UserRepository) UpdatePassword(id uint, password string) error {
	return r.DB.Model(&system_manager.User{}).Where("id = ?", id).Update("password", password).Error
}

// UpdateStatus 更新用户状态 *
func (r *UserRepository) UpdateStatus(id uint, status int) error {
	return r.DB.Model(&system_manager.User{}).Where("id = ?", id).Update("status", status).Error
}
