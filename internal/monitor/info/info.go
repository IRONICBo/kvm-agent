package info

import (
	"encoding/json"
	"kvm-agent/internal/monitor/metrics"
	"kvm-agent/internal/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// GetCpuInfo returns cpu info
func GetCpuInfo() (metrics.CpuInfo, error) {
	info, err := cpu.Info()
	if err != nil {
		return metrics.CpuInfo{}, err
	}

	cpuBasicInfos := make([]metrics.CpuBasicInfo, 0)
	for _, v := range info {
		cpuBasicInfos = append(cpuBasicInfos, metrics.CpuBasicInfo{
			CPU:        v.CPU,
			VendorID:   v.VendorID,
			Family:     v.Family,
			Model:      v.Model,
			Stepping:   v.Stepping,
			PhysicalID: v.PhysicalID,
			CoreID:     v.CoreID,
			Cores:      v.Cores,
			ModelName:  v.ModelName,
			Mhz:        v.Mhz,
			CacheSize:  v.CacheSize,
			Flags:      v.Flags,
			Microcode:  v.Microcode,
		})
	}

	return metrics.CpuInfo{
		CpuBasicInfos: cpuBasicInfos,
	}, nil
}

// GetCpuInfoJson returns cpu info in json format
func GetCpuInfoJsonCompressed() string {
	cpuInfo, _ := GetCpuInfo()
	marshal, err := json.Marshal(cpuInfo)
	if err != nil {
		return ""
	}

	marshal, err = utils.CompressText(string(marshal))
	if err != nil {
		return ""
	}

	return utils.Base64Encode(marshal)
}

// GetMemInfo returns mem info
func GetMemInfo() (metrics.MemInfo, error) {
	devices, err := mem.SwapDevices()
	if err != nil {
		return metrics.MemInfo{}, err
	}

	swapDevices := make([]metrics.SwapDevice, 0)
	for _, v := range devices {
		swapDevices = append(swapDevices, metrics.SwapDevice{
			Name:      v.Name,
			UsedBytes: v.UsedBytes,
			FreeBytes: v.FreeBytes,
		})
	}

	return metrics.MemInfo{
		SwapDevices: swapDevices,
	}, nil
}

// GetMemInfoJson returns mem info in json format
func GetMemInfoJsonCompressed() string {
	memInfo, _ := GetMemInfo()
	marshal, err := json.Marshal(memInfo)
	if err != nil {
		return ""
	}

	marshal, err = utils.CompressText(string(marshal))
	if err != nil {
		return ""
	}

	return utils.Base64Encode(marshal)
}

// GetDiskInfo returns disk info
func GetDiskInfo() (metrics.DiskInfo, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return metrics.DiskInfo{}, err
	}

	diskPartitions := make([]metrics.PartitionStat, 0)
	for _, v := range partitions {
		diskPartitions = append(diskPartitions, metrics.PartitionStat{
			Device:     v.Device,
			Mountpoint: v.Mountpoint,
			Fstype:     v.Fstype,
			Opts:       v.Opts,
		})
	}

	return metrics.DiskInfo{
		PartitionStats: diskPartitions,
	}, nil
}

// GetDiskInfoJson returns disk info in json format
func GetDiskInfoJsonCompressed() string {
	diskInfo, _ := GetDiskInfo()
	marshal, err := json.Marshal(diskInfo)
	if err != nil {
		return ""
	}

	marshal, err = utils.CompressText(string(marshal))
	if err != nil {
		return ""
	}

	return utils.Base64Encode(marshal)
}

// GetNetInfo returns net info
func GetNetInfo() (metrics.NetInfo, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return metrics.NetInfo{}, err
	}

	interfaceInfos := make([]metrics.InterfaceInfo, 0)
	for _, v := range interfaces {
		interfaceAddrs := make([]metrics.InterfaceAddr, 0)
		for _, v2 := range v.Addrs {
			interfaceAddrs = append(interfaceAddrs, metrics.InterfaceAddr{
				Addr: v2.Addr,
			})
		}

		interfaceInfos = append(interfaceInfos, metrics.InterfaceInfo{
			MTU:          v.MTU,
			Name:         v.Name,
			HardwareAddr: v.HardwareAddr,
			Flags:        v.Flags,
			Addrs:        interfaceAddrs,
		})
	}

	return metrics.NetInfo{
		InterfaceInfos: interfaceInfos,
	}, nil
}

// GetNetInfoJson returns net info in json format
func GetNetInfoJsonCompressed() string {
	netInfo, _ := GetNetInfo()
	marshal, err := json.Marshal(netInfo)
	if err != nil {
		return ""
	}

	marshal, err = utils.CompressText(string(marshal))
	if err != nil {
		return ""
	}

	return utils.Base64Encode(marshal)
}

// GetHostInfo returns host info
func GetHostInfo() (metrics.HostInfo, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return metrics.HostInfo{}, err
	}

	return metrics.HostInfo{
		HostInfoStat: &metrics.HostInfoStat{
			Hostname:             hostInfo.Hostname,
			Uptime:               hostInfo.Uptime,
			BootTime:             hostInfo.BootTime,
			Procs:                hostInfo.Procs,
			OS:                   hostInfo.OS,
			Platform:             hostInfo.Platform,
			PlatformFamily:       hostInfo.PlatformFamily,
			PlatformVersion:      hostInfo.PlatformVersion,
			KernelVersion:        hostInfo.KernelVersion,
			KernelArch:           hostInfo.KernelArch,
			VirtualizationSystem: hostInfo.VirtualizationSystem,
			VirtualizationRole:   hostInfo.VirtualizationRole,
			HostID:               hostInfo.HostID,
		},
	}, nil
}

// GetHostInfoJson returns host info in json format and base64 encoded
func GetHostInfoJsonCompressed() string {
	hostInfo, _ := GetHostInfo()
	marshal, err := json.Marshal(hostInfo)
	if err != nil {
		return ""
	}

	marshal, err = utils.CompressText(string(marshal))
	if err != nil {
		return ""
	}

	return utils.Base64Encode(marshal)
}
