package service

import (
	assModel "backend/internal/model/asset_management"
	assRepo "backend/internal/repository/asset_management"
	"backend/pkg/logger"
	"context"

	"go.uber.org/zap"
)

type HostMonitorService struct {
	promClient *assRepo.PrometheusClient
	hostRepo   *assRepo.HostRepository
}

func NewHostMonitorService(promClient *assRepo.PrometheusClient, hostRepo *assRepo.HostRepository) *HostMonitorService {
	return &HostMonitorService{
		promClient: promClient,
		hostRepo:   hostRepo,
	}
}

// 获取所有主机的监控指标
func (s *HostMonitorService) GetAllHostsMetrics(ctx context.Context) ([]assModel.HostMetric, error) {
	hosts, err := s.hostRepo.ListForMonitoring(ctx)
	if err != nil {
		return nil, err
	}

	var results []assModel.HostMetric
	for _, host := range hosts {
		metrics, err := s.promClient.GetHostMetrics(ctx, host.IPAddress)
		if err != nil {
			// 记录错误但继续处理其他主机
			logger.Logger.Error("get host metrics failed",
				zap.Uint("host_id", host.ID),
				zap.Error(err))
			continue
		}

		results = append(results, assModel.HostMetric{
			HostID:      host.ID,
			CPUUsage:    metrics.CPUUsage,
			MemoryUsage: metrics.MemoryUsage,
			DiskUsage:   metrics.DiskUsage,
			RecordedAt:  metrics.RecordedAt,
		})
	}

	return results, nil
}

// 更新数据库中的监控指标
func (s *HostMonitorService) UpdateHostMetrics(ctx context.Context) error {
	hosts, err := s.hostRepo.ListForMonitoring(ctx)
	if err != nil {
		return err
	}

	for _, host := range hosts {
		metrics, err := s.promClient.GetHostMetrics(ctx, host.IPAddress)
		if err != nil {
			logger.Logger.Error("get host metrics failed",
				zap.Uint("host_id", host.ID),
				zap.Error(err))
			continue
		}

		// 保存到数据库
		err = s.hostRepo.SaveMetrics(ctx, host.ID, metrics)
		if err != nil {
			logger.Logger.Error("save host metrics failed",
				zap.Uint("host_id", host.ID),
				zap.Error(err))
		}
	}

	return nil
}
