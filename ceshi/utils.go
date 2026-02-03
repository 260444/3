package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 发送 HTTP 请求
func makeRequest(method, endpoint string, body interface{}, headers map[string]string) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	fullURL := APIBase + endpoint
	req, err := http.NewRequest(method, fullURL, reqBody)
	if err != nil {
		return nil, err
	}

	// 打印详细的请求信息
	fmt.Printf("\n--- 请求详情 ---\n")
	fmt.Printf("方法: %s\n", method)
	fmt.Printf("URL: %s\n", fullURL)
	if body != nil {
		fmt.Printf("请求体: %+v\n", body)
	}
	if len(headers) > 0 {
		fmt.Printf("请求头: %+v\n", headers)
	}
	fmt.Printf("----------------\n")

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return httpClient.Do(req)
}

// 获取认证头
func getAuthHeaders() map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + testData.Token,
	}
}

// 打印测试结果
func printResult(testName string, passed bool, resp *http.Response, err error) {
	stats.Total++
	if passed {
		stats.Passed++
		fmt.Printf("\n✓ %s - PASS\n", testName)
	} else {
		stats.Failed++
		fmt.Printf("\n✗ %s - FAIL\n", testName)
	}

	if err != nil {
		fmt.Printf("  错误: %v\n", err)
		return
	}

	if resp != nil {
		fmt.Printf("  状态码: %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, body, "", "  ")
		fmt.Printf("  响应: %s\n", prettyJSON.String())
	}
}

// 打印测试结果汇总
func printSummary() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("测试结果汇总")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("总计: %d 个测试\n", stats.Total)
	fmt.Printf("通过: %d 个\n", stats.Passed)
	fmt.Printf("失败: %d 个\n", stats.Failed)
	if stats.Total > 0 {
		fmt.Printf("成功率: %.1f%%\n", float64(stats.Passed)/float64(stats.Total)*100)
	}
	fmt.Println(strings.Repeat("=", 60))
}

// 检查服务器状态
func checkServer() bool {
	resp, err := httpClient.Get(BaseURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode < 500
}