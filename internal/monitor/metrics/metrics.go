package metrics

type MetricStat struct {
	UUID      string `json:"uuid"`
	Timestamp int64  `json:"timestamp"`

	CpuStat     *CpuStat     `json:"cpu_stat"`
	MemStat     *MemStat     `json:"mem_stat"`
	DiskStat    *DiskStat    `json:"disk_stat"`
	NetStat     *NetStat     `json:"net_stat"`
	ProcessStat *ProcessStat `json:"process_stat"`
}
