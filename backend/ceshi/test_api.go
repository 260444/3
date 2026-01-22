package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// 配置常量
const (
	BaseURL = "http://localhost:8080"
	APIBase = BaseURL + "/api/v1"
	Timeout = 10 * time.Second
)

// 测试数据存储
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

// 响应结构
type APIResponse struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

// 分页响应结构
type PaginatedResponse struct {
	Message string `json:"message"`
	Data    struct {
		List     interface{} `json:"list"`
		Total    int64       `json:"total"`
		Page     int         `json:"page"`
		PageSize int         `json:"page_size"`
	} `json:"data"`
}

// 全局测试数据
var testData = &TestData{
	// 管理员账号（需要预先在数据库中创建）
	Username:     "admin",
	Password:     "Zhy20250730!",
	TestUsername: fmt.Sprintf("testuser_%d", time.Now().Unix()),
	TestPassword: "Test123456",
	MenuIDs:      []uint{116, 117},
}

// HTTP 客户端
var httpClient = &http.Client{
	Timeout: Timeout,
}

// 测试结果统计
type TestStats struct {
	Total  int
	Passed int
	Failed int
}

var stats TestStats

// 主函数
func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("API 测试脚本 (Go)")
	fmt.Printf("测试地址: %s\n", BaseURL)
	fmt.Println(strings.Repeat("=", 60))

	// 检查服务器是否运行
	if !checkServer() {
		fmt.Printf("\n✗ 无法连接到服务器 %s\n", BaseURL)
		fmt.Println("请确保后端服务已启动")
		return
	}
	fmt.Println("\n✓ 服务器运行中")

	// 执行测试
	runTests()

	// 打印测试结果汇总
	printSummary()
}

// 检查服务器状态
func checkServer() bool {
	resp, err := httpClient.Get(BaseURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode < 500
}

// 运行所有测试
func runTests() {
	// 公开接口测试
	//testCaptcha()
	testLogin()
	//testLogout()

	// 用户管理测试
	//testCreateUser()
	//testGetUsers()
	//testGetUserInfo()
	//testUpdateUser()
	//testUpdateUserStatus()
	//testChangePassword()

	// 角色管理测试
	//testCreateRole()
	//testGetRoles()
	//testGetRole()
	//testUpdateRole()

	// 菜单管理测试
	//testCreateMenu()
	//testGetMenuTree()
	//testGetAllMenus()
	//testUpdateMenu()

	// 权限管理测试
	testAssignMenuToRole()
	//testGetRoleMenus()
	//testAddPolicy()
	//testGetPolicies()
	//testGetAllPolicies()
	//testRemovePolicy()
	//testRemoveMenuFromRole()

	// 操作日志测试
	//testGetOperationLogs()

	// 清理测试数据（可选）
	// testDeleteMenu()
	// testDeleteRole()
	// testDeleteUser()
}

// 发送 HTTP 请求
func makeRequest(method, endpoint string, body interface{}, headers map[string]string) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, APIBase+endpoint, reqBody)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return httpClient.Do(req)
}

// 获取认证头
func getAuthHeaders() map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + testData.Token,
	}
}

// 打印测试结果
func printResult(testName string, passed bool, resp *http.Response, err error) {
	stats.Total++
	if passed {
		stats.Passed++
		fmt.Printf("\n✓ %s - PASS\n", testName)
	} else {
		stats.Failed++
		fmt.Printf("\n✗ %s - FAIL\n", testName)
	}

	if err != nil {
		fmt.Printf("  错误: %v\n", err)
		return
	}

	if resp != nil {
		fmt.Printf("  状态码: %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, body, "", "  ")
		fmt.Printf("  响应: %s\n", prettyJSON.String())
	}
}

// 打印测试结果汇总
func printSummary() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("测试结果汇总")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("总计: %d 个测试\n", stats.Total)
	fmt.Printf("通过: %d 个\n", stats.Passed)
	fmt.Printf("失败: %d 个\n", stats.Failed)
	if stats.Total > 0 {
		fmt.Printf("成功率: %.1f%%\n", float64(stats.Passed)/float64(stats.Total)*100)
	}
	fmt.Println(strings.Repeat("=", 60))
}

// ==================== 公开接口测试 ====================

// 测试验证码
func testCaptcha() {
	resp, err := makeRequest("GET", "/captcha", nil, nil)
	if err != nil {
		printResult("获取验证码", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取验证码", passed, resp, nil)
}

// 测试创建用户
func testCreateUser() {
	data := map[string]interface{}{
		"username": testData.TestUsername,
		"password": testData.TestPassword,
		"email":    testData.TestUsername + "@example.com",
		"nickname": "测试用户",
	}

	resp, err := makeRequest("POST", "/users", data, getAuthHeaders())
	if err != nil {
		printResult("创建用户", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := false
	if resp.StatusCode == 200 {
		var apiResp APIResponse
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &apiResp)
		if apiResp.Data != nil {
			if dataMap, ok := apiResp.Data.(map[string]interface{}); ok {
				if id, ok := dataMap["id"].(float64); ok {
					testData.UserID = uint(id)
				}
			}
		}
		passed = true
	}
	printResult("创建用户", passed, resp, nil)
}

// 测试用户登录
func testLogin() {
	data := map[string]interface{}{
		"username": testData.Username,
		"password": testData.Password,
	}

	resp, err := makeRequest("POST", "/login", data, nil)
	if err != nil {
		printResult("用户登录", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := false
	if resp.StatusCode == 200 {
		var apiResp APIResponse
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &apiResp)
		if apiResp.Data != nil {
			if dataMap, ok := apiResp.Data.(map[string]interface{}); ok {
				if token, ok := dataMap["token"].(string); ok {
					testData.Token = token
				}
				if user, ok := dataMap["user"].(map[string]interface{}); ok {
					if id, ok := user["id"].(float64); ok {
						testData.UserID = uint(id)
					}
				}
			}
		}
		passed = true
	}
	printResult("用户登录", passed, resp, nil)
}

// 测试退出登录
func testLogout() {
	resp, err := makeRequest("POST", "/logout", nil, getAuthHeaders())
	if err != nil {
		printResult("退出登录", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("退出登录", passed, resp, nil)
}

// ==================== 用户管理测试 ====================

// 测试获取用户列表
func testGetUsers() {
	resp, err := makeRequest("GET", "/users?page=1&page_size=10", nil, getAuthHeaders())
	if err != nil {
		printResult("获取用户列表", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取用户列表", passed, resp, nil)
}

// 测试获取用户信息
func testGetUserInfo() {
	resp, err := makeRequest("GET", fmt.Sprintf("/users/%d", testData.UserID), nil, getAuthHeaders())
	if err != nil {
		printResult("获取用户信息", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取用户信息", passed, resp, nil)
}

// 测试更新用户信息
func testUpdateUser() {
	data := map[string]interface{}{
		"username": testData.TestUsername,
		"email":    testData.TestUsername + "_updated@example.com",
		"nickname": "测试用户（已更新）",
	}

	resp, err := makeRequest("PUT", fmt.Sprintf("/users/%d", testData.UserID), data, getAuthHeaders())
	if err != nil {
		printResult("更新用户信息", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("更新用户信息", passed, resp, nil)
}

// 测试更新用户状态
func testUpdateUserStatus() {
	data := map[string]interface{}{
		"status": 1,
	}

	resp, err := makeRequest("PUT", fmt.Sprintf("/users/%d/status", testData.UserID), data, getAuthHeaders())
	if err != nil {
		printResult("更新用户状态", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("更新用户状态", passed, resp, nil)
}

// 测试修改密码
func testChangePassword() {
	data := map[string]interface{}{
		"old_password": testData.TestPassword,
		"new_password": "NewTest123456",
	}

	resp, err := makeRequest("PUT", "/users/change-password", data, getAuthHeaders())
	if err != nil {
		printResult("修改密码", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	if passed {
		testData.TestPassword = data["new_password"].(string)
	}
	printResult("修改密码", passed, resp, nil)
}

// ==================== 角色管理测试 ====================

// 测试创建角色
func testCreateRole() {
	data := map[string]interface{}{
		"name":        fmt.Sprintf("测试角色_%d", time.Now().Unix()),
		"description": "这是一个测试角色",
		"status":      1,
	}

	resp, err := makeRequest("POST", "/roles", data, getAuthHeaders())
	if err != nil {
		printResult("创建角色", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := false
	if resp.StatusCode == 200 {
		var apiResp APIResponse
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &apiResp)
		if apiResp.Data != nil {
			if dataMap, ok := apiResp.Data.(map[string]interface{}); ok {
				if id, ok := dataMap["id"].(float64); ok {
					testData.RoleID = uint(id)
				}
			}
		}
		passed = true
	}
	printResult("创建角色", passed, resp, nil)
}

// 测试获取角色列表
func testGetRoles() {
	resp, err := makeRequest("GET", "/roles?page=1&page_size=10", nil, getAuthHeaders())
	if err != nil {
		printResult("获取角色列表", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取角色列表", passed, resp, nil)
}

// 测试获取角色详情
func testGetRole() {
	resp, err := makeRequest("GET", fmt.Sprintf("/roles/%d", testData.RoleID), nil, getAuthHeaders())
	if err != nil {
		printResult("获取角色详情", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取角色详情", passed, resp, nil)
}

// 测试更新角色
func testUpdateRole() {
	data := map[string]interface{}{
		"name":        fmt.Sprintf("测试角色_%d_updated", time.Now().Unix()),
		"description": "这是一个更新后的测试角色",
		"status":      1,
	}

	resp, err := makeRequest("PUT", fmt.Sprintf("/roles/%d", testData.RoleID), data, getAuthHeaders())
	if err != nil {
		printResult("更新角色", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("更新角色", passed, resp, nil)
}

// ==================== 菜单管理测试 ====================

// 测试创建菜单
func testCreateMenu() {
	data := map[string]interface{}{
		"name":      fmt.Sprintf("test_menu_%d", time.Now().Unix()),
		"title":     "测试菜单",
		"path":      "/test",
		"component": "Test",
		"icon":      "test-icon",
		"sort":      100,
		"status":    1,
	}

	resp, err := makeRequest("POST", "/menus", data, getAuthHeaders())
	if err != nil {
		printResult("创建菜单", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := false
	if resp.StatusCode == 200 {
		var apiResp APIResponse
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &apiResp)
		if apiResp.Data != nil {
			if dataMap, ok := apiResp.Data.(map[string]interface{}); ok {
				if id, ok := dataMap["id"].(float64); ok {
					testData.MenuID = uint(id)
				}
			}
		}
		passed = true
	}
	printResult("创建菜单", passed, resp, nil)
}

// 测试获取菜单树
func testGetMenuTree() {
	url := "/menus"
	fmt.Printf("请求URL: %s%s\n", APIBase, url)

	resp, err := makeRequest("GET", url, nil, getAuthHeaders())
	if err != nil {
		printResult("获取菜单树", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取菜单树", passed, resp, nil)
}

// 测试获取所有菜单
func testGetAllMenus() {
	url := "/menus/all"
	fmt.Printf("请求URL: %s%s\n", APIBase, url)

	resp, err := makeRequest("GET", url, nil, getAuthHeaders())
	if err != nil {
		printResult("获取所有菜单", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取所有菜单", passed, resp, nil)
}

// 测试获取菜单详情
func testGetMenu() {
	url := fmt.Sprintf("/menus/%d", testData.MenuID)
	fmt.Printf("请求URL: %s%s\n", APIBase, url)

	resp, err := makeRequest("GET", url, nil, getAuthHeaders())
	if err != nil {
		printResult("获取菜单详情", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取菜单详情", passed, resp, nil)
}

// 测试更新菜单
func testUpdateMenu() {
	data := map[string]interface{}{
		"name":      fmt.Sprintf("test_menu_%d_updated", time.Now().Unix()),
		"title":     "测试菜单（已更新）",
		"path":      "/test-updated",
		"component": "Test",
		"icon":      "test-icon-updated",
		"sort":      101,
		"status":    1,
	}

	resp, err := makeRequest("PUT", fmt.Sprintf("/menus/%d", testData.MenuID), data, getAuthHeaders())
	if err != nil {
		printResult("更新菜单", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("更新菜单", passed, resp, nil)
}

// ==================== 权限管理测试 ====================

// 测试为角色分配菜单权限
func testAssignMenuToRole() {
	data := map[string]interface{}{
		"menu_ids": []uint{testData.MenuIDs[0]},
	}

	resp, err := makeRequest("POST", fmt.Sprintf("/roles/%d/menus", testData.RoleID), data, getAuthHeaders())
	if err != nil {
		printResult("为角色分配菜单权限", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("为角色分配菜单权限", passed, resp, nil)
}

// 测试获取角色菜单权限
func testGetRoleMenus() {
	resp, err := makeRequest("GET", fmt.Sprintf("/roles/%d/menus", testData.RoleID), nil, getAuthHeaders())
	if err != nil {
		printResult("获取角色菜单权限", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取角色菜单权限", passed, resp, nil)
}

// 测试添加 Casbin 策略
func testAddPolicy() {
	data := map[string]interface{}{
		"path":   "/api/v1/users",
		"method": "GET",
	}

	resp, err := makeRequest("POST", fmt.Sprintf("/roles/%d/policies", testData.RoleID), data, getAuthHeaders())
	if err != nil {
		printResult("添加 Casbin 策略", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("添加 Casbin 策略", passed, resp, nil)
}

// 测试获取角色 Casbin 策略
func testGetPolicies() {
	resp, err := makeRequest("GET", fmt.Sprintf("/roles/%d/policies", testData.RoleID), nil, getAuthHeaders())
	if err != nil {
		printResult("获取角色 Casbin 策略", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取角色 Casbin 策略", passed, resp, nil)
}

// 测试获取所有 Casbin 策略
func testGetAllPolicies() {
	resp, err := makeRequest("GET", "/policies", nil, getAuthHeaders())
	if err != nil {
		printResult("获取所有 Casbin 策略", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取所有 Casbin 策略", passed, resp, nil)
}

// 测试移除 Casbin 策略
func testRemovePolicy() {
	data := map[string]interface{}{
		"path":   "/api/v1/users",
		"method": "GET",
	}

	resp, err := makeRequest("DELETE", fmt.Sprintf("/roles/%d/policies", testData.RoleID), data, getAuthHeaders())
	if err != nil {
		printResult("移除 Casbin 策略", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("移除 Casbin 策略", passed, resp, nil)
}

// 测试移除角色菜单权限
func testRemoveMenuFromRole() {
	data := map[string]interface{}{
		"menu_ids": []uint{testData.MenuID},
	}

	resp, err := makeRequest("DELETE", fmt.Sprintf("/roles/%d/menus", testData.RoleID), data, getAuthHeaders())
	if err != nil {
		printResult("移除角色菜单权限", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("移除角色菜单权限", passed, resp, nil)
}

// ==================== 操作日志测试 ====================

// 测试获取操作日志
func testGetOperationLogs() {
	resp, err := makeRequest("GET", "/operation-logs?page=1&page_size=10", nil, getAuthHeaders())
	if err != nil {
		printResult("获取操作日志", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取操作日志", passed, resp, nil)
}

// ==================== 清理测试数据（可选） ====================

// 测试删除菜单
func testDeleteMenu() {
	resp, err := makeRequest("DELETE", fmt.Sprintf("/menus/%d", testData.MenuID), nil, getAuthHeaders())
	if err != nil {
		printResult("删除菜单", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("删除菜单", passed, resp, nil)
}

// 测试删除角色
func testDeleteRole() {
	resp, err := makeRequest("DELETE", fmt.Sprintf("/roles/%d", testData.RoleID), nil, getAuthHeaders())
	if err != nil {
		printResult("删除角色", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("删除角色", passed, resp, nil)
}

// 测试删除用户
func testDeleteUser() {
	resp, err := makeRequest("DELETE", fmt.Sprintf("/users/%d", testData.UserID), nil, getAuthHeaders())
	if err != nil {
		printResult("删除用户", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("删除用户", passed, resp, nil)
}
