package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 测试创建用户
func testCreateUser() {
	fmt.Println("\n--- 测试创建用户 ---")

	userData := map[string]interface{}{
		"username": TestUsername,
		"password": TestPassword,
		"email":    TestEmail,
		"phone":    TestPhone,
		"nickname": "测试用户",
		"status":   1,
		"role_id":  2,
	}

	resp, body, err := sendAuthenticatedPost(BaseURL+"/api/v1/users", userData)
	if err != nil {
		addTestResult("创建用户", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		addTestResult("创建用户", true, "用户创建成功")
	} else {
		addTestResult("创建用户", false, fmt.Sprintf("创建失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取用户列表
func testGetUsers() {
	fmt.Println("\n--- 测试获取用户列表 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/users?page=1&page_size=10")
	if err != nil {
		addTestResult("获取用户列表", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取用户列表", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取用户列表", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取用户列表", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取用户详情
func testGetUserInfo() {
	fmt.Println("\n--- 测试获取用户详情 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/users/1")
	if err != nil {
		addTestResult("获取用户详情", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取用户详情", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取用户详情", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取用户详情", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取当前用户信息
func testGetCurrentUser() {
	fmt.Println("\n--- 测试获取当前用户信息 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/users/profile")
	if err != nil {
		addTestResult("获取当前用户信息", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取当前用户信息", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取当前用户信息", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取当前用户信息", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试更新用户
func testUpdateUser() {
	fmt.Println("\n--- 测试更新用户 ---")

	updateData := map[string]interface{}{
		"nickname": "更新的昵称",
		"email":    "updated@example.com",
	}

	resp, body, err := sendAuthenticatedPut(BaseURL+"/api/v1/users/1", updateData)
	if err != nil {
		addTestResult("更新用户", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("更新用户", true, "用户更新成功")
	} else {
		addTestResult("更新用户", false, fmt.Sprintf("更新失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试更新用户状态
func testUpdateUserStatus() {
	fmt.Println("\n--- 测试更新用户状态 ---")

	statusData := map[string]interface{}{
		"status": 0,
	}

	resp, body, err := sendAuthenticatedPut(BaseURL+"/api/v1/users/1/status", statusData)
	if err != nil {
		addTestResult("更新用户状态", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("更新用户状态", true, "用户状态更新成功")
	} else {
		addTestResult("更新用户状态", false, fmt.Sprintf("更新失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试修改密码
func testChangePassword() {
	fmt.Println("\n--- 测试修改密码 ---")

	passwordData := map[string]interface{}{
		"old_password": TestPassword,
		"new_password": "newpass123",
	}

	resp, body, err := sendAuthenticatedPut(BaseURL+"/api/v1/users/change-password", passwordData)
	if err != nil {
		addTestResult("修改密码", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("修改密码", true, "密码修改成功")
	} else {
		addTestResult("修改密码", false, fmt.Sprintf("修改失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试重置密码
func testResetPassword() {
	fmt.Println("\n--- 测试重置密码 ---")

	passwordData := map[string]interface{}{
		"new_password": "resetpass123",
	}

	resp, body, err := sendAuthenticatedPut(BaseURL+"/api/v1/users/1/reset-password", passwordData)
	if err != nil {
		addTestResult("重置密码", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("重置密码", true, "密码重置成功")
	} else {
		addTestResult("重置密码", false, fmt.Sprintf("重置失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取用户角色列表
func testGetUserRoles() {
	fmt.Println("\n--- 测试获取用户角色列表 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/users-roles/admin")
	if err != nil {
		addTestResult("获取用户角色列表", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取用户角色列表", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取用户角色列表", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取用户角色列表", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试为用户分配角色
func testAssignRole() {
	fmt.Println("\n--- 测试为用户分配角色 ---")

	roleData := map[string]interface{}{
		"role_ident": "test_role",
	}

	resp, body, err := sendAuthenticatedPost(BaseURL+"/api/v1/users-roles/testuser", roleData)
	if err != nil {
		addTestResult("为用户分配角色", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("为用户分配角色", true, "角色分配成功")
	} else {
		addTestResult("为用户分配角色", false, fmt.Sprintf("分配失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试移除用户角色
func testRemoveRole() {
	fmt.Println("\n--- 测试移除用户角色 ---")

	// 移除用户角色时通常不需要请求体，或者根据实际API需求调整

	resp, body, err := sendAuthenticatedDelete(BaseURL + "/api/v1/users-roles/testuser")
	if err != nil {
		addTestResult("移除用户角色", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("移除用户角色", true, "角色移除成功")
	} else {
		addTestResult("移除用户角色", false, fmt.Sprintf("移除失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试删除用户
func testDeleteUser() {
	fmt.Println("\n--- 测试删除用户 ---")

	resp, body, err := sendAuthenticatedDelete(BaseURL + "/api/v1/users/999") // 使用不存在的ID
	if err != nil {
		addTestResult("删除用户", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	// 期望返回404或其他错误状态码
	if resp.StatusCode != http.StatusOK {
		addTestResult("删除用户", true, fmt.Sprintf("删除成功或正确处理了不存在的用户，状态码: %d", resp.StatusCode))
	} else {
		addTestResult("删除用户", false, fmt.Sprintf("意外的成功响应，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}
