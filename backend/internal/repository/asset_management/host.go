package asset_management

import (
	assModel "backend/internal/model/asset_management"
	"gorm.io/gorm"
)

type HostRepository struct {
	DB *gorm.DB
}

func NewHostRepository(db *gorm.DB) *HostRepository {
	return &HostRepository{DB: db}
}

// Create 创建主机
func (r *HostRepository) Create(host *assModel.Host) error {
	return r.DB.Create(host).Error
}

// GetByID 根据ID获取主机
func (r *HostRepository) GetByID(id uint) (*assModel.Host, error) {
	var host assModel.Host
	err := r.DB.Preload("Group").Where("id = ?", id).First(&host).Error
	return &host, err
}

// List 获取主机列表
func (r *HostRepository) List(page, pageSize int, hostname, ipAddress string, groupID *uint, status *int8, osType string) ([]assModel.Host, int64, error) {
	var hosts []assModel.Host
	var total int64

	query := r.DB.Model(&assModel.Host{}).Preload("Group")

	// 添加查询条件
	if hostname != "" {
		query = query.Where("hostname LIKE ?", "%"+hostname+"%")
	}
	if ipAddress != "" {
		query = query.Where("ip_address LIKE ?", "%"+ipAddress+"%")
	}
	if groupID != nil {
		query = query.Where("group_id = ?", *groupID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if osType != "" {
		query = query.Where("os_type = ?", osType)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&hosts).Error; err != nil {
		return nil, 0, err
	}

	return hosts, total, nil
}

// Update 更新主机
func (r *HostRepository) Update(host *assModel.Host) error {
	return r.DB.Save(host).Error
}

// Delete 删除主机（软删除）
func (r *HostRepository) Delete(id uint) error {
	return r.DB.Delete(&assModel.Host{}, id).Error
}

// BatchDelete 批量删除主机
func (r *HostRepository) BatchDelete(ids []uint) (int64, error) {
	result := r.DB.Where("id IN ?", ids).Delete(&assModel.Host{})
	return result.RowsAffected, result.Error
}

// UpdateStatus 更新主机状态
func (r *HostRepository) UpdateStatus(id uint, status int8) error {
	return r.DB.Model(&assModel.Host{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateMonitoring 更新监控状态
func (r *HostRepository) UpdateMonitoring(id uint, monitoringEnabled int8) error {
	return r.DB.Model(&assModel.Host{}).Where("id = ?", id).Update("monitoring_enabled", monitoringEnabled).Error
}

// GetByHostname 根据主机名获取主机
func (r *HostRepository) GetByHostname(hostname string) (*assModel.Host, error) {
	var host assModel.Host
	err := r.DB.Where("hostname = ?", hostname).First(&host).Error
	return &host, err
}

// GetByIPAddress 根据IP地址获取主机
func (r *HostRepository) GetByIPAddress(ipAddress string) (*assModel.Host, error) {
	var host assModel.Host
	err := r.DB.Where("ip_address = ?", ipAddress).First(&host).Error
	return &host, err
}

// GetHostsWithGroup 获取主机及其关联的主机组信息
func (r *HostRepository) GetHostsWithGroup(id uint) (*assModel.Host, error) {
	var host assModel.Host
	err := r.DB.Preload("Group").Where("id = ?", id).First(&host).Error
	return &host, err
}

// GetStatistics 获取主机统计信息
//func (r *HostRepository) GetStatistics() (*assModel.HostStatisticsResponse, error) {
//	var stats assModel.HostStatisticsResponse
//
//	// 获取总主机数
//	r.DB.Model(&assModel.Host{}).Count(&stats.TotalHosts)
//
//	// 获取各状态主机数
//	r.DB.Model(&assModel.Host{}).Where("status = ?", 1).Count(&stats.OnlineHosts)
//	r.DB.Model(&assModel.Host{}).Where("status = ?", 0).Count(&stats.OfflineHosts)
//	r.DB.Model(&assModel.Host{}).Where("status = ?", -1).Count(&stats.FaultHosts)
//
//	// 获取监控状态统计
//	r.DB.Model(&assModel.Host{}).Where("monitoring_enabled = ?", 1).Count(&stats.EnabledMonitoring)
//	r.DB.Model(&assModel.Host{}).Where("monitoring_enabled = ?", 0).Count(&stats.DisabledMonitoring)
//
//	// 按操作系统类型统计
//	stats.ByOSType = make(map[string]int)
//	r.DB.Model(&assModel.Host{}).Select("os_type, count(*) as count").Group("os_type").Scan(&stats.ByOSType)
//
//	// 按主机组统计
//	var groupStats []struct {
//		GroupID   uint
//		GroupName string
//		HostCount int
//	}
//	r.DB.Table("hosts h").
//		Select("hg.id as group_id, hg.name as group_name, count(h.id) as host_count").
//		Joins("left join host_groups hg on h.group_id = hg.id").
//		Group("hg.id, hg.name").
//		Scan(&groupStats)
//
//	stats.ByGroup = make([]assModel.HostGroupStatistics, len(groupStats))
//	for i, gs := range groupStats {
//		stats.ByGroup[i] = assModel.HostGroupStatistics{
//			GroupID:   gs.GroupID,
//			GroupName: gs.GroupName,
//			HostCount: gs.HostCount,
//		}
//	}
//
//	return &stats, nil
//}
