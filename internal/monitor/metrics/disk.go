package metrics

// DiskInfo disk info.
type DiskInfo struct {
	PartitionStats []PartitionStat `json:"partition_stats"`
}

// DiskStat disk stat.
type DiskStat struct {
	PartitionWithUsageAndIOStat []PartitionWithUsageAndIOStat `json:"partition_with_usage_stats"`
	// DiskIOCountersStats     []DiskIOCountersStat     `json:"disk_io_counters_stats"`
}

// DiskIOCountersStat io counters stat.
type DiskIOCountersStat struct {
	ReadCount        uint64 `json:"read_count"`
	MergedReadCount  uint64 `json:"merged_read_count"`
	WriteCount       uint64 `json:"write_count"`
	MergedWriteCount uint64 `json:"merged_write_count"`
	ReadBytes        uint64 `json:"read_bytes"`
	WriteBytes       uint64 `json:"write_bytes"`
	ReadTime         uint64 `json:"read_time"`
	WriteTime        uint64 `json:"write_time"`
	IopsInProgress   uint64 `json:"iops_in_progress"`
	IoTime           uint64 `json:"io_time"`
	WeightedIO       uint64 `json:"weighted_io"`
	Name             string `json:"name"`
	SerialNumber     string `json:"serial_number"`
	Label            string `json:"label"`
}

// PartitionWithUsageAndIOStat partition stat with usage and io stat.
type PartitionWithUsageAndIOStat struct {
	// partition stat
	Device     string `json:"device"`
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Opts       string `json:"opts"`

	// usage stat
	Path              string  `json:"path"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`

	// io counters stat
	ReadCount        uint64 `json:"read_count"`
	MergedReadCount  uint64 `json:"merged_read_count"`
	WriteCount       uint64 `json:"write_count"`
	MergedWriteCount uint64 `json:"merged_write_count"`
	ReadBytes        uint64 `json:"read_bytes"`
	WriteBytes       uint64 `json:"write_bytes"`
	ReadTime         uint64 `json:"read_time"`
	WriteTime        uint64 `json:"write_time"`
	IopsInProgress   uint64 `json:"iops_in_progress"`
	IoTime           uint64 `json:"io_time"`
	WeightedIO       uint64 `json:"weighted_io"`
	Name             string `json:"name"`
	SerialNumber     string `json:"serial_number"`
	Label            string `json:"label"`
}

// PartitionStat partition stat.
type PartitionStat struct {
	Device     string `json:"device"`
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Opts       string `json:"opts"`
}

// UsageStat usage stat.
type UsageStat struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}
