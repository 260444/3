package system_manager

import (
	"backend/internal/model/system_manager"

	"gorm.io/gorm"
)

// OperationLogRepository 操作日志数据访问层
type OperationLogRepository struct {
	DB *gorm.DB
}

// NewOperationLogRepository 创建操作日志仓库
func NewOperationLogRepository(db *gorm.DB) *OperationLogRepository {
	return &OperationLogRepository{DB: db}
}

// Create 创建操作日志
func (r *OperationLogRepository) Create(log *system_manager.OperationLog) error {
	return r.DB.Create(log).Error
}

// GetByID 根据ID获取操作日志
func (r *OperationLogRepository) GetByID(id uint) (*system_manager.OperationLog, error) {
	var log system_manager.OperationLog
	err := r.DB.First(&log, id).Error
	return &log, err
}

// List 获取操作日志列表
func (r *OperationLogRepository) List(limit, offset int) ([]system_manager.OperationLog, error) {
	var logs []system_manager.OperationLog
	err := r.DB.Offset(offset).Limit(limit).Order("created_at DESC").Find(&logs).Error
	return logs, err
}

// GetTotal 获取操作日志总数
func (r *OperationLogRepository) GetTotal() (int64, error) {
	var count int64
	err := r.DB.Model(&system_manager.OperationLog{}).Count(&count).Error
	return count, err
}

// Delete 删除操作日志
func (r *OperationLogRepository) Delete(id uint) error {
	return r.DB.Delete(&system_manager.OperationLog{}, id).Error
}
