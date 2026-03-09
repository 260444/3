package asset_management

import (
	assModel "backend/internal/model/asset_management"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"gorm.io/gorm"
)

type HostMetricRepository struct {
	DB     *gorm.DB
	client *http.Client
}

func NewHostMetricRepository(db *gorm.DB) *HostMetricRepository {
	return &HostMetricRepository{
		DB:     db,
		client: &http.Client{Timeout: 10 * time.Second},
	}
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

type PromQLResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]interface{} `json:"metric"`
			Value  []interface{}          `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// GetCPUUsage 查询 CPU 使用率
func (pc *HostMetricRepository) GetCPUUsage(ctx context.Context) (float64, error) {
	query := `100 - (avg by(instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)`
	return pc.querySingleValue(ctx, query)
}

// GetMemoryUsage 查询内存使用率
func (pc *HostMetricRepository) GetMemoryUsage(ctx context.Context) (float64, error) {
	query := `(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100`
	return pc.querySingleValue(ctx, query)
}

// GetDiskUsage 查询磁盘使用率
func (pc *HostMetricRepository) GetDiskUsage(ctx context.Context, mountpoint string) (float64, error) {
	query := fmt.Sprintf(`(1 - (node_filesystem_avail_bytes{mountpoint="%s"} / node_filesystem_size_bytes{mountpoint="%s"})) * 100`,
		mountpoint, mountpoint)
	return pc.querySingleValue(ctx, query)
}

// GetHostMetrics 查询主机监控指标（用于你的资产管理系统）
func (pc *HostMetricRepository) GetHostMetrics(ctx context.Context, hostIP string) (*assModel.HostMetric, error) {
	// CPU 使用率
	cpuUsage, err := pc.GetCPUUsage(ctx)
	if err != nil {
		return nil, err
	}

	// 内存使用率
	memoryUsage, err := pc.GetMemoryUsage(ctx)
	if err != nil {
		return nil, err
	}

	// 磁盘使用率
	diskUsage, err := pc.GetDiskUsage(ctx, "/")
	if err != nil {
		return nil, err
	}

	return &assModel.HostMetric{
		CPUUsage:    cpuUsage,
		MemoryUsage: memoryUsage,
		DiskUsage:   diskUsage,
		RecordedAt:  time.Now(),
	}, nil
}

func (pc *HostMetricRepository) querySingleValue(ctx context.Context, query string) (float64, error) {
	u, _ := url.Parse("http://localhost:9090" + "/api/v1/query")
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	resp, err := pc.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var promResp PromQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&promResp); err != nil {
		return 0, err
	}

	if promResp.Status != "success" || len(promResp.Data.Result) == 0 {
		return 0, fmt.Errorf("no data returned from prometheus")
	}

	// Prometheus 返回的 value 格式：[timestamp, value]
	value := promResp.Data.Result[0].Value[1]
	var result float64
	fmt.Sscanf(value.(string), "%f", &result)

	return result, nil
}
