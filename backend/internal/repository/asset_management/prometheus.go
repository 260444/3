package asset_management

import (
	assModel "backend/internal/model/asset_management"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type PrometheusClient struct {
	baseURL string
	client  *http.Client
}

func NewPrometheusClient(baseURL string) *PrometheusClient {
	return &PrometheusClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

type PromQLResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]interface{} `json:"metric"`
			Value  []interface{}          `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// 查询 CPU 使用率
func (pc *PrometheusClient) GetCPUUsage(ctx context.Context) (float64, error) {
	query := `100 - (avg by(instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)`
	return pc.querySingleValue(ctx, query)
}

// 查询内存使用率
func (pc *PrometheusClient) GetMemoryUsage(ctx context.Context) (float64, error) {
	query := `(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100`
	return pc.querySingleValue(ctx, query)
}

// 查询磁盘使用率
func (pc *PrometheusClient) GetDiskUsage(ctx context.Context, mountpoint string) (float64, error) {
	query := fmt.Sprintf(`(1 - (node_filesystem_avail_bytes{mountpoint="%s"} / node_filesystem_size_bytes{mountpoint="%s"})) * 100`,
		mountpoint, mountpoint)
	return pc.querySingleValue(ctx, query)
}

// 查询主机监控指标（用于你的资产管理系统）
func (pc *PrometheusClient) GetHostMetrics(ctx context.Context, hostIP string) (*assModel.HostMetric, error) {
	// CPU 使用率
	cpuUsage, err := pc.GetCPUUsage(ctx)
	if err != nil {
		return nil, err
	}

	// 内存使用率
	memoryUsage, err := pc.GetMemoryUsage(ctx)
	if err != nil {
		return nil, err
	}

	// 磁盘使用率
	diskUsage, err := pc.GetDiskUsage(ctx, "/")
	if err != nil {
		return nil, err
	}

	return &assModel.HostMetric{
		CPUUsage:    cpuUsage,
		MemoryUsage: memoryUsage,
		DiskUsage:   diskUsage,
		RecordedAt:  time.Now(),
	}, nil
}

func (pc *PrometheusClient) querySingleValue(ctx context.Context, query string) (float64, error) {
	u, _ := url.Parse(pc.baseURL + "/api/v1/query")
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	resp, err := pc.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var promResp PromQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&promResp); err != nil {
		return 0, err
	}

	if promResp.Status != "success" || len(promResp.Data.Result) == 0 {
		return 0, fmt.Errorf("no data returned from prometheus")
	}

	// Prometheus 返回的 value 格式：[timestamp, value]
	value := promResp.Data.Result[0].Value[1]
	var result float64
	fmt.Sscanf(value.(string), "%f", &result)

	return result, nil
}
