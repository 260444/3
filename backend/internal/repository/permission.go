package repository

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

// PermissionRepository 权限数据访问层
type PermissionRepository struct {
	DB *gorm.DB
}

// NewPermissionRepository 创建权限仓库
func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}

// Create 创建权限
func (r *PermissionRepository) Create(permission *model.Permission) error {
	return r.DB.Create(permission).Error
}

// GetByID 根据ID获取权限
func (r *PermissionRepository) GetByID(id uint) (*model.Permission, error) {
	var permission model.Permission
	err := r.DB.First(&permission, id).Error
	return &permission, err
}

// Update 更新权限
func (r *PermissionRepository) Update(permission *model.Permission) error {
	return r.DB.Save(permission).Error
}

// Delete 删除权限
func (r *PermissionRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Permission{}, id).Error
}

// List 获取权限列表
func (r *PermissionRepository) List(limit, offset int) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.DB.Offset(offset).Limit(limit).Find(&permissions).Error
	return permissions, err
}

// GetTotal 获取权限总数
func (r *PermissionRepository) GetTotal() (int64, error) {
	var count int64
	err := r.DB.Model(&model.Permission{}).Count(&count).Error
	return count, err
}

// GetByPathAndMethod 根据路径和方法获取权限
func (r *PermissionRepository) GetByPathAndMethod(path, method string) (*model.Permission, error) {
	var permission model.Permission
	err := r.DB.Where("path = ? AND method = ?", path, method).First(&permission).Error
	return &permission, err
}

// UpdateStatus 更新权限状态
func (r *PermissionRepository) UpdateStatus(id uint, status int8) error {
	return r.DB.Model(&model.Permission{}).Where("id = ?", id).Update("status", status).Error
}
