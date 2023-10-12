package metrics

// HostInfo contains common host information.
type HostInfo struct {
	HostInfoStat *HostInfoStat `json:"host_info_stat"`
}

// HostInfoStat contains common host information.
type HostInfoStat struct {
	Hostname             string `json:"hostname"`
	Uptime               uint64 `json:"uptime"`
	BootTime             uint64 `json:"boot_time"`
	Procs                uint64 `json:"procs"`            // number of processes
	OS                   string `json:"os"`               // ex: freebsd, linux
	Platform             string `json:"platform"`         // ex: ubuntu, linuxmint
	PlatformFamily       string `json:"platform_family"`  // ex: debian, rhel
	PlatformVersion      string `json:"platform_version"` // version of the complete OS
	KernelVersion        string `json:"kernel_version"`   // version of the OS kernel (if available)
	KernelArch           string `json:"kernel_arch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
	VirtualizationSystem string `json:"virtualization_system"`
	VirtualizationRole   string `json:"virtualization_role"` // guest or host
	HostID               string `json:"hostid"`              // ex: uuid

	SystemdInfos []SystemdInfo `json:"systemd_infos"`
}

// SystemdInfo contains systemd information.
type SystemdInfo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// TODO: ignore
// TemperatureStat contains temperature information.
type TemperatureStat struct {
	SensorKey   string  `json:"sensorKey"`
	Temperature float64 `json:"sensorTemperature"`
}
