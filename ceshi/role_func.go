package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

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