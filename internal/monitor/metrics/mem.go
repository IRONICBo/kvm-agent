package metrics

// MemInfo mem info.
type MemInfo struct {
	SwapDevices []SwapDevice `json:"swap_devices"`
}

// SwapDevice swap device.
type SwapDevice struct {
	Name      string `json:"name"`
	UsedBytes uint64 `json:"used_bytes"`
	FreeBytes uint64 `json:"free_bytes"`
}

// MemStat mem stat.
type MemStat struct {
	SwapMemoryStat      *SwapMemoryStat      `json:"swap_memory_stat"`
	VirtualMemoryStat   *VirtualMemoryStat   `json:"virtual_memory_stat"`
	VirtualMemoryExStat *VirtualMemoryExStat `json:"virtual_memory_ex_stat"`
}

// SwapMemoryStat swap memory stat.
type SwapMemoryStat struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
	Sin         uint64  `json:"sin"`
	Sout        uint64  `json:"sout"`
	PgIn        uint64  `json:"pgin"`
	PgOut       uint64  `json:"pgout"`
	PgFault     uint64  `json:"pgfault"`

	// Linux specific numbers
	// https://www.kernel.org/doc/Documentation/cgroup-v2.txt
	PgMajFault uint64 `json:"pgmajfault"`
}

// VirtualMemoryExStat virtual memory ex stat.
type VirtualMemoryExStat struct {
	ActiveFile   uint64 `json:"activefile"`
	InactiveFile uint64 `json:"inactivefile"`
	ActiveAnon   uint64 `json:"activeanon"`
	InactiveAnon uint64 `json:"inactiveanon"`
	Unevictable  uint64 `json:"unevictable"`
}

// VirtualMemoryStat virtual memory stat.
type VirtualMemoryStat struct {
	// Total amount of RAM on this system
	Total uint64 `json:"total"`

	// RAM available for programs to allocate
	//
	// This value is computed from the kernel specific values.
	Available uint64 `json:"available"`

	// RAM used by programs
	//
	// This value is computed from the kernel specific values.
	Used uint64 `json:"used"`

	// Percentage of RAM used by programs
	//
	// This value is computed from the kernel specific values.
	UsedPercent float64 `json:"used_percent"`

	// This is the kernel's notion of free memory; RAM chips whose bits nobody
	// cares about the value of right now. For a human consumable number,
	// Available is what you really want.
	Free uint64 `json:"free"`

	// OS X / BSD specific numbers:
	// http://www.macyourself.com/2010/02/17/what-is-free-wired-active-and-inactive-system-memory-ram/
	Active   uint64 `json:"active"`
	Inactive uint64 `json:"inactive"`
	Wired    uint64 `json:"wired"`

	// FreeBSD specific numbers:
	// https://reviews.freebsd.org/D8467
	Laundry uint64 `json:"laundry"`

	// Linux specific numbers
	// https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html
	// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
	// https://www.kernel.org/doc/Documentation/vm/overcommit-accounting
	Buffers        uint64 `json:"buffers"`
	Cached         uint64 `json:"cached"`
	Writeback      uint64 `json:"writeback"`
	Dirty          uint64 `json:"dirty"`
	WritebackTmp   uint64 `json:"writebacktmp"`
	Shared         uint64 `json:"shared"`
	Slab           uint64 `json:"slab"`
	SReclaimable   uint64 `json:"sreclaimable"`
	SUnreclaim     uint64 `json:"sunreclaim"`
	PageTables     uint64 `json:"pagetables"`
	SwapCached     uint64 `json:"swapcached"`
	CommitLimit    uint64 `json:"commitlimit"`
	CommittedAS    uint64 `json:"committedas"`
	HighTotal      uint64 `json:"hightotal"`
	HighFree       uint64 `json:"highfree"`
	LowTotal       uint64 `json:"lowtotal"`
	LowFree        uint64 `json:"lowfree"`
	SwapTotal      uint64 `json:"swaptotal"`
	SwapFree       uint64 `json:"swapfree"`
	Mapped         uint64 `json:"mapped"`
	VMallocTotal   uint64 `json:"vmalloctotal"`
	VMallocUsed    uint64 `json:"vmallocused"`
	VMallocChunk   uint64 `json:"vmallocchunk"`
	HugePagesTotal uint64 `json:"hugepagestotal"`
	HugePagesFree  uint64 `json:"hugepagesfree"`
	HugePageSize   uint64 `json:"hugepagesize"`
}
