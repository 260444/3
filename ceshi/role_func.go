package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 测试创建角色
func testCreateRole() {
	fmt.Println("\n--- 测试创建角色 ---")

	roleData := map[string]interface{}{
		"name":        "test_role",
		"description": "测试角色",
		"status":      1,
	}

	resp, body, err := sendAuthenticatedPost(BaseURL+"/api/v1/roles", roleData)
	if err != nil {
		addTestResult("创建角色", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		addTestResult("创建角色", true, "角色创建成功")
	} else {
		addTestResult("创建角色", false, fmt.Sprintf("创建失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取角色列表
func testGetRoles() {
	fmt.Println("\n--- 测试获取角色列表 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/roles?page=1&page_size=10")
	if err != nil {
		addTestResult("获取角色列表", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取角色列表", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取角色列表", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取角色列表", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取角色详情
func testGetRole() {
	fmt.Println("\n--- 测试获取角色详情 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/roles/1")
	if err != nil {
		addTestResult("获取角色详情", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取角色详情", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取角色详情", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取角色详情", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试更新角色
func testUpdateRole() {
	fmt.Println("\n--- 测试更新角色 ---")

	updateData := map[string]interface{}{
		"description": "更新的角色描述",
	}

	resp, body, err := sendAuthenticatedPut(BaseURL+"/api/v1/roles/1", updateData)
	if err != nil {
		addTestResult("更新角色", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("更新角色", true, "角色更新成功")
	} else {
		addTestResult("更新角色", false, fmt.Sprintf("更新失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试删除角色
func testDeleteRole() {
	fmt.Println("\n--- 测试删除角色 ---")

	resp, body, err := sendAuthenticatedDelete(BaseURL + "/api/v1/roles/999") // 使用不存在的ID
	if err != nil {
		addTestResult("删除角色", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	// 期望返回404或其他错误状态码
	if resp.StatusCode != http.StatusOK {
		addTestResult("删除角色", true, fmt.Sprintf("删除成功或正确处理了不存在的角色，状态码: %d", resp.StatusCode))
	} else {
		addTestResult("删除角色", false, fmt.Sprintf("意外的成功响应，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}
