package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/pkg/casbin"
	"errors"
	"fmt"
	"time"

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

// CreateUser 创建用户
func (s *UserService) CreateUser(username, password, email, nickname string) (*model.User, error) {
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
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Nickname: nickname,
		Status:   1, // 默认启用
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login 用户登录
func (s *UserService) Login(username, password string) (*model.User, error) {
	user, err := s.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if user.Status == 0 {
		return nil, errors.New("用户已被禁用")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 更新最后登录时间和IP
	now := time.Now()
	user.LastLoginAt = &now
	// 注意：在实际应用中需要从请求中获取IP
	// user.LastLoginIP = 获取IP的逻辑
	_ = s.UserRepo.Update(user)

	return user, nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.UserRepo.GetByID(id)
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(id uint, user *model.User) error {
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

// AddRoleForUser 为用户分配角色
func (s *UserService) AddRoleForUser(userID uint, roleIdent string) error {
	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 构造 Casbin 的 subject，例如 user_1
	sub := fmt.Sprintf("user_%d", user.ID)
	// 将用户添加到角色组 (g policy)
	// g, user_1, role_admin
	_, err = casbin.Enforcer.AddGroupingPolicy(sub, roleIdent)
	return err
}

// RemoveRoleForUser 移除用户的角色
func (s *UserService) RemoveRoleForUser(userID uint, roleIdent string) error {
	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	sub := fmt.Sprintf("user_%d", user.ID)
	_, err = casbin.Enforcer.RemoveGroupingPolicy(sub, roleIdent)
	return err
}

// GetUserRoles 获取用户的角色列表
func (s *UserService) GetUserRoles(userID uint) ([]string, error) {
	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	sub := fmt.Sprintf("user_%d", user.ID)
	// 获取用户的所有角色 (g policy 中的第二项)
	roles, err := casbin.Enforcer.GetRolesForUser(sub)
	return roles, err
}

// UpdateUserStatus 更新用户状态
func (s *UserService) UpdateUserStatus(id uint, status int) error {
	return s.UserRepo.UpdateStatus(id, status)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id uint) error {
	return s.UserRepo.Delete(id)
}

// GetUsers 获取用户列表
func (s *UserService) GetUsers(limit, offset int) ([]model.User, int64, error) {
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
