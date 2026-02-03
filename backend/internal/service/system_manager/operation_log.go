package system_manager

import (
	"backend/internal/model/system_manager"
	repository "backend/internal/repository/system_manager"
)

// OperationLogService 操作日志服务
type OperationLogService struct {
	OperationLogRepo *repository.OperationLogRepository
}

// NewOperationLogService 创建操作日志服务
func NewOperationLogService(operationLogRepo *repository.OperationLogRepository) *OperationLogService {
	return &OperationLogService{
		OperationLogRepo: operationLogRepo,
	}
}

// CreateOperationLog 创建操作日志
func (s *OperationLogService) CreateOperationLog(log *system_manager.OperationLog) error {
	return s.OperationLogRepo.Create(log)
}

// GetOperationLogByID 根据ID获取操作日志
func (s *OperationLogService) GetOperationLogByID(id uint) (*system_manager.OperationLog, error) {
	return s.OperationLogRepo.GetByID(id)
}

// GetOperationLogs 获取操作日志列表
func (s *OperationLogService) GetOperationLogs(limit, offset int) ([]system_manager.OperationLog, int64, error) {
	logs, err := s.OperationLogRepo.List(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.OperationLogRepo.GetTotal()
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// DeleteOperationLog 删除操作日志
func (s *OperationLogService) DeleteOperationLog(id uint) error {
	return s.OperationLogRepo.Delete(id)
}
