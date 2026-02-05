package asset_management

import (
	"time"

	"gorm.io/gorm"
)

// Host 主机模型
type Host struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Hostname         string         `gorm:"size:100;not null;uniqueIndex" json:"hostname"`                 // 主机名
	IPAddress        string         `gorm:"size:45;not null;uniqueIndex" json:"ip_address"`                // IP地址
	Port             uint16         `gorm:"default:22" json:"port"`                                        // SSH端口
	Username         string         `gorm:"size:50;not null" json:"username"`                              // 登录用户名
	Password         string         `gorm:"size:255;not null" json:"password"`                             // 加密后的密码
	OSType           string         `gorm:"size:20;default:'linux'" json:"os_type"`                        // 操作系统类型
	CPUCores         *uint16        `json:"cpu_cores"`                                                     // CPU核心数
	MemoryGB         *uint16        `json:"memory_gb"`                                                     // 内存大小(GB)
	DiskSpaceGB      *uint32        `json:"disk_space_gb"`                                                 // 磁盘空间(GB)
	GroupID          uint           `gorm:"not null" json:"group_id"`                                      // 所属主机组ID
	Status           int8           `gorm:"default:1" json:"status"`                                       // 主机状态: 1-在线, 0-离线, -1-故障
	MonitoringEnable int8           `gorm:"column:monitoring_enabled;default:1" json:"monitoring_enabled"` // 监控是否启用
	LastHeartbeat    *time.Time     `json:"last_heartbeat"`                                                // 最后心跳时间
	Description      string         `gorm:"size:500" json:"description"`                                   // 主机描述
	CreatedBy        *uint          `json:"created_by"`                                                    // 创建人用户ID
	UpdatedBy        *uint          `json:"updated_by"`                                                    // 更新人用户ID
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Group            *HostGroup     `gorm:"foreignKey:GroupID" json:"group"` // 关联的主机组
}

// HostGroup 主机组模型
type HostGroup struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:100;not null;uniqueIndex" json:"name"` // 主机组名称
	Description string         `gorm:"size:500" json:"description"`               // 描述信息
	Status      int8           `gorm:"default:1" json:"status"`                   // 状态: 1-启用, 0-禁用
	CreatedBy   *uint          `json:"created_by"`                                // 创建人用户ID
	UpdatedBy   *uint          `json:"updated_by"`                                // 更新人用户ID
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Hosts       []Host         `gorm:"foreignKey:GroupID" json:"hosts"` // 关联的主机
}

// HostMetric 主机监控指标模型
type HostMetric struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	HostID      uint      `gorm:"not null" json:"host_id"`                         // 主机ID
	MetricType  string    `gorm:"size:30;not null" json:"metric_type"`             // 指标类型: cpu,memory,disk,network
	MetricName  string    `gorm:"size:50;not null" json:"metric_name"`             // 指标名称
	MetricValue float64   `gorm:"type:decimal(10,2);not null" json:"metric_value"` // 指标值
	Unit        string    `gorm:"size:20" json:"unit"`                             // 单位
	RecordedAt  time.Time `gorm:"not null" json:"recorded_at"`                     // 记录时间
	Host        *Host     `gorm:"foreignKey:HostID" json:"host"`                   // 关联的主机
}

// HostCreateRequest 创建主机请求
type HostCreateRequest struct {
	Hostname    string  `json:"hostname" binding:"required"`
	IPAddress   string  `json:"ip_address" binding:"required"`
	Port        uint16  `json:"port" binding:"required"`
	Username    string  `json:"username" binding:"required"`
	Password    string  `json:"password" binding:"required"`
	OSType      string  `json:"os_type" binding:"required,oneof=linux windows"`
	CPUCores    *uint16 `json:"cpu_cores"`
	MemoryGB    *uint16 `json:"memory_gb"`
	DiskSpaceGB *uint32 `json:"disk_space_gb"`
	GroupID     uint    `json:"group_id" binding:"required"`
	Description string  `json:"description"`
}

// HostUpdateRequest 更新主机请求
type HostUpdateRequest struct {
	Hostname    string  `json:"hostname"`
	IPAddress   string  `json:"ip_address"`
	Port        uint16  `json:"port"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	OSType      string  `json:"os_type" binding:"omitempty,oneof=linux windows"`
	CPUCores    *uint16 `json:"cpu_cores"`
	MemoryGB    *uint16 `json:"memory_gb"`
	DiskSpaceGB *uint32 `json:"disk_space_gb"`
	GroupID     uint    `json:"group_id"`
	Description string  `json:"description"`
}

// HostStatusUpdateRequest 更新主机状态请求
type HostStatusUpdateRequest struct {
	Status int8 `json:"status" binding:"required,oneof=1 2 3"`
}

// HostMonitoringUpdateRequest 更新监控状态请求
type HostMonitoringUpdateRequest struct {
	MonitoringEnabled int8 `json:"monitoring_enabled" binding:"required,oneof=1 2"`
}

// HostGroupCreateRequest 创建主机组请求
type HostGroupCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// HostGroupUpdateRequest 更新主机组请求
type HostGroupUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// HostGroupStatusUpdateRequest 更新主机组状态请求
type HostGroupStatusUpdateRequest struct {
	Status int8 `json:"status" binding:"required,oneof=1 2"`
}

// HostMetricsRequest 上报主机指标请求
type HostMetricsRequest struct {
	HostID  uint            `json:"host_id" binding:"required"`
	Metrics []HostMetricDTO `json:"metrics" binding:"required"`
}

// HostMetricDTO 主机指标DTO
type HostMetricDTO struct {
	MetricType  string  `json:"metric_type" binding:"required"`
	MetricName  string  `json:"metric_name" binding:"required"`
	MetricValue float64 `json:"metric_value" binding:"required"`
	Unit        string  `json:"unit"`
}

// HostStatisticsResponse 主机统计信息响应
type HostStatisticsResponse struct {
	TotalHosts         int                   `json:"total_hosts"`
	OnlineHosts        int                   `json:"online_hosts"`
	OfflineHosts       int                   `json:"offline_hosts"`
	FaultHosts         int                   `json:"fault_hosts"`
	EnabledMonitoring  int                   `json:"enabled_monitoring"`
	DisabledMonitoring int                   `json:"disabled_monitoring"`
	ByOSType           map[string]int        `json:"by_os_type"`
	ByGroup            []HostGroupStatistics `json:"by_group"`
}

// HostGroupStatistics 主机组统计信息
type HostGroupStatistics struct {
	GroupID   uint   `json:"group_id"`
	GroupName string `json:"group_name"`
	HostCount int    `json:"host_count"`
}

// TableName 设置表名
func (Host) TableName() string {
	return "hosts"
}

func (HostGroup) TableName() string {
	return "host_groups"
}

func (HostMetric) TableName() string {
	return "host_metrics"
}
