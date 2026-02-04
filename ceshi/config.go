package main

// 测试配置常量
const (
	TestUsername = "testuser"
	TestPassword = "testpass123"
	TestEmail    = "test@example.com"
	TestPhone    = "13800138000"

	AdminUsername = "admin"
	AdminPassword = "123456"
)

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

// TestStats 测试统计结构
type TestStats struct {
	Total  int
	Passed int
	Failed int
}
