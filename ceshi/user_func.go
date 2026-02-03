package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// 测试创建用户
func testCreateUser() {
	data := map[string]interface{}{
		"username": testData.TestUsername,
		"password": testData.TestPassword,
		"email":    testData.TestUsername + "@example.com",
		"nickname": "测试用户",
	}

	resp, err := makeRequest("POST", "/users", data, getAuthHeaders())
	if err != nil {
		printResult("创建用户", false, nil, err)
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
					testData.UserID = uint(id)
				}
			}
		}
		passed = true
	}
	printResult("创建用户", passed, resp, nil)
}

// 测试获取用户列表
func testGetUsers() {
	resp, err := makeRequest("GET", "/users?page=1&page_size=10", nil, getAuthHeaders())
	if err != nil {
		printResult("获取用户列表", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取用户列表", passed, resp, nil)
}

// 测试获取用户信息
func testGetUserInfo() {
	resp, err := makeRequest("GET", fmt.Sprintf("/users/%d", testData.UserID), nil, getAuthHeaders())
	if err != nil {
		printResult("获取用户信息", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取用户信息", passed, resp, nil)
}

// 测试更新用户信息
func testUpdateUser() {
	data := map[string]interface{}{
		"username": testData.TestUsername,
		"email":    testData.TestUsername + "_updated@example.com",
		"nickname": "测试用户（已更新）",
	}

	resp, err := makeRequest("PUT", fmt.Sprintf("/users/%d", testData.UserID), data, getAuthHeaders())
	if err != nil {
		printResult("更新用户信息", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("更新用户信息", passed, resp, nil)
}

// 测试更新用户状态
func testUpdateUserStatus() {
	data := map[string]interface{}{
		"status": 1,
	}

	resp, err := makeRequest("PUT", fmt.Sprintf("/users/%d/status", testData.UserID), data, getAuthHeaders())
	if err != nil {
		printResult("更新用户状态", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("更新用户状态", passed, resp, nil)
}

// 测试修改密码
func testChangePassword() {
	data := map[string]interface{}{
		"old_password": testData.TestPassword,
		"new_password": "NewTest123456",
	}

	resp, err := makeRequest("PUT", "/users/change-password", data, getAuthHeaders())
	if err != nil {
		printResult("修改密码", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	if passed {
		testData.TestPassword = data["new_password"].(string)
	}
	printResult("修改密码", passed, resp, nil)
}