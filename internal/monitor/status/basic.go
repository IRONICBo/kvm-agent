package status

import (
	"fmt"
	"kvm-agent/internal/log"
	"kvm-agent/internal/monitor/metrics"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func GetCpuStat() *metrics.CpuStat {
	// Get raw data
	cpuTimeStats, err := cpu.Times(true)
	if err != nil {
		log.Errorf("GetCpuStat", "cpu.Times(true) error: %s", err.Error())
		cpuTimeStats = []cpu.TimesStat{}
	}

	cpuPercentStats, err := cpu.Percent(0, true)
	if err != nil {
		log.Errorf("GetCpuStat", "cpu.Percent(0, true) error: %s", err.Error())
		cpuPercentStats = []float64{}
	}

	cpuLoad, err := load.Avg()
	if err != nil {
		log.Errorf("GetCpuStat", "load.Avg() error: %s", err.Error())
		cpuLoad = &load.AvgStat{}
	}

	// Fill data and allocate
	metric := &metrics.CpuStat{}
	metric.CPUTimeStats = make([]metrics.CPUTimeStat, 0)
	metric.CPUPercents = make([]float64, 0)
	metric.CPULoad = &metrics.AvgStat{}

	for _, cpuTimeStat := range cpuTimeStats {
		tempCPUTimeStat := metrics.CPUTimeStat{
			CPU:       cpuTimeStat.CPU,
			User:      cpuTimeStat.User,
			System:    cpuTimeStat.System,
			Idle:      cpuTimeStat.Idle,
			Nice:      cpuTimeStat.Nice,
			Iowait:    cpuTimeStat.Iowait,
			Irq:       cpuTimeStat.Irq,
			Softirq:   cpuTimeStat.Softirq,
			Steal:     cpuTimeStat.Steal,
			Guest:     cpuTimeStat.Guest,
			GuestNice: cpuTimeStat.GuestNice,
		}
		metric.CPUTimeStats = append(metric.CPUTimeStats, tempCPUTimeStat)
	}
	for _, cpuPercentStat := range cpuPercentStats {
		metric.CPUPercents = append(metric.CPUPercents, cpuPercentStat)
	}
	metric.CPULoad = &metrics.AvgStat{
		Load1:  cpuLoad.Load1,
		Load5:  cpuLoad.Load5,
		Load15: cpuLoad.Load15,
	}

	return metric
}

func GetMemStat() *metrics.MemStat {
	// get raw data
	swapMem, err := mem.SwapMemory()
	if err != nil {
		log.Errorf("GetMemStat", "mem.SwapMemory() error: %s", err.Error())
		swapMem = &mem.SwapMemoryStat{}
	}

	virtualMem, err := mem.VirtualMemory()
	if err != nil {
		log.Errorf("GetMemStat", "mem.VirtualMemory() error: %s", err.Error())
		virtualMem = &mem.VirtualMemoryStat{}
	}

	virtualMemEx, err := mem.VirtualMemoryEx()
	if err != nil {
		log.Errorf("GetMemStat", "mem.VirtualMemoryEx() error: %s", err.Error())
		virtualMemEx = &mem.VirtualMemoryExStat{}
	}

	// fill data and allocate
	metric := &metrics.MemStat{}
	metric.SwapMemoryStat = &metrics.SwapMemoryStat{
		Total:       swapMem.Total,
		Used:        swapMem.Used,
		Free:        swapMem.Free,
		UsedPercent: swapMem.UsedPercent,
		Sin:         swapMem.Sin,
		Sout:        swapMem.Sout,
		PgIn:        swapMem.PgIn,
		PgOut:       swapMem.PgOut,
		PgFault:     swapMem.PgFault,
		PgMajFault:  swapMem.PgMajFault,
	}
	metric.VirtualMemoryStat = &metrics.VirtualMemoryStat{
		Total:     virtualMem.Total,
		Available: virtualMem.Available,
	}
	metric.VirtualMemoryExStat = &metrics.VirtualMemoryExStat{
		ActiveFile:   virtualMemEx.ActiveFile,
		InactiveFile: virtualMemEx.InactiveFile,
		ActiveAnon:   virtualMemEx.ActiveAnon,
		InactiveAnon: virtualMemEx.InactiveAnon,
		Unevictable:  virtualMemEx.Unevictable,
	}

	return metric
}

func GetDiskStat() *metrics.DiskStat {
	// Get raw data
	partitionStats, err := disk.Partitions(true)
	if err != nil {
		log.Errorf("GetDiskStat", "disk.Partitions(true) error: %s", err.Error())
		partitionStats = []disk.PartitionStat{}
	}

	metric := &metrics.DiskStat{}
	metric.PartitionWithUsageAndIOStat = make([]metrics.PartitionWithUsageAndIOStat, 0)

	// foreach partitionStats
	for _, partitionStat := range partitionStats {
		usage, err := disk.Usage(partitionStat.Mountpoint)
		if err != nil {
			log.Errorf("GetDiskStat", "disk.Usage(%s) error: %s", partitionStat.Mountpoint, err.Error())
			usage = &disk.UsageStat{}
		}

		ioCounters, err := disk.IOCounters(partitionStat.Device)
		if err != nil {
			log.Errorf("GetDiskStat", "disk.IOCounters(%s) error: %s", partitionStat.Device, err.Error())
			ioCounters = make(map[string]disk.IOCountersStat, 0)
		}

		// Map length is 0 or 1
		// Get IOCountersStat instance
		var diskIOCountersStat metrics.DiskIOCountersStat
		for _, ioCountersStat := range ioCounters {
			diskIOCountersStat = metrics.DiskIOCountersStat{
				ReadCount:        ioCountersStat.ReadCount,
				MergedReadCount:  ioCountersStat.MergedReadCount,
				WriteCount:       ioCountersStat.WriteCount,
				MergedWriteCount: ioCountersStat.MergedWriteCount,
				ReadBytes:        ioCountersStat.ReadBytes,
				WriteBytes:       ioCountersStat.WriteBytes,
				ReadTime:         ioCountersStat.ReadTime,
				WriteTime:        ioCountersStat.WriteTime,
				IopsInProgress:   ioCountersStat.IopsInProgress,
				IoTime:           ioCountersStat.IoTime,
				WeightedIO:       ioCountersStat.WeightedIO,
				Name:             ioCountersStat.Name,
				SerialNumber:     ioCountersStat.SerialNumber,
				Label:            ioCountersStat.Label,
			}
		}

		partitionWithUsageStat := metrics.PartitionWithUsageAndIOStat{
			Device:     partitionStat.Device,
			Mountpoint: partitionStat.Mountpoint,
			Fstype:     partitionStat.Fstype,
			Opts:       partitionStat.Opts,

			Path:              usage.Path,
			Total:             usage.Total,
			Free:              usage.Free,
			Used:              usage.Used,
			UsedPercent:       usage.UsedPercent,
			InodesTotal:       usage.InodesTotal,
			InodesUsed:        usage.InodesUsed,
			InodesFree:        usage.InodesFree,
			InodesUsedPercent: usage.InodesUsedPercent,

			ReadCount:        diskIOCountersStat.ReadCount,
			MergedReadCount:  diskIOCountersStat.MergedReadCount,
			WriteCount:       diskIOCountersStat.WriteCount,
			MergedWriteCount: diskIOCountersStat.MergedWriteCount,
			ReadBytes:        diskIOCountersStat.ReadBytes,
			WriteBytes:       diskIOCountersStat.WriteBytes,
			ReadTime:         diskIOCountersStat.ReadTime,
			WriteTime:        diskIOCountersStat.WriteTime,
			IopsInProgress:   diskIOCountersStat.IopsInProgress,
			IoTime:           diskIOCountersStat.IoTime,
			WeightedIO:       diskIOCountersStat.WeightedIO,
			Name:             diskIOCountersStat.Name,
			SerialNumber:     diskIOCountersStat.SerialNumber,
			Label:            diskIOCountersStat.Label,
		}
		metric.PartitionWithUsageAndIOStat = append(metric.PartitionWithUsageAndIOStat, partitionWithUsageStat)
	}

	return metric
}

func GetNetStat() *metrics.NetStat {
	// Get raw data
	connections, err := net.Connections("all")
	if err != nil {
		log.Errorf("GetNetStat", "net.Connections(all) error: %s", err.Error())
		connections = []net.ConnectionStat{}
	}
	cks, err := net.ConntrackStats(false)
	if err != nil {
		log.Errorf("GetNetStat", "net.ConntrackStats(false) error: %s", err.Error())
		cks = []net.ConntrackStat{}
	}
	fc, err := net.FilterCounters()
	if err != nil {
		log.Errorf("GetNetStat", "net.FilterCounters() error: %s", err.Error())
		fc = []net.FilterStat{}
	}
	netioc, err := net.IOCounters(true)
	if err != nil {
		log.Errorf("GetNetStat", "net.IOCounters(true) error: %s", err.Error())
		netioc = []net.IOCountersStat{}
	}
	pc, err := net.ProtoCounters(nil)
	if err != nil {
		log.Errorf("GetNetStat", "net.ProtoCounters(nil) error: %s", err.Error())
		pc = []net.ProtoCountersStat{}
	}

	// Fill data and allocate
	metric := &metrics.NetStat{}
	metric.ConnectionStats = make([]metrics.ConnectionStat, 0)
	metric.ConntrackStats = make([]metrics.ConntrackStat, 0)
	metric.FilterStats = make([]metrics.FilterStat, 0)
	metric.NetIOCountersStats = make(map[string]interface{}, 0)
	metric.ProtoCountersStats = make(map[string]interface{}, 0)

	for _, connection := range connections {
		tempConnectionStat := metrics.ConnectionStat{
			Fd:     connection.Fd,
			Family: connection.Family,
			Type:   connection.Type,
			Laddr: metrics.Addr{
				IP:   connection.Laddr.IP,
				Port: connection.Laddr.Port,
			},
			Raddr: metrics.Addr{
				IP:   connection.Raddr.IP,
				Port: connection.Raddr.Port,
			},
			Status: connection.Status,
			Uids:   connection.Uids,
			Pid:    connection.Pid,
		}
		metric.ConnectionStats = append(metric.ConnectionStats, tempConnectionStat)
	}

	for _, ck := range cks {
		tempConntrackStat := metrics.ConntrackStat{
			Entries:       ck.Entries,
			Searched:      ck.Searched,
			Found:         ck.Found,
			New:           ck.New,
			Invalid:       ck.Invalid,
			Ignore:        ck.Ignore,
			Delete:        ck.Delete,
			DeleteList:    ck.DeleteList,
			Insert:        ck.Insert,
			InsertFailed:  ck.InsertFailed,
			Drop:          ck.Drop,
			EarlyDrop:     ck.EarlyDrop,
			IcmpError:     ck.IcmpError,
			ExpectNew:     ck.ExpectNew,
			ExpectCreate:  ck.ExpectCreate,
			ExpectDelete:  ck.ExpectDelete,
			SearchRestart: ck.SearchRestart,
		}
		metric.ConntrackStats = append(metric.ConntrackStats, tempConntrackStat)
	}

	for _, f := range fc {
		tempFilterStat := metrics.FilterStat{
			ConnTrackCount: f.ConnTrackCount,
			ConnTrackMax:   f.ConnTrackMax,
		}
		metric.FilterStats = append(metric.FilterStats, tempFilterStat)
	}

	for _, n := range netioc {
		tempNetIOCountersStat := metrics.NetIOCountersStat{
			BytesSent:   n.BytesSent,
			BytesRecv:   n.BytesRecv,
			PacketsSent: n.PacketsSent,
			PacketsRecv: n.PacketsRecv,
			Errin:       n.Errin,
			Errout:      n.Errout,
			Dropin:      n.Dropin,
			Dropout:     n.Dropout,
		}
		metric.NetIOCountersStats[n.Name] = tempNetIOCountersStat // key is interface name
	}

	for _, p := range pc {
		tempProtoCountersStat := metrics.ProtoCountersStat{
			Protocol: p.Protocol,
			Stats:    p.Stats,
		}
		metric.ProtoCountersStats[tempProtoCountersStat.Protocol] = tempProtoCountersStat.Stats // key is protocol name
	}

	return metric
}

func GetProcessStat() *metrics.ProcessStat {
	// Get raw data
	processes, err := process.Processes()
	if err != nil {
		log.Errorf("GetProcessStat", "metrics.GetProcesses() error: %s", err.Error())
		processes = []*process.Process{}
	}

	// Fill data and allocate
	metric := &metrics.ProcessStat{}
	metric.Processes = make([]metrics.Process, 0)

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			log.Errorf("GetProcessStat", "p.Name() error: %s", err.Error())
			name = ""
		}
		status, err := p.Status()
		if err != nil {
			log.Errorf("GetProcessStat", "p.Status() error: %s", err.Error())
			status = ""
		}
		parent, err := p.Ppid()
		if err != nil {
			log.Errorf("GetProcessStat", "p.Ppid() error: %s", err.Error())
			parent = 0
		}
		ncs, err := p.NumCtxSwitches()
		if err != nil {
			log.Errorf("GetProcessStat", "p.NumCtxSwitches() error: %s", err.Error())
			ncs = &process.NumCtxSwitchesStat{}
		}
		uids, err := p.Uids()
		if err != nil {
			log.Errorf("GetProcessStat", "p.Uids() error: %s", err.Error())
			uids = []int32{}
		}
		groups, err := p.Groups()
		if err != nil {
			log.Errorf("GetProcessStat", "p.Groups() error: %s", err.Error())
			groups = []int32{}
		}
		gids, err := p.Gids()
		if err != nil {
			log.Errorf("GetProcessStat", "p.Gids() error: %s", err.Error())
			gids = []int32{}
		}
		memInfo, err := p.MemoryInfo()
		if err != nil {
			log.Errorf("GetProcessStat", "p.MemoryInfo() error: %s", err.Error())
			memInfo = &process.MemoryInfoStat{}
		}
		tgid, err := p.Tgid()
		if err != nil {
			log.Errorf("GetProcessStat", "p.Tgid() error: %s", err.Error())
			tgid = 0
		}

		tempProcessStat := metrics.Process{
			Pid:    p.Pid,
			Name:   name,
			Status: status,
			Parent: parent,
			NumCtxSwitches: &metrics.NumCtxSwitchesStat{
				Voluntary:   ncs.Voluntary,
				Involuntary: ncs.Involuntary,
			},
			Uids:   uids,
			Gids:   gids,
			Groups: groups,
			MemInfo: &metrics.MemoryInfoStat{
				RSS:    memInfo.RSS,
				VMS:    memInfo.VMS,
				HWM:    memInfo.HWM,
				Data:   memInfo.Data,
				Stack:  memInfo.Stack,
				Locked: memInfo.Locked,
				Swap:   memInfo.Swap,
			},
			SigInfo:     &metrics.SignalInfoStat{},
			LastCPUTime: time.Now(),
			Tgid:        tgid,
		}
		metric.Processes = append(metric.Processes, tempProcessStat)
	}

	return metric
}

func GetAllStat(uuid string, timeout int) metrics.MetricStat {
	metric := metrics.MetricStat{
		UUID:      uuid,
		Timestamp: time.Now().Unix(),
	}

	var wg sync.WaitGroup

	wg.Add(5)
	go func() {
		metric.CpuStat = GetCpuStat()
		wg.Done()
	}()

	go func() {
		metric.MemStat = GetMemStat()
		wg.Done()
	}()

	go func() {
		metric.DiskStat = GetDiskStat()
		wg.Done()
	}()

	go func() {
		metric.NetStat = GetNetStat()
		wg.Done()
	}()

	go func() {
		metric.ProcessStat = GetProcessStat()
		wg.Done()
	}()

	timeoutEvent := time.Duration(timeout) * time.Second
	select {
	case <-time.After(timeoutEvent):
		log.Errorf("GetAllStat", "timeout :%s", timeoutEvent.String())

		// Calc alert data
		metric.AlertStat = calcAlertData(&metric)
		fmt.Printf("data: %#v", metric)

		return metric
	}
}

// CalcAlertData calculate alert data.
// This function is used to calculate the average value of the data.
func calcAlertData(metric *metrics.MetricStat) *metrics.AlertStat {
	alertStat := &metrics.AlertStat{}

	// CPU
	cpuPercentSum := 0.0
	for _, cpuPercent := range metric.CpuStat.CPUPercents {
		cpuPercentSum += cpuPercent
	}
	alertStat.AverageCpuPercent = cpuPercentSum / float64(len(metric.CpuStat.CPUPercents))

	// MEM
	alertStat.AverageMemPercent = metric.MemStat.VirtualMemoryStat.UsedPercent

	// DISK
	// diskPercentSum := 0.0
	// for _, partitionWithUsageAndIOStat := range metric.DiskStat.PartitionWithUsageAndIOStat {
	// 	diskPercentSum += partitionWithUsageAndIOStat.UsedPercent
	// }
	// alertStat.AverageDiskPercent = diskPercentSum / float64(len(metric.DiskStat.PartitionWithUsageAndIOStat))

	// Now only use the "/" path
	for _, partitionWithUsageAndIOStat := range metric.DiskStat.PartitionWithUsageAndIOStat {
		if partitionWithUsageAndIOStat.Path == "/" {
			alertStat.AverageDiskPercent = partitionWithUsageAndIOStat.UsedPercent
		}
	}

	// NET
	alertStat.ConnectNetCount = len(metric.NetStat.ConnectionStats)

	// LOAD
	alertStat.Load1 = metric.CpuStat.CPULoad.Load1
	alertStat.Load5 = metric.CpuStat.CPULoad.Load5
	alertStat.Load15 = metric.CpuStat.CPULoad.Load15

	return alertStat
}
