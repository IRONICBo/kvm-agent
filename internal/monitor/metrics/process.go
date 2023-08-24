package metrics

import (
	"time"
)

// SimpleProcessStat simple process stat.
type SimpleProcessStat struct {
	SimpleProcessItem []SimpleProcessItem `json:"simple_process_item"`
}

// SimpleProcessItem simple process item.
type SimpleProcessItem struct {
	PID     string `json:"pid"`
	USER    string `json:"user"`
	PR      string `json:"pr"`
	NI      string `json:"ni"`
	VIRT    string `json:"virt"`
	RES     string `json:"res"`
	SHR     string `json:"shr"`
	S       string `json:"s"`
	CPU     string `json:"cpu"`
	MEM     string `json:"mem"`
	TIME    string `json:"time"`
	COMMAND string `json:"command"`
}

type ProcessStat struct {
	Processes []Process `json:"processes"`
}

type Process struct {
	Pid            int32               `json:"pid"`
	Name           string              `json:"name"`
	Status         string              `json:"status"`
	Parent         int32               `json:"parent"`
	NumCtxSwitches *NumCtxSwitchesStat `json:"num_ctx_switches"`
	Uids           []int32             `json:"uids"`
	Gids           []int32             `json:"gids"`
	Groups         []int32             `json:"groups"`
	NumThreads     int32               `json:"num_threads"`
	MemInfo        *MemoryInfoStat     `json:"mem_info"`
	SigInfo        *SignalInfoStat     `json:"sig_info"`
	CreateTime     int64               `json:"create_time"`

	LastCPUTimes *CPUTimeStat `json:"last_cpu_times"`
	LastCPUTime  time.Time    `json:"last_cpu_time"`

	Tgid int32 `json:"tgid"`
}

// OpenFilesStat open files stat.
type OpenFilesStat struct {
	Path string `json:"path"`
	Fd   uint64 `json:"fd"`
}

// MemoryInfoStat memory info stat.
type MemoryInfoStat struct {
	RSS    uint64 `json:"rss"`    // bytes
	VMS    uint64 `json:"vms"`    // bytes
	HWM    uint64 `json:"hwm"`    // bytes
	Data   uint64 `json:"data"`   // bytes
	Stack  uint64 `json:"stack"`  // bytes
	Locked uint64 `json:"locked"` // bytes
	Swap   uint64 `json:"swap"`   // bytes
}

// SignalInfoStat signal info stat.
type SignalInfoStat struct {
	PendingProcess uint64 `json:"pending_process"`
	PendingThread  uint64 `json:"pending_thread"`
	Blocked        uint64 `json:"blocked"`
	Ignored        uint64 `json:"ignored"`
	Caught         uint64 `json:"caught"`
}

// RlimitStat rlimit stat.
type RlimitStat struct {
	Resource int32  `json:"resource"`
	Soft     int32  `json:"soft"` //TODO too small. needs to be uint64
	Hard     int32  `json:"hard"` //TODO too small. needs to be uint64
	Used     uint64 `json:"used"`
}

// IOCountersStat io counters stat.
type IOCountersStat struct {
	ReadCount  uint64 `json:"read_count"`
	WriteCount uint64 `json:"write_count"`
	ReadBytes  uint64 `json:"read_bytes"`
	WriteBytes uint64 `json:"write_bytes"`
}

// NumCtxSwitchesStat num ctx switches stat.
type NumCtxSwitchesStat struct {
	Voluntary   int64 `json:"voluntary"`
	Involuntary int64 `json:"involuntary"`
}

// PageFaultsStat page faults stat.
type PageFaultsStat struct {
	MinorFaults      uint64 `json:"minor_faults"`
	MajorFaults      uint64 `json:"major_faults"`
	ChildMinorFaults uint64 `json:"child_minor_faults"`
	ChildMajorFaults uint64 `json:"child_major_faults"`
}
