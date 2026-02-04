package main

import (
	"fmt"
	"net/http"
	"time"
)

const BaseURL = "http://localhost:8081"

// 主函数
func main() {
	fmt.Println(fmt.Sprintf("=%58s", "="))
	fmt.Println("API 测试脚本 (Go)")
	fmt.Printf("测试地址: %s\n", BaseURL)
	fmt.Println(fmt.Sprintf("=%58s", "="))

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

// 检查服务器是否运行
func checkServer() bool {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(BaseURL + "/api/v1/captcha")
	if err != nil {
		// 如果没有健康检查端点，尝试访问根路径
		resp, err = client.Get(BaseURL)
		if err != nil {
			return false
		}
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// 运行所有测试
func runTests() {
	// 公开接口测试
	testCaptcha()
	testLogin()
	testLogout()

	// 用户管理测试
	testCreateUser()
	testGetUsers()
	testGetUserInfo()
	testGetCurrentUser()
	testUpdateUser()
	testUpdateUserStatus()
	testChangePassword()
	testResetPassword()
	testGetUserRoles()
	testAssignRole()
	testRemoveRole()

	// 角色管理测试
	testCreateRole()
	testGetRoles()
	testGetRole()
	testUpdateRole()
	testDeleteRole()

	// 菜单管理测试
	testCreateMenu()
	testGetUserMenus()
	testGetAllMenus()
	testUpdateMenu()
	testDeleteMenu()

	// 权限管理测试
	testCreatePermission()
	testGetPermissions()
	testGetAllPermissions()
	testGetPermission()
	testUpdatePermission()
	testUpdatePermissionStatus()
	testDeletePermission()

	// 角色菜单关联测试
	testAssignMenuToRole()
	testGetRoleMenus()
	testRemoveMenuFromRole()

	// 权限策略管理测试
	testAddPolicy()
	testGetPolicies()
	testRemovePolicy()

	// 操作日志测试
	testGetOperationLogs()
	testDeleteOperationLog()
}

// 测试结果统计
var (
	totalTests  = 0
	passedTests = 0
	failedTests = 0
	testResults = make([]string, 0)
)

// 添加测试结果
func addTestResult(testName string, passed bool, message string) {
	totalTests++
	if passed {
		passedTests++
		testResults = append(testResults, fmt.Sprintf("✓ %s: %s", testName, message))
	} else {
		failedTests++
		testResults = append(testResults, fmt.Sprintf("✗ %s: %s", testName, message))
	}
}

// 打印测试摘要
func printSummary() {
	fmt.Println(fmt.Sprintf("\n=%58s", "="))
	fmt.Println("测试结果汇总")
	fmt.Println(fmt.Sprintf("=%58s", "="))

	for _, result := range testResults {
		fmt.Println(result)
	}

	fmt.Println(fmt.Sprintf("=%58s", "="))
	fmt.Printf("总计: %d | 通过: %d | 失败: %d\n", totalTests, passedTests, failedTests)

	if failedTests == 0 {
		fmt.Println("🎉 所有测试通过!")
	} else {
		fmt.Printf("❌ %d 个测试失败\n", failedTests)
	}
}
