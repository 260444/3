package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var authToken string

// 测试验证码功能
func testCaptcha() {
	fmt.Println("\n--- 测试验证码功能 ---")

	resp, body, err := sendGet(BaseURL + "/api/v1/captcha")
	if err != nil {
		addTestResult("验证码测试", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("验证码测试", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}

		if data, ok := result["data"].(map[string]interface{}); ok {
			if id, ok := data["id"].(string); ok && id != "" {
				addTestResult("验证码测试", true, "验证码生成成功")
				return
			}
		}
		addTestResult("验证码测试", false, "验证码数据格式不正确")
	} else {
		addTestResult("验证码测试", false, fmt.Sprintf("验证码生成失败，状态码: %d", resp.StatusCode))
	}
}

// 测试登录功能
func testLogin() {
	fmt.Println("\n--- 测试登录功能 ---")

	loginData := map[string]interface{}{
		"username": "admin",
		"password": "123456",
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		addTestResult("登录测试", false, fmt.Sprintf("JSON序列化失败: %v", err))
		return
	}

	resp, body, err := sendPost(BaseURL+"/api/v1/login", jsonData)
	if err != nil {
		addTestResult("登录测试", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("登录测试", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}

		if data, ok := result["data"].(map[string]interface{}); ok {
			if token, ok := data["token"].(string); ok && token != "" {
				authToken = token
				addTestResult("登录测试", true, "登录成功")
				return
			}
		}
		addTestResult("登录测试", false, "未找到有效的token")
	} else {
		addTestResult("登录测试", false, fmt.Sprintf("登录失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试登出功能
func testLogout() {
	fmt.Println("\n--- 测试登出功能 ---")

	if authToken == "" {
		addTestResult("登出测试", false, "未登录，跳过测试")
		return
	}

	req, err := http.NewRequest("POST", BaseURL+"/api/v1/logout", nil)
	if err != nil {
		addTestResult("登出测试", false, fmt.Sprintf("创建请求失败: %v", err))
		return
	}

	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		addTestResult("登出测试", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		addTestResult("登出测试", false, fmt.Sprintf("读取响应失败: %v", err))
		return
	}

	if resp.StatusCode == http.StatusOK {
		authToken = ""
		addTestResult("登出测试", true, "登出成功")
	} else {
		addTestResult("登出测试", false, fmt.Sprintf("登出失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}
