package asset_management

import (
	assModel "backend/internal/model/asset_management"
	assRepo "backend/internal/repository/asset_management"
	"backend/pkg/response"
	"errors"

	"gorm.io/gorm"
)

type HostGroupService struct {
	hostGroupRepo *assRepo.HostGroupRepository
	hostRepo      *assRepo.HostRepository
}

func NewHostGroupService(
	hostGroupRepo *assRepo.HostGroupRepository,
	hostRepo *assRepo.HostRepository,
) *HostGroupService {
	return &HostGroupService{
		hostGroupRepo: hostGroupRepo,
		hostRepo:      hostRepo,
	}
}

// CreateHostGroup 创建主机组
func (s *HostGroupService) CreateHostGroup(req *assModel.HostGroupCreateRequest, userID uint) (*assModel.HostGroup, error) {
	// 检查主机组名称是否已存在
	if _, err := s.hostGroupRepo.GetByName(req.Name); err == nil {
		return nil, response.ErrValidationError
	}

	group := &assModel.HostGroup{
		Name:        req.Name,
		Description: req.Description,
		Status:      1, // 默认启用
		CreatedBy:   &userID,
	}

	if err := s.hostGroupRepo.Create(group); err != nil {
		return nil, response.ErrDatabaseError
	}

	return group, nil
}

// GetHostGroupByID 根据ID获取主机组
func (s *HostGroupService) GetHostGroupByID(id uint) (*assModel.HostGroup, error) {
	group, err := s.hostGroupRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}
	return group, nil
}

// ListHostGroups 获取主机组列表
func (s *HostGroupService) ListHostGroups(page, pageSize int, name string, status *int8) ([]assModel.HostGroup, int64, error) {
	groups, total, err := s.hostGroupRepo.List(page, pageSize, name, status)
	if err != nil {
		return nil, 0, response.ErrDatabaseError
	}

	// 为每个主机组添加主机数量
	for i := range groups {
		var hostCount int64
		s.hostRepo.DB.Model(&assModel.Host{}).Where("group_id = ?", groups[i].ID).Count(&hostCount)
		// 可以在这里添加一个HostCount字段到HostGroup模型中
	}

	return groups, total, nil
}

// UpdateHostGroup 更新主机组
func (s *HostGroupService) UpdateHostGroup(id uint, req *assModel.HostGroupUpdateRequest, userID uint) (*assModel.HostGroup, error) {
	group, err := s.hostGroupRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}

	// 检查名称是否重复（排除自己）
	if req.Name != "" && req.Name != group.Name {
		if existing, _ := s.hostGroupRepo.GetByName(req.Name); existing != nil && existing.ID != id {
			return nil, response.ErrValidationError
		}
		group.Name = req.Name
	}

	if req.Description != "" {
		group.Description = req.Description
	}

	group.UpdatedBy = &userID

	if err := s.hostGroupRepo.Update(group); err != nil {
		return nil, response.ErrDatabaseError
	}

	return group, nil
}

// DeleteHostGroup 删除主机组
func (s *HostGroupService) DeleteHostGroup(id uint) error {
	// 检查主机组是否存在
	_, err := s.hostGroupRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return response.ErrDatabaseError
	}

	// 检查主机组是否有关联的主机
	hasHosts, err := s.hostGroupRepo.CheckGroupHasHosts(id)
	if err != nil {
		return response.ErrDatabaseError
	}
	if hasHosts {
		return response.ErrValidationError // 可以定义专门的错误类型
	}

	if err := s.hostGroupRepo.Delete(id); err != nil {
		return response.ErrDatabaseError
	}

	return nil
}

// UpdateHostGroupStatus 更新主机组状态
func (s *HostGroupService) UpdateHostGroupStatus(id uint, status int8) error {
	// 检查主机组是否存在
	if _, err := s.hostGroupRepo.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return response.ErrDatabaseError
	}

	if err := s.hostGroupRepo.UpdateStatus(id, status); err != nil {
		return response.ErrDatabaseError
	}

	return nil
}

// GetHostGroupWithHosts 获取主机组及其主机信息
func (s *HostGroupService) GetHostGroupWithHosts(id uint) (*assModel.HostGroup, error) {
	group, err := s.hostGroupRepo.GetGroupWithHosts(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, response.ErrDatabaseError
	}
	return group, nil
}
