我将为您实施以下计划，以满足您的需求：

1.  **后端修改 (Backend)**:
    *   **移除用户表中的 `RoleID`**: 修改 `backend/internal/model/user.go`，删除 `RoleID` 和 `Role` 字段，因为我们将转而使用 Casbin 的 `g` 策略（用户-角色分组）来管理角色。
    *   **更新 UserService**:
        *   修改 `CreateUser` 和 `UpdateUser`，移除对 `RoleID` 的处理。
        *   添加 `AddRoleForUser(userID uint, roleIdent string)` 方法：使用 Casbin API (`AddGroupingPolicy`) 将用户绑定到角色。
        *   添加 `RemoveRoleForUser(userID uint, roleIdent string)` 方法。
        *   添加 `GetUserRoles(userID uint)` 方法：查询 Casbin 获取用户拥有的角色。
    *   **更新 UserHandler**:
        *   添加 `AssignRole` 接口：处理分配角色的请求（POST `/users/:id/roles`）。
        *   添加 `GetUserRoles` 接口：获取用户的角色列表。
    *   **更新 Middleware**: 修改 Casbin 中间件，使其支持多角色鉴权（如果之前是基于 `RoleID` 的单角色逻辑，现在需要适配 Casbin 的 `g` 策略）。

2.  **前端修改 (Frontend)**:
    *   **修改 UserManageView.vue**:
        *   **移除**: 在新增/编辑用户弹窗中，移除“角色选择”下拉框。
        *   **新增**: 在用户列表的操作列（编辑按钮旁边），添加“分配角色”按钮。
        *   **新增弹窗**: 点击“分配角色”按钮后，弹出一个新窗口，列出所有可用角色（复选框或穿梭框），允许为该用户勾选一个或多个角色。
        *   **API 调用**: 对接后端新的分配角色接口。

3.  **验证**:
    *   验证新增用户时不再需要选择角色。
    *   验证可以通过新按钮为用户分配角色，且分配后权限生效。

请确认是否执行此计划。注意：这涉及数据库结构的变更（删除 `role_id` 列）和鉴权逻辑的重大调整。