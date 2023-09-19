package metrics

// AlertStat alert stat.
type AlertStat struct {
	AverageCpuPercent  float64 `json:"average_cpu_percent"`
	AverageMemPercent  float64 `json:"average_mem_percent"`  // For virtual memory used percent
	AverageDiskPercent float64 `json:"average_disk_percent"` // For path "/"
	ConnectNetCount    int     `json:"connect_net_count"`

	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}
