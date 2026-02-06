// Package di 依赖注入容器
package di

import (
	sysHandler "backend/api/handler/system_manager"
	sysRepo "backend/internal/repository/system_manager"
	sysService "backend/internal/service/system_manager"

	assHandler "backend/api/handler/asset_management"
	assRepo "backend/internal/repository/asset_management"
	assService "backend/internal/service/asset_management"

	"backend/pkg/casbin"
	"backend/pkg/database"
	myredis "backend/pkg/redis"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// Container 依赖注入容器接口
type Container interface {
	GetUserRepository() *sysRepo.UserRepository
	GetRoleRepository() *sysRepo.RoleRepository
	GetMenuRepository() *sysRepo.MenuRepository
	GetOperationLogRepository() *sysRepo.OperationLogRepository
	GetPermissionRepository() *sysRepo.PermissionRepository
	GetRoleMenuRepository() *sysRepo.RoleMenuRepository
	GetHostRepository() *assRepo.HostRepository
	GetHostGroupRepository() *assRepo.HostGroupRepository
	GetHostMetricRepository() *assRepo.HostMetricRepository

	GetUserService() *sysService.UserService
	GetRoleService() *sysService.RoleService
	GetMenuService() *sysService.MenuService
	GetOperationLogService() *sysService.OperationLogService
	GetPermissionService() *sysService.PermissionService
	GetRoleMenuService() *sysService.RoleMenuService
	GetHostService() *assService.HostService
	GetHostGroupService() *assService.HostGroupService
	GetHostMetricService() *assService.HostMetricService
	GetCredentialService() *assService.CredentialService

	GetUserHandler() *sysHandler.UserHandler
	GetRoleHandler() *sysHandler.RoleHandler
	GetMenuHandler() *sysHandler.MenuHandler
	GetOperationLogHandler() *sysHandler.OperationLogHandler
	GetPermissionHandler() *sysHandler.PermissionHandler
	GetRoleMenuHandler() *sysHandler.RoleMenuHandler
	GetHostHandler() *assHandler.HostHandler
	GetHostGroupHandler() *assHandler.HostGroupHandler
	GetCredentialHandler() *assHandler.CredentialHandler

	GetDB() *gorm.DB
	GetRedis() *redis.Client

	// Close 关闭所有资源
	Close() error
}

// containerImpl 依赖注入容器实现
type containerImpl struct {
	// 数据库连接
	db    *gorm.DB
	redis *redis.Client

	// Repositories
	userRepo         *sysRepo.UserRepository
	roleRepo         *sysRepo.RoleRepository
	menuRepo         *sysRepo.MenuRepository
	operationLogRepo *sysRepo.OperationLogRepository
	permissionRepo   *sysRepo.PermissionRepository
	roleMenuRepo     *sysRepo.RoleMenuRepository
	hostRepo         *assRepo.HostRepository
	hostGroupRepo    *assRepo.HostGroupRepository
	hostMetricRepo   *assRepo.HostMetricRepository
	credentialRepo   *assRepo.CredentialRepository

	// Services
	userService         *sysService.UserService
	roleService         *sysService.RoleService
	menuService         *sysService.MenuService
	operationLogService *sysService.OperationLogService
	permissionService   *sysService.PermissionService
	roleMenuService     *sysService.RoleMenuService
	hostService         *assService.HostService
	hostGroupService    *assService.HostGroupService
	hostMetricService   *assService.HostMetricService
	credentialService   *assService.CredentialService

	// Handlers
	userHandler         *sysHandler.UserHandler
	roleHandler         *sysHandler.RoleHandler
	menuHandler         *sysHandler.MenuHandler
	operationLogHandler *sysHandler.OperationLogHandler
	permissionHandler   *sysHandler.PermissionHandler
	roleMenuHandler     *sysHandler.RoleMenuHandler
	hostHandler         *assHandler.HostHandler
	hostGroupHandler    *assHandler.HostGroupHandler
	credentialHandler   *assHandler.CredentialHandler
}

// InitializeContainer 初始化依赖注入容器
func InitializeContainer() (Container, error) {
	// 初始化数据库
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	// 初始化Redis
	redisClient, err := myredis.InitRedis()
	if err != nil {
		// Redis不是必需的，记录警告但不中断
		// logger.Logger.Warn("Redis初始化失败", zap.Error(err))
		redisClient = nil
	}

	// 初始化Casbin
	if err := casbin.InitCasbinWithGormAdapter(db); err != nil {
		return nil, err
	}

	container := &containerImpl{
		db:    db,
		redis: redisClient,
	}

	// 初始化各层依赖
	container.initRepositories()
	container.initServices()
	container.initHandlers()

	return container, nil
}

// initRepositories 初始化Repository层
func (c *containerImpl) initRepositories() {
	c.userRepo = sysRepo.NewUserRepository(c.db)
	c.roleRepo = sysRepo.NewRoleRepository(c.db)
	c.menuRepo = sysRepo.NewMenuRepository(c.db)
	c.operationLogRepo = sysRepo.NewOperationLogRepository(c.db)
	c.permissionRepo = sysRepo.NewPermissionRepository(c.db)
	c.roleMenuRepo = sysRepo.NewRoleMenuRepository(c.db)
	c.hostRepo = assRepo.NewHostRepository(c.db)
	c.hostGroupRepo = assRepo.NewHostGroupRepository(c.db)
	c.hostMetricRepo = assRepo.NewHostMetricRepository(c.db)
	c.credentialRepo = assRepo.NewCredentialRepository(c.db)
}

// initServices 初始化Service层
func (c *containerImpl) initServices() {
	c.userService = sysService.NewUserService(c.userRepo, c.roleRepo)
	c.roleService = sysService.NewRoleService(c.roleRepo)
	c.menuService = sysService.NewMenuService(c.menuRepo)
	c.operationLogService = sysService.NewOperationLogService(c.operationLogRepo)
	c.permissionService = sysService.NewPermissionService(c.roleRepo, c.menuRepo, c.permissionRepo)
	c.roleMenuService = sysService.NewRoleMenuService(c.roleMenuRepo)
	c.credentialService = assService.NewCredentialService(c.credentialRepo)
	c.hostService = assService.NewHostService(c.hostRepo, c.hostGroupRepo)
	c.hostGroupService = assService.NewHostGroupService(c.hostGroupRepo, c.hostRepo)
	c.hostMetricService = assService.NewHostMetricService(c.hostMetricRepo, c.hostRepo)
}

// initHandlers 初始化Handler层
func (c *containerImpl) initHandlers() {
	c.userHandler = sysHandler.NewUserHandler(c.userService)
	c.roleHandler = sysHandler.NewRoleHandler(c.roleService)
	c.menuHandler = sysHandler.NewMenuHandler(c.menuService)
	c.operationLogHandler = sysHandler.NewOperationLogHandler(c.operationLogService)
	c.permissionHandler = sysHandler.NewPermissionHandler(c.permissionService)
	c.roleMenuHandler = sysHandler.NewRoleMenuHandler(c.roleMenuService)
	c.hostHandler = assHandler.NewHostHandler(c.hostService, c.hostMetricService, c.credentialService)
	c.hostGroupHandler = assHandler.NewHostGroupHandler(c.hostGroupService)
	c.credentialHandler = assHandler.NewCredentialHandler(c.credentialService)
}

// Getters for repositories
func (c *containerImpl) GetUserRepository() *sysRepo.UserRepository {
	return c.userRepo
}

func (c *containerImpl) GetRoleRepository() *sysRepo.RoleRepository {
	return c.roleRepo
}

func (c *containerImpl) GetMenuRepository() *sysRepo.MenuRepository {
	return c.menuRepo
}

func (c *containerImpl) GetOperationLogRepository() *sysRepo.OperationLogRepository {
	return c.operationLogRepo
}

func (c *containerImpl) GetPermissionRepository() *sysRepo.PermissionRepository {
	return c.permissionRepo
}

func (c *containerImpl) GetRoleMenuRepository() *sysRepo.RoleMenuRepository {
	return c.roleMenuRepo
}

func (c *containerImpl) GetHostRepository() *assRepo.HostRepository {
	return c.hostRepo
}

func (c *containerImpl) GetHostGroupRepository() *assRepo.HostGroupRepository {
	return c.hostGroupRepo
}

func (c *containerImpl) GetHostMetricRepository() *assRepo.HostMetricRepository {
	return c.hostMetricRepo
}

// Getters for services
func (c *containerImpl) GetUserService() *sysService.UserService {
	return c.userService
}

func (c *containerImpl) GetRoleService() *sysService.RoleService {
	return c.roleService
}

func (c *containerImpl) GetMenuService() *sysService.MenuService {
	return c.menuService
}

func (c *containerImpl) GetOperationLogService() *sysService.OperationLogService {
	return c.operationLogService
}

func (c *containerImpl) GetPermissionService() *sysService.PermissionService {
	return c.permissionService
}

func (c *containerImpl) GetRoleMenuService() *sysService.RoleMenuService {
	return c.roleMenuService
}

func (c *containerImpl) GetHostService() *assService.HostService {
	return c.hostService
}

func (c *containerImpl) GetHostGroupService() *assService.HostGroupService {
	return c.hostGroupService
}

func (c *containerImpl) GetHostMetricService() *assService.HostMetricService {
	return c.hostMetricService
}

func (c *containerImpl) GetCredentialService() *assService.CredentialService {
	return c.credentialService
}

// Getters for handlers
func (c *containerImpl) GetUserHandler() *sysHandler.UserHandler {
	return c.userHandler
}

func (c *containerImpl) GetRoleHandler() *sysHandler.RoleHandler {
	return c.roleHandler
}

func (c *containerImpl) GetMenuHandler() *sysHandler.MenuHandler {
	return c.menuHandler
}

func (c *containerImpl) GetOperationLogHandler() *sysHandler.OperationLogHandler {
	return c.operationLogHandler
}

func (c *containerImpl) GetPermissionHandler() *sysHandler.PermissionHandler {
	return c.permissionHandler
}

func (c *containerImpl) GetRoleMenuHandler() *sysHandler.RoleMenuHandler {
	return c.roleMenuHandler
}

func (c *containerImpl) GetHostHandler() *assHandler.HostHandler {
	return c.hostHandler
}

func (c *containerImpl) GetHostGroupHandler() *assHandler.HostGroupHandler {
	return c.hostGroupHandler
}

func (c *containerImpl) GetCredentialHandler() *assHandler.CredentialHandler {
	return c.credentialHandler
}

// Getters for infrastructure
func (c *containerImpl) GetDB() *gorm.DB {
	return c.db
}

func (c *containerImpl) GetRedis() *redis.Client {
	return c.redis
}

// Close 关闭资源
func (c *containerImpl) Close() error {
	// 关闭数据库连接
	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	// 关闭Redis连接
	if c.redis != nil {
		if err := c.redis.Close(); err != nil {
			return err
		}
	}

	return nil
}
