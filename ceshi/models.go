package main

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// APIResponse 通用API响应结构
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// User 用户结构
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
	RoleID   uint   `json:"role_id"`
}

// Role 角色结构
type Role struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

// Menu 菜单结构
type Menu struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Path      string `json:"path"`
	Component string `json:"component"`
	ParentID  *uint  `json:"parent_id"`
	Icon      string `json:"icon"`
	Sort      int    `json:"sort"`
	IsHidden  bool   `json:"is_hidden"`
	Status    int    `json:"status"`
}

// Permission 权限结构
type Permission struct {
	ID          uint   `json:"id"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
}
