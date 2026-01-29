I will update the `d:\ai\3\backend\docs\routes.csv` file with the correct route information extracted from `d:\ai\3\backend\api\router\router.go`.

The new content will be:

```csv
API路径,HTTP方法,处理函数,备注
/api/v1/login,POST,userHandler.Login,用户登录
/api/v1/captcha,GET,GenerateCaptcha,验证码接口
/api/v1/logout,POST,userHandler.Logout,退出登录
/api/v1/users,POST,userHandler.CreateUser,创建用户
/api/v1/users,GET,userHandler.GetUsers,获取用户列表
/api/v1/users/:id,GET,userHandler.GetUserInfo,获取用户信息
/api/v1/users/:id,PUT,userHandler.UpdateUser,更新用户信息
/api/v1/users/:id/status,PUT,userHandler.UpdateUserStatus,更新用户状态
/api/v1/users/:id,DELETE,userHandler.DeleteUser,删除用户
/api/v1/users/change-password,PUT,userHandler.ChangePassword,修改密码
/api/v1/users/:id/reset-password,PUT,userHandler.ResetPassword,重置密码
/api/v1/users-roles/:username,POST,userHandler.AssignRole,为用户分配角色
/api/v1/users-roles/:username,DELETE,userHandler.RemoveRole,移除用户的角色
/api/v1/users-roles/:username,GET,userHandler.GetUserRoles,获取用户的角色列表
/api/v1/roles,POST,roleHandler.CreateRole,创建角色
/api/v1/roles,GET,roleHandler.GetRoles,获取角色列表
/api/v1/roles/:id,GET,roleHandler.GetRole,获取角色详情
/api/v1/roles/:id,PUT,roleHandler.UpdateRole,更新角色
/api/v1/roles/:id,DELETE,roleHandler.DeleteRole,删除角色
/api/v1/menus,POST,menuHandler.CreateMenu,创建菜单
/api/v1/menus,GET,menuHandler.GetUserMenus,查询用户可见菜单（包含子菜单）
/api/v1/menus/all,GET,menuHandler.GetAllMenus,查询所有菜单（包含子菜单）
/api/v1/menus/:id,PUT,menuHandler.UpdateMenu,更新菜单
/api/v1/menus/:id,DELETE,menuHandler.DeleteMenu,删除菜单
/api/v1/roles/:id/menus,POST,roleMenuHandler.AssignMenuToRole,为角色分配菜单权限
/api/v1/roles/:id/menus,GET,roleMenuHandler.GetRoleMenus,获取角色的菜单权限
/api/v1/roles/:id/menus,DELETE,roleMenuHandler.RemoveMenuFromRole,移除角色的菜单权限
/api/v1/roles/:id/policies,POST,permissionHandler.AddPolicy,添加Casbin策略
/api/v1/roles/:id/policies,DELETE,permissionHandler.RemovePolicy,移除Casbin策略
/api/v1/roles/:id/policies,GET,permissionHandler.GetPolicies,获取角色的Casbin策略
/api/v1/permissions,POST,permissionHandler.CreatePermission,创建权限
/api/v1/permissions,GET,permissionHandler.GetPermissions,分页查询，获取权限列表
/api/v1/permissions/all,GET,permissionHandler.GetAllPermissions,不进行分页查询，获取所有权限
/api/v1/permissions/:id,GET,permissionHandler.GetPermission,获取权限详情
/api/v1/permissions/:id,PUT,permissionHandler.UpdatePermission,更新权限
/api/v1/permissions/:id/status,PUT,permissionHandler.UpdatePermissionStatus,更新权限状态
/api/v1/permissions/:id,DELETE,permissionHandler.DeletePermission,删除权限
```