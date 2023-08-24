package metrics

// CpuInfo cpu info.
type CpuInfo struct {
	CpuBasicInfos []CpuBasicInfo `json:"cpu_basic_infos"`
}

// CpuBasicInfo cpu basic info.
type CpuBasicInfo struct {
	CPU        int32    `json:"cpu"`
	VendorID   string   `json:"vendor_id"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   int32    `json:"stepping"`
	PhysicalID string   `json:"physical_id"`
	CoreID     string   `json:"core_id"`
	Cores      int32    `json:"cores"`
	ModelName  string   `json:"model_name"`
	Mhz        float64  `json:"mhz"`
	CacheSize  int32    `json:"cache_size"`
	Flags      []string `json:"flags"`
	Microcode  string   `json:"microcode"`
}

// CpuStat cpu stat.
type CpuStat struct {
	CPUTimeStats []CPUTimeStat `json:"cpu_time_stats"`
	CPUPercents  []float64     `json:"cpu_percents"`
	CPULoad      *AvgStat      `json:"cpu_load"`
}

// CPUTimeStat time stat.
type CPUTimeStat struct {
	CPU       string  `json:"cpu"`
	User      float64 `json:"user"`
	System    float64 `json:"system"`
	Idle      float64 `json:"idle"`
	Nice      float64 `json:"nice"`
	Iowait    float64 `json:"iowait"`
	Irq       float64 `json:"irq"`
	Softirq   float64 `json:"softirq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guest_nice"`
}

// AvgStat cpu avg stat.
type AvgStat struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}
