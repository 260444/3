package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 测试为角色分配菜单权限
func testAssignMenuToRole() {
	fmt.Println("\n--- 测试为角色分配菜单权限 ---")

	menuData := map[string]interface{}{
		"menu_ids": []uint{1, 2, 3},
	}

	resp, body, err := sendAuthenticatedPost(BaseURL+"/api/v1/roles/1/menus", menuData)
	if err != nil {
		addTestResult("为角色分配菜单权限", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("为角色分配菜单权限", true, "菜单权限分配成功")
	} else {
		addTestResult("为角色分配菜单权限", false, fmt.Sprintf("分配失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试获取角色菜单权限
func testGetRoleMenus() {
	fmt.Println("\n--- 测试获取角色菜单权限 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/roles/1/menus")
	if err != nil {
		addTestResult("获取角色菜单权限", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取角色菜单权限", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取角色菜单权限", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取角色菜单权限", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试移除角色菜单权限
func testRemoveMenuFromRole() {
	fmt.Println("\n--- 测试移除角色菜单权限 ---")

	// 移除菜单权限时通常不需要请求体，或者根据实际API需求调整

	resp, body, err := sendAuthenticatedDelete(BaseURL + "/api/v1/roles/1/menus")
	if err != nil {
		addTestResult("移除角色菜单权限", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		addTestResult("移除角色菜单权限", true, "菜单权限移除成功")
	} else {
		addTestResult("移除角色菜单权限", false, fmt.Sprintf("移除失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}
