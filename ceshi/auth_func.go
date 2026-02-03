package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// 测试验证码
func testCaptcha() {
	resp, err := makeRequest("GET", "/captcha", nil, nil)
	if err != nil {
		printResult("获取验证码", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取验证码", passed, resp, nil)
}

// 测试用户登录
func testLogin() {
	// 动态生成测试用户名
	testData.TestUsername = generateTestUsername()
	
	data := map[string]interface{}{
		"username": testData.Username,
		"password": testData.Password,
	}

	resp, err := makeRequest("POST", "/login", data, nil)
	if err != nil {
		printResult("用户登录", false, nil, err)
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
				if token, ok := dataMap["token"].(string); ok {
					testData.Token = token
				}
				if user, ok := dataMap["user"].(map[string]interface{}); ok {
					if id, ok := user["id"].(float64); ok {
						testData.UserID = uint(id)
					}
				}
			}
		}
		passed = true
	}
	printResult("用户登录", passed, resp, nil)
}

// 测试退出登录
func testLogout() {
	resp, err := makeRequest("POST", "/logout", nil, getAuthHeaders())
	if err != nil {
		printResult("退出登录", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("退出登录", passed, resp, nil)
}

// 生成测试用户名
func generateTestUsername() string {
	return fmt.Sprintf("testuser_%d", time.Now().Unix())
}