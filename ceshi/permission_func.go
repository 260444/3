package main

import (
	"fmt"
)

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