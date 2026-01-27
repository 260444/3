我已了解您的需求。您已经在数据库 `Role` 表中添加了 `ident` 列，现在需要同步更新后端代码和前端页面。

以下是详细的实施计划：

### 后端修改 (Backend)

1.  **修改模型文件** `backend/internal/model/role.go`:
    *   在 `Role` 结构体中添加 `Ident` 字段。
    *   添加相应的 GORM 标签（`uniqueIndex`, `size`, `not null`）和 JSON 标签。

### 前端修改 (Frontend)

1.  **修改角色管理页面** `frontend/src/views/role/RoleManageView.vue`:
    *   **列表展示**: 在表格中新增一列 "角色标识" (`ident`)，放在 "ID" 列之后。
    *   **表单数据**: 在 `roleForm` 对象中初始化 `ident` 字段。
    *   **编辑/新增弹窗**: 在表单中添加 "角色标识" 的输入框。
    *   **表单验证**: 在 `roleRules` 中添加 `ident` 的验证规则（必填）。

### 验证计划

1.  **后端**: 编译并通过。
2.  **前端**: 运行前端项目，检查角色列表是否显示标识列，检查新增/编辑角色时是否可以输入和保存标识。

请确认是否执行此计划。