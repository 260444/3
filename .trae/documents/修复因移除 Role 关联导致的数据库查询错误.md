我已查明错误原因。在之前的修改中，我移除了 `User` 模型中的 `RoleID` 和 `Role` 字段，但是 `UserRepository` 中的查询方法（`GetByID`, `GetByUsername`, `GetByEmail`, `List`）仍然包含 `.Preload("Role")`，导致 GORM 报错 `Role: unsupported relations for schema User`。

此外，`Role` 模型中仍然保留了 `Users []User` 字段和 `foreignKey:RoleID` 标签，这也会导致问题，因为 `User` 表中已经没有 `role_id` 列了。

我将执行以下修复计划：

1.  **修复 `Role` 模型 (Backend)**:
    *   修改 `backend/internal/model/role.go`，移除 `Users` 字段。

2.  **修复 `UserRepository` (Backend)**:
    *   修改 `backend/internal/repository/user.go`。
    *   移除所有方法中的 `.Preload("Role")` 调用。

3.  **验证**:
    *   重新编译后端并启动。
    *   验证获取用户列表接口是否恢复正常。

请确认执行此修复。