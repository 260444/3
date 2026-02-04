package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 测试创建菜单
func testCreateMenu() {
	fmt.Println("\n--- 测试创建菜单 ---")

	menuData := map[string]interface{}{
		"name":      "test_menu",
		"title":     "测试菜单",
		"path":      "/test",
		"component": "TestView",
		"icon":      "el-icon-setting",
		"sort":      99,
		"is_hidden": false,
		"status":    1,
	}

	resp, body, err := sendAuthenticatedPost(BaseURL+"/api/v1/menus", menuData)
	if err != nil {
		addTestResult("创建菜单", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		addTestResult("创建菜单", true, "菜单创建成功")
	} else {
		addTestResult("创建菜单", false, fmt.Sprintf("创建失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取用户菜单
func testGetUserMenus() {
	fmt.Println("\n--- 测试获取用户菜单 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/menus")
	if err != nil {
		addTestResult("获取用户菜单", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取用户菜单", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取用户菜单", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取用户菜单", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取所有菜单
func testGetAllMenus() {
	fmt.Println("\n--- 测试获取所有菜单 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/menus/all")
	if err != nil {
		addTestResult("获取所有菜单", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取所有菜单", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取所有菜单", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取所有菜单", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试更新菜单
func testUpdateMenu() {
	fmt.Println("\n--- 测试更新菜单 ---")

	updateData := map[string]interface{}{
		"title": "更新的菜单标题",
		"sort":  100,
	}

	resp, body, err := sendAuthenticatedPut(BaseURL+"/api/v1/menus/1", updateData)
	if err != nil {
		addTestResult("更新菜单", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("更新菜单", true, "菜单更新成功")
	} else {
		addTestResult("更新菜单", false, fmt.Sprintf("更新失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试删除菜单
func testDeleteMenu() {
	fmt.Println("\n--- 测试删除菜单 ---")

	resp, body, err := sendAuthenticatedDelete(BaseURL + "/api/v1/menus/999") // 使用不存在的ID
	if err != nil {
		addTestResult("删除菜单", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	// 期望返回404或其他错误状态码
	if resp.StatusCode != http.StatusOK {
		addTestResult("删除菜单", true, fmt.Sprintf("删除成功或正确处理了不存在的菜单，状态码: %d", resp.StatusCode))
	} else {
		addTestResult("删除菜单", false, fmt.Sprintf("意外的成功响应，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}
