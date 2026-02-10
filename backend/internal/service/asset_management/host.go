package asset_management

import (
	assModel "backend/internal/model/asset_management"
	assRepo "backend/internal/repository/asset_management"

	"backend/pkg/logger"
	"backend/pkg/response"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HostService struct {
	hostRepo      *assRepo.HostRepository
	hostGroupRepo *assRepo.HostGroupRepository
	// hostMetricRepo *assRepo.HostMetricRepository
	// credentialRepo *assRepo.CredentialRepository
}

func NewHostService(
	hostRepo *assRepo.HostRepository,
	hostGroupRepo *assRepo.HostGroupRepository,
	// hostMetricRepo *assRepo.HostMetricRepository,
	// credentialService *assRepo.CredentialRepository,
) *HostService {
	return &HostService{
		hostRepo:      hostRepo,
		hostGroupRepo: hostGroupRepo,
		// hostMetricRepo: hostMetricRepo,
		// credentialRepo: credentialService,
	}
}

// CreateHost 创建主机
func (s *HostService) CreateHost(req *assModel.HostCreateRequest, userID uint) (*assModel.Host, error) {
	// 检查主机名是否已存在
	if _, err := s.hostRepo.GetByHostname(req.Hostname); err == nil {
		return nil, response.ErrValidationError
	}

	// 检查IP地址是否已存在
	if _, err := s.hostRepo.GetByIPAddress(req.IPAddress); err == nil {
		return nil, response.ErrValidationError
	}

	// 检查主机组是否存在
	if _, err := s.hostGroupRepo.GetByID(req.GroupID); err != nil {
		return nil, response.ErrValidationError
	}

	host := &assModel.Host{
		Hostname:         req.Hostname,
		IPAddress:        req.IPAddress,
		Port:             req.Port,
		OSType:           req.OSType,
		CPUCores:         req.CPUCores,
		MemoryGB:         req.MemoryGB,
		DiskSpaceGB:      req.DiskSpaceGB,
		GroupID:          req.GroupID,
		Status:           1, // 默认在线
		MonitoringEnable: 1, // 默认启用监控
		Description:      req.Description,
		CreatedBy:        &userID,
		UpdatedBy:        &userID,
	}

	if err := s.hostRepo.Create(host); err != nil {
		logger.Logger.Info("写入主机信息失败", zap.String("hostname", host.Hostname), zap.Error(err))
		return nil, response.ErrDatabaseError
	}

	// 如果提供了凭据ID，则建立关联关系
	if len(req.CredentialIDs) > 0 {
		// 开始事务处理凭据关联
		tx := s.hostRepo.DB.Begin()
		defer func() {
			if r := tx.Rollback(); r != nil {
				logger.Logger.Error("事务回滚错误", zap.Error(r.Error))
			}
		}()

		// 检查所有凭据ID是否存在
		var existingCredentials []assModel.Credential
		if err := tx.Find(&existingCredentials, "id IN ?", req.CredentialIDs).Error; err != nil {
			tx.Rollback()
			return nil, response.ErrDatabaseError
		}

		if len(existingCredentials) != len(req.CredentialIDs) {
			tx.Rollback()
			return nil, errors.New("存在不存在的凭据ID")
		}

		// 建立主机与凭据的关联
		if err := tx.Model(host).Association("Credentials").Append(&existingCredentials); err != nil {
			tx.Rollback()
			return nil, response.ErrDatabaseError
		}

		if err := tx.Commit().Error; err != nil {
			logger.Logger.Error("提交事务失败", zap.Error(err))
			return nil, response.ErrDatabaseError
		}
	}

	return host, nil
}

// GetHostByID 根据ID获取主机
func (s *HostService) GetHostByID(id uint) (*assModel.Host, error) {
	host, err := s.hostRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}
	return host, nil
}

// ListHosts 获取主机列表
func (s *HostService) ListHosts(page, pageSize int, hostname, ipAddress string, groupID *uint, status *int8, osType string) ([]assModel.Host, int64, error) {
	hosts, total, err := s.hostRepo.List(page, pageSize, hostname, ipAddress, groupID, status, osType)
	if err != nil {
		return nil, 0, response.ErrDatabaseError
	}
	return hosts, total, nil
}

// UpdateHost 更新主机
func (s *HostService) UpdateHost(id uint, req *assModel.HostUpdateRequest, userID uint) (*assModel.Host, error) {
	host, err := s.hostRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}

	// 检查主机名是否重复（排除自己）
	if req.Hostname != "" && req.Hostname != host.Hostname {
		if existing, _ := s.hostRepo.GetByHostname(req.Hostname); existing != nil && existing.ID != id {
			return nil, response.ErrValidationError
		}
		host.Hostname = req.Hostname
	}

	// 检查IP地址是否重复（排除自己）
	if req.IPAddress != "" && req.IPAddress != host.IPAddress {
		if existing, _ := s.hostRepo.GetByIPAddress(req.IPAddress); existing != nil && existing.ID != id {
			return nil, response.ErrValidationError
		}
		host.IPAddress = req.IPAddress
	}

	// 更新其他字段
	if req.Port != 0 {
		host.Port = req.Port
	}
	if req.OSType != "" {
		host.OSType = req.OSType
	}
	if req.CPUCores != nil {
		host.CPUCores = req.CPUCores
	}
	if req.MemoryGB != nil {
		host.MemoryGB = req.MemoryGB
	}
	if req.DiskSpaceGB != nil {
		host.DiskSpaceGB = req.DiskSpaceGB
	}
	if req.GroupID != 0 {
		// 检查主机组是否存在
		if _, err := s.hostGroupRepo.GetByID(req.GroupID); err != nil {
			return nil, response.ErrValidationError
		}
		host.GroupID = req.GroupID
	}
	if req.Description != "" {
		host.Description = req.Description
	}

	host.UpdatedBy = &userID

	if err := s.hostRepo.Update(host); err != nil {
		return nil, response.ErrDatabaseError
	}

	// 如果提供了凭据ID，则更新关联关系
	if req.CredentialIDs != nil {
		// 开始事务处理凭据关联
		tx := s.hostRepo.DB.Begin()
		defer func() {
			// 只有在panic等异常情况下才回滚
			if r := recover(); r != nil {
				tx.Rollback()
				logger.Logger.Error("事务因异常回滚", zap.Any("panic", r))
				// 重新抛出panic
				panic(r)
			}
		}()

		// 清除现有的凭据关联
		if err := tx.Model(host).Association("Credentials").Clear(); err != nil {
			tx.Rollback()
			return nil, response.ErrDatabaseError
		}

		// 如果提供了新的凭据ID列表，则建立新的关联
		if len(req.CredentialIDs) > 0 {
			// 检查所有凭据ID是否存在
			var existingCredentials []assModel.Credential
			if err := tx.Find(&existingCredentials, "id IN ?", req.CredentialIDs).Error; err != nil {
				tx.Rollback()
				return nil, response.ErrDatabaseError
			}

			if len(existingCredentials) != len(req.CredentialIDs) {
				tx.Rollback()
				return nil, errors.New("存在不存在的凭据ID")
			}

			// 建立主机与凭据的关联
			if err := tx.Model(host).Association("Credentials").Append(&existingCredentials); err != nil {
				tx.Rollback()
				return nil, response.ErrDatabaseError
			}
		}

		if err := tx.Commit().Error; err != nil {
			logger.Logger.Error("提交事务失败", zap.Error(err))
			return nil, response.ErrDatabaseError
		}
	}

	return host, nil
}

// DeleteHost 删除主机
func (s *HostService) DeleteHost(id uint) error {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return response.ErrDatabaseError
	}

	if err := s.hostRepo.Delete(id); err != nil {
		return response.ErrDatabaseError
	}

	return nil
}

// BatchDeleteHosts 批量删除主机
func (s *HostService) BatchDeleteHosts(ids []uint) (int64, error) {
	if len(ids) == 0 {
		return 0, response.ErrValidationError
	}

	affected, err := s.hostRepo.BatchDelete(ids)
	if err != nil {
		return 0, response.ErrDatabaseError
	}

	return affected, nil
}

// UpdateHostStatus 更新主机状态
func (s *HostService) UpdateHostStatus(id uint, status int8) error {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return response.ErrDatabaseError
	}

	if err := s.hostRepo.UpdateStatus(id, status); err != nil {
		return response.ErrDatabaseError
	}

	return nil
}

// UpdateHostMonitoring 更新主机监控状态
func (s *HostService) UpdateHostMonitoring(id uint, monitoringEnabled int8) error {
	// 检查主机是否存在
	if _, err := s.hostRepo.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return response.ErrDatabaseError
	}

	if err := s.hostRepo.UpdateMonitoring(id, monitoringEnabled); err != nil {
		return response.ErrDatabaseError
	}

	return nil
}

// GetHostStatistics 获取主机统计信息
//func (s *HostService) GetHostStatistics() (*assModel.HostStatisticsResponse, error) {
//	stats, err := s.hostRepo.GetStatistics()
//	if err != nil {
//		return nil, response.ErrDatabaseError
//	}
//	return stats, nil
//}
