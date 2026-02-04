package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 测试获取操作日志
func testGetOperationLogs() {
	fmt.Println("\n--- 测试获取操作日志 ---")

	resp, body, err := sendAuthenticatedGet(BaseURL + "/api/v1/operation-logs?page=1&page_size=10")
	if err != nil {
		addTestResult("获取操作日志", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			addTestResult("获取操作日志", false, fmt.Sprintf("解析响应失败: %v", err))
			return
		}
		addTestResult("获取操作日志", true, fmt.Sprintf("获取成功，数据: %s", prettyPrintJSON(result["data"])))
	} else {
		addTestResult("获取操作日志", false, fmt.Sprintf("获取失败，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}

// 测试删除操作日志
func testDeleteOperationLog() {
	fmt.Println("\n--- 测试删除操作日志 ---")

	resp, body, err := sendAuthenticatedDelete(BaseURL + "/api/v1/operation-logs/999") // 使用不存在的ID
	if err != nil {
		addTestResult("删除操作日志", false, fmt.Sprintf("请求失败: %v", err))
		return
	}
	defer resp.Body.Close()

	// 期望返回404或其他错误状态码
	if resp.StatusCode != http.StatusOK {
		addTestResult("删除操作日志", true, fmt.Sprintf("删除成功或正确处理了不存在的日志，状态码: %d", resp.StatusCode))
	} else {
		addTestResult("删除操作日志", false, fmt.Sprintf("意外的成功响应，状态码: %d, 响应: %s", resp.StatusCode, string(body)))
	}
}
