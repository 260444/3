package asset_management

import (
	assModel "backend/internal/model/asset_management"
	"time"

	"gorm.io/gorm"
)

type HostMetricRepository struct {
	DB *gorm.DB
}

func NewHostMetricRepository(db *gorm.DB) *HostMetricRepository {
	return &HostMetricRepository{DB: db}
}

// Create 创建主机指标
func (r *HostMetricRepository) Create(metric *assModel.HostMetric) error {
	return r.DB.Create(metric).Error
}

// BatchCreate 批量创建主机指标
func (r *HostMetricRepository) BatchCreate(metrics []assModel.HostMetric) error {
	return r.DB.Create(&metrics).Error
}

// GetHistory 获取主机指标历史数据
func (r *HostMetricRepository) GetHistory(hostID uint, metricType, metricName string, startTime, endTime *time.Time, page, pageSize int) ([]assModel.HostMetric, int64, error) {
	var metrics []assModel.HostMetric
	var total int64

	query := r.DB.Model(&assModel.HostMetric{}).Where("host_id = ?", hostID)

	// 添加查询条件
	if metricType != "" {
		query = query.Where("metric_type = ?", metricType)
	}
	if metricName != "" {
		query = query.Where("metric_name = ?", metricName)
	}
	if startTime != nil {
		query = query.Where("recorded_at >= ?", *startTime)
	}
	if endTime != nil {
		query = query.Where("recorded_at <= ?", *endTime)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("recorded_at DESC").Find(&metrics).Error; err != nil {
		return nil, 0, err
	}

	return metrics, total, nil
}

// GetLatest 获取主机最新指标
func (r *HostMetricRepository) GetLatest(hostID uint) ([]assModel.HostMetric, error) {
	var metrics []assModel.HostMetric

	// 获取每种指标类型的最新记录
	subQuery := r.DB.Model(&assModel.HostMetric{}).
		Select("MAX(id) as id").
		Where("host_id = ?", hostID).
		Group("metric_type, metric_name")

	err := r.DB.Where("id IN (?)", subQuery).Find(&metrics).Error
	return metrics, err
}

// DeleteOldMetrics 删除过期的历史指标数据
func (r *HostMetricRepository) DeleteOldMetrics(beforeTime time.Time) (int64, error) {
	result := r.DB.Where("recorded_at < ?", beforeTime).Delete(&assModel.HostMetric{})
	return result.RowsAffected, result.Error
}

// GetMetricsByTimeRange 获取指定时间范围内的指标数据
func (r *HostMetricRepository) GetMetricsByTimeRange(hostID uint, metricType, metricName string, startTime, endTime time.Time) ([]assModel.HostMetric, error) {
	var metrics []assModel.HostMetric
	err := r.DB.Where("host_id = ? AND metric_type = ? AND metric_name = ? AND recorded_at BETWEEN ? AND ?",
		hostID, metricType, metricName, startTime, endTime).
		Order("recorded_at ASC").
		Find(&metrics).Error
	return metrics, err
}

// GetAverageMetrics 获取平均指标值
func (r *HostMetricRepository) GetAverageMetrics(hostID uint, metricType, metricName string, startTime, endTime time.Time) (float64, error) {
	var avgValue float64
	err := r.DB.Model(&assModel.HostMetric{}).
		Where("host_id = ? AND metric_type = ? AND metric_name = ? AND recorded_at BETWEEN ? AND ?",
			hostID, metricType, metricName, startTime, endTime).
		Select("COALESCE(AVG(metric_value), 0)").
		Scan(&avgValue).Error
	return avgValue, err
}

// GetMaxMinMetrics 获取最大最小指标值
func (r *HostMetricRepository) GetMaxMinMetrics(hostID uint, metricType, metricName string, startTime, endTime time.Time) (max, min float64, err error) {
	var result struct {
		Max float64 `json:"max"`
		Min float64 `json:"min"`
	}

	err = r.DB.Model(&assModel.HostMetric{}).
		Where("host_id = ? AND metric_type = ? AND metric_name = ? AND recorded_at BETWEEN ? AND ?",
			hostID, metricType, metricName, startTime, endTime).
		Select("MAX(metric_value) as max, MIN(metric_value) as min").
		Scan(&result).Error

	return result.Max, result.Min, err
}
