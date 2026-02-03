package main

// TestData 测试数据结构
type TestData struct {
	Token        string
	UserID       uint
	RoleID       uint
	MenuID       uint
	MenuIDs      []uint
	LogID        uint
	Username     string
	Password     string
	TestUsername string
	TestPassword string
}

// APIResponse API响应结构
type APIResponse struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

// PaginatedResponse 分页响应结构
type PaginatedResponse struct {
	Message string `json:"message"`
	Data    struct {
		List     interface{} `json:"list"`
		Total    int64       `json:"total"`
		Page     int         `json:"page"`
		PageSize int         `json:"page_size"`
	} `json:"data"`
}

// TestStats 测试统计结构
type TestStats struct {
	Total  int
	Passed int
	Failed int
}