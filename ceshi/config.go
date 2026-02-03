package main

import (
	"net/http"
	"time"
)

// 配置常量
const (
	BaseURL = "http://localhost:8080"
	APIBase = BaseURL + "/api/v1"
	Timeout = 10 * time.Second
)

// HTTP 客户端
var httpClient = &http.Client{
	Timeout: Timeout,
}

// 全局测试数据
var testData = &TestData{
	// 管理员账号（需要预先在数据库中创建）
	Username:     "admin",
	Password:     "Zhy20250730!",
	TestUsername: "",
	TestPassword: "Test123456",
	MenuIDs:      []uint{116, 117},
}

var stats TestStats