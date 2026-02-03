package main

// 测试获取操作日志
func testGetOperationLogs() {
	resp, err := makeRequest("GET", "/operation-logs?page=1&page_size=10", nil, getAuthHeaders())
	if err != nil {
		printResult("获取操作日志", false, nil, err)
		return
	}
	defer resp.Body.Close()

	passed := resp.StatusCode == 200
	printResult("获取操作日志", passed, resp, nil)
}