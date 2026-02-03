package main

import (
	"fmt"
)

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
	//testGetMenuByID()
	//testUpdateMenu()

	// 权限管理测试
	//testAssignMenuToRole()
	//testGetRoleMenus()
	//testAddPolicy()
	//testGetPolicies()
	//testGetAllPolicies()
	//testRemovePolicy()
	//testRemoveMenuFromRole()

	// 操作日志测试
	//testGetOperationLogs()

	// 清理测试数据（可选）
	//testDeleteMenu()
	//testDeleteRole()
	//testDeleteUser()
}