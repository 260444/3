package asset_management

import (
	assModel "backend/internal/model/asset_management"
	assRepo "backend/internal/repository/asset_management"
	"backend/pkg/response"
	"errors"
	"time"

	"gorm.io/gorm"
)

type HostMetricService struct {
	hostMetricRepo *assRepo.HostMetricRepository
	hostRepo       *assRepo.HostRepository
}

func NewHostMetricService(
	hostMetricRepo *assRepo.HostMetricRepository,
	hostRepo *assRepo.HostRepository,
) *HostMetricService {
	return &HostMetricService{
		hostMetricRepo: hostMetricRepo,
		hostRepo:       hostRepo,
	}
}

// ReportHostMetrics 上报主机指标
func (s *HostMetricService) ReportHostMetrics(req *assModel.HostMetricsRequest) (int64, error) {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(req.HostID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, response.ErrValidationError
		}
		return 0, response.ErrDatabaseError
	}

	// 转换指标数据
	var metrics []assModel.HostMetric
	for _, metricDTO := range req.Metrics {
		metric := assModel.HostMetric{
			HostID:      req.HostID,
			MetricType:  metricDTO.MetricType,
			MetricName:  metricDTO.MetricName,
			MetricValue: metricDTO.MetricValue,
			Unit:        metricDTO.Unit,
			RecordedAt:  time.Now(),
		}
		metrics = append(metrics, metric)
	}

	// 批量插入指标数据
	if err := s.hostMetricRepo.BatchCreate(metrics); err != nil {
		return 0, response.ErrDatabaseError
	}

	return int64(len(metrics)), nil
}

// GetHostMetricsHistory 获取主机指标历史数据
func (s *HostMetricService) GetHostMetricsHistory(hostID uint, metricType, metricName string, startTime, endTime *time.Time, page, pageSize int) ([]assModel.HostMetric, int64, error) {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(hostID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, response.ErrNotFound
		}
		return nil, 0, response.ErrDatabaseError
	}

	metrics, total, err := s.hostMetricRepo.GetHistory(hostID, metricType, metricName, startTime, endTime, page, pageSize)
	if err != nil {
		return nil, 0, response.ErrDatabaseError
	}

	return metrics, total, nil
}

// GetHostLatestMetrics 获取主机最新指标
func (s *HostMetricService) GetHostLatestMetrics(hostID uint) ([]assModel.HostMetric, error) {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(hostID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}

	metrics, err := s.hostMetricRepo.GetLatest(hostID)
	if err != nil {
		return nil, response.ErrDatabaseError
	}

	return metrics, nil
}

// GetHostMetricsByTimeRange 获取指定时间范围内的指标数据
func (s *HostMetricService) GetHostMetricsByTimeRange(hostID uint, metricType, metricName string, startTime, endTime time.Time) ([]assModel.HostMetric, error) {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(hostID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}

	metrics, err := s.hostMetricRepo.GetMetricsByTimeRange(hostID, metricType, metricName, startTime, endTime)
	if err != nil {
		return nil, response.ErrDatabaseError
	}

	return metrics, nil
}

// GetHostMetricsStatistics 获取主机指标统计信息
func (s *HostMetricService) GetHostMetricsStatistics(hostID uint, metricType, metricName string, startTime, endTime time.Time) (map[string]interface{}, error) {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(hostID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}

	// 获取平均值
	avgValue, err := s.hostMetricRepo.GetAverageMetrics(hostID, metricType, metricName, startTime, endTime)
	if err != nil {
		return nil, response.ErrDatabaseError
	}

	// 获取最大最小值
	maxValue, minValue, err := s.hostMetricRepo.GetMaxMinMetrics(hostID, metricType, metricName, startTime, endTime)
	if err != nil {
		return nil, response.ErrDatabaseError
	}

	statistics := map[string]interface{}{
		"average": avgValue,
		"maximum": maxValue,
		"minimum": minValue,
	}

	return statistics, nil
}

// CleanupOldMetrics 清理过期的历史指标数据
func (s *HostMetricService) CleanupOldMetrics(beforeTime time.Time) (int64, error) {
	affected, err := s.hostMetricRepo.DeleteOldMetrics(beforeTime)
	if err != nil {
		return 0, response.ErrDatabaseError
	}
	return affected, nil
}

// GetHostMetricsOverview 获取主机指标概览
func (s *HostMetricService) GetHostMetricsOverview(hostID uint) (map[string]interface{}, error) {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(hostID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}

	// 获取最新的各项指标
	latestMetrics, err := s.hostMetricRepo.GetLatest(hostID)
	if err != nil {
		return nil, response.ErrDatabaseError
	}

	// 按指标类型分组
	metricsByType := make(map[string][]assModel.HostMetric)
	for _, metric := range latestMetrics {
		metricsByType[metric.MetricType] = append(metricsByType[metric.MetricType], metric)
	}

	overview := make(map[string]interface{})
	for metricType, metrics := range metricsByType {
		typeMetrics := make(map[string]interface{})
		for _, metric := range metrics {
			typeMetrics[metric.MetricName] = map[string]interface{}{
				"value": metric.MetricValue,
				"unit":  metric.Unit,
				"time":  metric.RecordedAt,
			}
		}
		overview[metricType] = typeMetrics
	}

	return overview, nil
}
