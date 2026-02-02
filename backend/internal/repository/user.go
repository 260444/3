package repository

import (
	"backend/internal/model"
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
func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

// GetByID 根据ID获取用户 *
func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

// GetByUsername 根据用户名获取用户 *
func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// GetPermissionsByUsername 根据用户名获取用户和权限信息
func (r *UserRepository) UserWithRoleInfo(username string) (*model.UserWithRoleInfo, error) {
	var userWithRole model.UserWithRoleInfo
	err := r.DB.Table("users").
		Select("users.*, roles.ident as ident").
		Joins("INNER JOIN roles ON users.role_id = roles.id").
		Where("users.username = ?", username).
		Scan(&userWithRole).Error

	fmt.Println("userWithRole:", userWithRole.RoleIdent)
	return &userWithRole, err
}

// GetByEmail 根据邮箱获取用户
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Update 更新用户 *
func (r *UserRepository) Update(user *model.User) error {
	return r.DB.Save(user).Error
}

// Delete 删除用户 *
func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&model.User{}, id).Error
}

// List 获取用户列表 *
func (r *UserRepository) List(limit, offset int) ([]model.User, error) {
	var users []model.User
	err := r.DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

// GetTotal 获取用户总数 *
func (r *UserRepository) GetTotal() (int64, error) {
	var count int64
	err := r.DB.Model(&model.User{}).Count(&count).Error
	return count, err
}

// UpdatePassword 更新密码
func (r *UserRepository) UpdatePassword(id uint, password string) error {
	return r.DB.Model(&model.User{}).Where("id = ?", id).Update("password", password).Error
}

// UpdateStatus 更新用户状态 *
func (r *UserRepository) UpdateStatus(id uint, status int) error {
	return r.DB.Model(&model.User{}).Where("id = ?", id).Update("status", status).Error
}
