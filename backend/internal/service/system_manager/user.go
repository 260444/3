// Package system_manager 提供系统管理相关的业务逻辑服务。
//
// 该包包含用户服务、角色服务、菜单服务、权限服务等核心业务逻辑的实现。
// 所有服务都依赖对应的 Repository 层进行数据访问。
package system_manager

import (
	"backend/pkg/casbin"
	"backend/pkg/logger"
	"errors"
	"time"

	"go.uber.org/zap"

	sysModel "backend/internal/model/system_manager"
	sysRepository "backend/internal/repository/system_manager"
	"golang.org/x/crypto/bcrypt"
)

// UserService 提供用户相关的业务逻辑处理。
//
// 该服务负责用户的创建、登录、信息管理、密码修改、角色分配等功能。
// 使用 bcrypt 对密码进行加密存储。
type UserService struct {
	UserRepo *sysRepository.UserRepository
	RoleRepo *sysRepository.RoleRepository
}

// NewUserService 创建一个新的 UserService 实例。
//
// 参数：
//   - userRepo: 用户数据访问层实例
//   - roleRepo: 角色数据访问层实例
//
// 返回：
//   - *UserService: 用户服务实例
func NewUserService(userRepo *sysRepository.UserRepository, roleRepo *sysRepository.RoleRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
		RoleRepo: roleRepo,
	}
}

// CreateUser 创建新用户。
//
// 该函数会验证用户名和邮箱的唯一性，并对密码进行 bcrypt 加密后存储。
// 新创建的用户默认状态为正常（Status = 1）。
//
// 参数：
//   - username: 用户名（必须唯一）
//   - password: 密码（将被加密存储）
//   - email: 邮箱（可选，必须唯一）
//   - nickname: 昵称（可选）
//
// 返回：
//   - *sysModel.User: 创建成功的用户实例
//   - error: 如果用户名或邮箱已存在，返回相应错误
func (s *UserService) CreateUser(username, password, email, nickname string) (*sysModel.User, error) {
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

	user := &sysModel.User{

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

// Login 验证用户登录凭证并返回用户信息。
//
// 该函数会验证用户名和密码，检查用户状态，并更新最后登录时间。
// 密码验证使用 bcrypt 进行比对。
//
// 参数：
//   - username: 用户名
//   - password: 明文密码
//
// 返回：
//   - *sysModel.UserWithRoleInfo: 包含用户信息和角色标识符的结构体
//   - error: 如果用户名或密码错误、用户被禁用等情况下返回错误
func (s *UserService) Login(username, password string) (*sysModel.UserWithRoleInfo, error) {
	UserWithRole, err := s.UserRepo.UserWithRoleInfo(username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	user := UserWithRole.User
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
	_ = s.UserRepo.Update(&user)

	return UserWithRole, nil
}

// GetUserByID 根据用户 ID 获取用户信息。
//
// 参数：
//   - id: 用户 ID
//
// 返回：
//   - *sysModel.User: 用户实例
//   - error: 如果用户不存在，返回错误
func (s *UserService) GetUserByID(id uint) (*sysModel.User, error) {
	return s.UserRepo.GetByID(id)
}

// UpdateUser 更新用户的基本信息。
//
// 该函数只更新用户的基本信息（用户名、邮箱、昵称、手机号、头像），
// 不更新密码和状态。如需更新密码，请使用 ChangePassword 方法。
//
// 参数：
//   - id: 用户 ID
//   - user: 包含更新信息的用户实例
//
// 返回：
//   - error: 如果用户不存在或更新失败，返回错误
func (s *UserService) UpdateUser(id uint, user *sysModel.User) error {
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

// ChangePassword 修改用户密码。
//
// 该函数会验证原密码的正确性，然后将新密码加密后存储。
//
// 参数：
//   - id: 用户 ID
//   - oldPassword: 原密码（明文）
//   - newPassword: 新密码（明文）
//
// 返回：
//   - error: 如果用户不存在、原密码错误或加密失败，返回错误
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

// AddRoleForUser 为用户分配角色。
//
// 该函数会在 Casbin 中添加分组策略，并同步更新用户的 RoleID 字段。
//
// 参数：
//   - username: 用户名
//   - roleIdent: 角色标识符
//
// 返回：
//   - error: 如果用户不存在或添加策略失败，返回错误
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

// RemoveRoleForUser 移除用户的角色。
//
// 该函数会从 Casbin 中删除分组策略，并同步清空用户的 RoleID 字段。
//
// 参数：
//   - username: 用户名
//   - roleIdent: 角色标识符
//
// 返回：
//   - error: 如果用户不存在或删除策略失败，返回错误
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

// GetUserRoles 获取用户的角色列表。
//
// 该函数从 Casbin 中查询用户所属的所有角色。
//
// 参数：
//   - username: 用户名
//
// 返回：
//   - []string: 角色标识符列表
//   - error: 如果用户不存在或查询失败，返回错误
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

// UpdateUserStatus 更新用户状态。
//
// 用于启用或禁用用户账号。
//
// 参数：
//   - id: 用户 ID
//   - status: 状态值（1-正常，0-禁用）
//
// 返回：
//   - error: 如果更新失败，返回错误
func (s *UserService) UpdateUserStatus(id uint, status int) error {
	return s.UserRepo.UpdateStatus(id, status)
}

// DeleteUser 删除用户。
//
// 该函数会软删除用户（通过 GORM 的 DeletedAt 字段）。
//
// 参数：
//   - id: 用户 ID
//
// 返回：
//   - error: 如果删除失败，返回错误
func (s *UserService) DeleteUser(id uint) error {
	return s.UserRepo.Delete(id)
}

// GetUsers 获取用户列表（分页）。
//
// 该函数支持分页查询用户列表，并返回总记录数。
//
// 参数：
//   - limit: 每页记录数
//   - offset: 偏移量（用于分页）
//
// 返回：
//   - []sysModel.User: 用户列表
//   - int64: 总记录数
//   - error: 如果查询失败，返回错误
func (s *UserService) GetUsers(limit, offset int) ([]sysModel.User, int64, error) {
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

// ResetPassword 重置用户密码（管理员功能）。
//
// 该函数允许管理员直接重置用户密码，无需验证原密码。
// 新密码会被加密后存储。
//
// 参数：
//   - id: 用户 ID
//   - newPassword: 新密码（明文）
//
// 返回：
//   - error: 如果用户不存在或加密失败，返回错误
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
