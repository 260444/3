package operationtool

// ExecuteCommandRequest 批量执行命令请求
type ExecuteCommandRequest struct {
	HostIDs  []uint `json:"host_ids" binding:"required"`
	Commands string `json:"commands" binding:"required"`
}

// ExecuteCommandResponse 批量执行命令响应
type ExecuteCommandResponse struct {
	Total   int                 `json:"total"`
	Success int                 `json:"success"`
	Failed  int                 `json:"failed"`
	Results []HostCommandResult `json:"results"`
}

// HostCommandResult 单台主机执行结果
type HostCommandResult struct {
	HostID      uint   `json:"host_id"`
	Hostname    string `json:"hostname"`
	IPAddress   string `json:"ip_address"`
	Success     bool   `json:"success"`
	Output      string `json:"output,omitempty"`
	Error       string `json:"error,omitempty"`
	ExecuteTime int64  `json:"execute_time"` // 毫秒
}
