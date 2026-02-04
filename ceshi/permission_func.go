package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 测试创建权限资源
func testCreatePermission() {
	fmt.Println("\n--- 测试创建权限资源 ---")

	permissionData := map[string]interface{}{
		"path":        "/api/test",
		"method":      "GET",
		"description": "测试权限",
		"status":      1,
	}

	resp, body, err := sendAuthenticatedPost(BaseURL+"/api/v1/permissions", permissionData)
	if err != nil {
		addTestResult("创建权限资源", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		addTestResult("创建权限资源", true, "权限资源创建成功")
	} else {
		addTestResult("创建权限资源", false, fmt.Sprintf("创建失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取权限资源列表
func testGetPermissions() {
	fmt.Println("\n--- 测试获取权限资源列表 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/permissions?page=1&page_size=10")
	if err != nil {
		addTestResult("获取权限资源列表", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取权限资源列表", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取权限资源列表", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取权限资源列表", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取所有权限资源
func testGetAllPermissions() {
	fmt.Println("\n--- 测试获取所有权限资源 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/permissions/all")
	if err != nil {
		addTestResult("获取所有权限资源", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取所有权限资源", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取所有权限资源", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取所有权限资源", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取权限资源详情
func testGetPermission() {
	fmt.Println("\n--- 测试获取权限资源详情 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/permissions/1")
	if err != nil {
		addTestResult("获取权限资源详情", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取权限资源详情", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取权限资源详情", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取权限资源详情", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试更新权限资源
func testUpdatePermission() {
	fmt.Println("\n--- 测试更新权限资源 ---")

	updateData := map[string]interface{}{
		"description": "更新的权限描述",
	}

	resp, body, err := sendAuthenticatedPut(BaseURL+"/api/v1/permissions/1", updateData)
	if err != nil {
		addTestResult("更新权限资源", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("更新权限资源", true, "权限资源更新成功")
	} else {
		addTestResult("更新权限资源", false, fmt.Sprintf("更新失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试更新权限状态
func testUpdatePermissionStatus() {
	fmt.Println("\n--- 测试更新权限状态 ---")

	statusData := map[string]interface{}{
		"status": 0,
	}

	resp, body, err := sendAuthenticatedPut(BaseURL+"/api/v1/permissions/1/status", statusData)
	if err != nil {
		addTestResult("更新权限状态", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("更新权限状态", true, "权限状态更新成功")
	} else {
		addTestResult("更新权限状态", false, fmt.Sprintf("更新失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试删除权限资源
func testDeletePermission() {
	fmt.Println("\n--- 测试删除权限资源 ---")

	resp, body, err := sendAuthenticatedDelete(BaseURL + "/api/v1/permissions/999") // 使用不存在的ID
	if err != nil {
		addTestResult("删除权限资源", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	// 期望返回404或其他错误状态码
	if resp.StatusCode != http.StatusOK {
		addTestResult("删除权限资源", true, fmt.Sprintf("删除成功或正确处理了不存在的权限，状态码: %d", resp.StatusCode))
	} else {
		addTestResult("删除权限资源", false, fmt.Sprintf("意外的成功响应，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试为角色添加策略
func testAddPolicy() {
	fmt.Println("\n--- 测试为角色添加策略 ---")

	policyData := map[string]interface{}{
		"path":   "/api/test",
		"method": "GET",
	}

	resp, body, err := sendAuthenticatedPost(BaseURL+"/api/v1/roles/1/policies", policyData)
	if err != nil {
		addTestResult("为角色添加策略", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("为角色添加策略", true, "策略添加成功")
	} else {
		addTestResult("为角色添加策略", false, fmt.Sprintf("添加失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取角色策略
func testGetPolicies() {
	fmt.Println("\n--- 测试获取角色策略 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/roles/1/policies")
	if err != nil {
		addTestResult("获取角色策略", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取角色策略", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取角色策略", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取角色策略", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试移除角色策略
func testRemovePolicy() {
	fmt.Println("\n--- 测试移除角色策略 ---")

	// 移除策略时通常不需要请求体，或者根据实际API需求调整

	resp, body, err := sendAuthenticatedDelete(BaseURL + "/api/v1/roles/1/policies")
	if err != nil {
		addTestResult("移除角色策略", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	// 注意：这里可能需要根据实际API调整请求体的发送方式
	if resp.StatusCode == http.StatusOK {
		addTestResult("移除角色策略", true, "策略移除成功")
	} else {
		addTestResult("移除角色策略", false, fmt.Sprintf("移除失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}
