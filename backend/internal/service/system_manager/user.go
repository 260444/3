package system_manager

import (
	"backend/internal/model/system_manager"
	repository "backend/internal/repository/system_manager"
	"backend/pkg/casbin"
	"backend/pkg/logger"
	"errors"
	"time"

	"go.uber.org/zap"

	"golang.org/x/crypto/bcrypt"
)

// UserService 用户服务
type UserService struct {
	UserRepo *repository.UserRepository
	RoleRepo *repository.RoleRepository
}

// NewUserService 创建用户服务
func NewUserService(userRepo *repository.UserRepository, roleRepo *repository.RoleRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
		RoleRepo: roleRepo,
	}
}

// CreateUser 创建用户 *
func (s *UserService) CreateUser(username, password, email, nickname string) (*system_manager.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.UserRepo.GetByUsername(username)

	if existingUser.Username != "" {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	if email != "" {
		existingUser, _ = s.UserRepo.GetByEmail(email)
		if existingUser.Email != "" {
			return nil, errors.New("邮箱已存在")
		}
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	logger.Logger.Info("加密免密:",
		zap.String("密码", password),
		zap.String("加密后的密码", string(hashedPassword)))

	if err != nil {
		return nil, err
	}

	user := &system_manager.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Nickname: nickname,
		Status:   1, // 默认启用
	}

	// 创建用户
	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Login 用户登录 *
func (s *UserService) Login(username, password string) (*system_manager.UserWithRoleInfo, error) {
	UserWithRole, err := s.UserRepo.UserWithRoleInfo(username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	user := UserWithRole.User
	if user.Status == 0 {
		return nil, errors.New("用户已被禁用")
	}

	logger.Logger.Info("用户登录请求",
		zap.String("username", username),
		zap.String("user.Password", user.Password),
		zap.String("password", password))
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 更新最后登录时间和IP
	now := time.Now()
	user.LastLoginAt = &now
	// 注意：在实际应用中需要从请求中获取IP
	// user.LastLoginIP = 获取IP的逻辑
	_ = s.UserRepo.Update(&user)

	return UserWithRole, nil
}

// GetUserByID 根据ID获取用户 *
func (s *UserService) GetUserByID(id uint) (*system_manager.User, error) {
	return s.UserRepo.GetByID(id)
}

// UpdateUser 更新用户  *
func (s *UserService) UpdateUser(id uint, user *system_manager.User) error {
	existingUser, err := s.UserRepo.GetByID(id)
	if err != nil {
		return errors.New("用户不存在")
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Nickname = user.Nickname
	existingUser.Phone = user.Phone
	existingUser.Avatar = user.Avatar

	return s.UserRepo.Update(existingUser)
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(id uint, oldPassword, newPassword string) error {
	user, err := s.UserRepo.GetByID(id)
	if err != nil {
		return errors.New("用户不存在")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("原密码错误")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.UserRepo.UpdatePassword(id, string(hashedPassword))
}

// AddRoleForUser 为用户分配角色 *
func (s *UserService) AddRoleForUser(username string, roleIdent string) error {
	user, err := s.UserRepo.GetByUsername(username)
	if err != nil {
		return err // 假设 GetByUsername 出错会返回错误
	}
	if user.ID == 0 {
		return errors.New("用户不存在")
	}

	// 构造 Casbin 的 subject，直接使用 username
	sub := user.Username
	// 将用户添加到角色组 (g policy)
	// g, username, role_ident
	_, err = casbin.Enforcer.AddGroupingPolicy(sub, roleIdent)
	if err != nil {
		return err
	}

	// 同步更新 RoleID
	role, err := s.RoleRepo.GetByIdent(roleIdent)
	if err == nil && role.ID > 0 {
		user.RoleID = &role.ID
		err = s.UserRepo.Update(user)
		if err != nil {
			logger.Logger.Error("更新用户角色时出错", zap.Uint("roleID: ", role.ID), zap.Error(err))
			return err
		}

	}
	return nil
}

// RemoveRoleForUser 移除用户的角色 *
func (s *UserService) RemoveRoleForUser(username string, roleIdent string) error {
	user, err := s.UserRepo.GetByUsername(username)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("用户不存在")
	}

	sub := user.Username
	_, err = casbin.Enforcer.RemoveGroupingPolicy(sub, roleIdent)
	if err != nil {
		return err
	}

	// 如果移除了角色，且当前 RoleID 对应的就是该角色，则清空 RoleID
	role, err := s.RoleRepo.GetByIdent(roleIdent)
	if err == nil && role.ID > 0 && user.RoleID != nil && *user.RoleID == role.ID {
		user.RoleID = nil
		err = s.UserRepo.Update(user)
		if err != nil {
			logger.Logger.Error("移除用户角色时出错", zap.Error(err))
			return err
		}
	}
	return nil
}

// GetUserRoles 获取用户的角色列表 *
func (s *UserService) GetUserRoles(username string) ([]string, error) {
	user, err := s.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	sub := user.Username
	// 获取用户的所有角色 (g policy 中的第二项)
	roles, err := casbin.Enforcer.GetRolesForUser(sub)
	return roles, err
}

// UpdateUserStatus 更新用户状态 *
func (s *UserService) UpdateUserStatus(id uint, status int) error {
	return s.UserRepo.UpdateStatus(id, status)
}

// DeleteUser 删除用户 *
func (s *UserService) DeleteUser(id uint) error {
	return s.UserRepo.Delete(id)
}

// GetUsers 获取用户列表  *
func (s *UserService) GetUsers(limit, offset int) ([]system_manager.User, int64, error) {
	users, err := s.UserRepo.List(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.UserRepo.GetTotal()
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// ResetPassword 重置用户密码（管理员功能）
func (s *UserService) ResetPassword(id uint, newPassword string) error {
	_, err := s.UserRepo.GetByID(id)
	if err != nil {
		return errors.New("用户不存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.UserRepo.UpdatePassword(id, string(hashedPassword))
}
