package asset_management

import (
	assModel "backend/internal/model/asset_management"
	"gorm.io/gorm"
)

type HostGroupRepository struct {
	DB *gorm.DB
}

func NewHostGroupRepository(db *gorm.DB) *HostGroupRepository {
	return &HostGroupRepository{DB: db}
}

// Create 创建主机组
func (r *HostGroupRepository) Create(group *assModel.HostGroup) error {
	return r.DB.Create(group).Error
}

// GetByID 根据ID获取主机组
func (r *HostGroupRepository) GetByID(id uint) (*assModel.HostGroup, error) {
	var group assModel.HostGroup
	err := r.DB.Preload("Hosts").Where("id = ?", id).First(&group).Error
	return &group, err
}

// List 获取主机组列表
func (r *HostGroupRepository) List(page, pageSize int, name string, status *int8) ([]assModel.HostGroup, int64, error) {
	var groups []assModel.HostGroup
	var total int64

	query := r.DB.Model(&assModel.HostGroup{})

	// 添加查询条件
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&groups).Error; err != nil {
		return nil, 0, err
	}

	// 为每个主机组添加主机数量统计
	for i := range groups {
		var hostCount int64
		r.DB.Model(&assModel.Host{}).Where("group_id = ?", groups[i].ID).Count(&hostCount)
		// 这里可以将hostCount存储到某个字段中，或者在service层处理
	}

	return groups, total, nil
}

// Update 更新主机组
func (r *HostGroupRepository) Update(group *assModel.HostGroup) error {
	return r.DB.Save(group).Error
}

// Delete 删除主机组（软删除）
func (r *HostGroupRepository) Delete(id uint) error {
	return r.DB.Delete(&assModel.HostGroup{}, id).Error
}

// UpdateStatus 更新主机组状态
func (r *HostGroupRepository) UpdateStatus(id uint, status int8) error {
	return r.DB.Model(&assModel.HostGroup{}).Where("id = ?", id).Update("status", status).Error
}

// GetByName 根据名称获取主机组
func (r *HostGroupRepository) GetByName(name string) (*assModel.HostGroup, error) {
	var group assModel.HostGroup
	err := r.DB.Where("name = ?", name).First(&group).Error
	return &group, err
}

// GetGroupWithHosts 获取主机组及其关联的主机信息
func (r *HostGroupRepository) GetGroupWithHosts(id uint) (*assModel.HostGroup, error) {
	var group assModel.HostGroup
	err := r.DB.Preload("Hosts").Where("id = ?", id).First(&group).Error
	return &group, err
}

// CheckGroupHasHosts 检查主机组是否有关联的主机
func (r *HostGroupRepository) CheckGroupHasHosts(id uint) (bool, error) {
	var count int64
	err := r.DB.Model(&assModel.Host{}).Where("group_id = ?", id).Count(&count).Error
	return count > 0, err
}
