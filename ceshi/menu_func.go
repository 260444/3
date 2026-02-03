package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

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
	resp, err := makeRequest("GET", "/menus", nil, getAuthHeaders())
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
	resp, err := makeRequest("GET", "/menus/all", nil, getAuthHeaders())
	if err != nil {
		printResult("获取所有菜单", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取所有菜单", passed, resp, nil)
}

// 测试获取菜单详情
func testGetMenuByID() {
	resp, err := makeRequest("GET", fmt.Sprintf("/menus/%d", testData.MenuID), nil, getAuthHeaders())
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