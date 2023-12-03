package snmp

import (
	"fmt"
	"kvm-agent/internal/log"
	"kvm-agent/internal/monitor/metrics"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

func GetSNMPStat() *metrics.CpuStat {
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

func GetAllStat(uuid string, timeout int) metrics.MetricStat {
	metric := metrics.MetricStat{
		UUID:      uuid,
		Timestamp: time.Now().Unix(),
	}

	var wg sync.WaitGroup

	wg.Add(5)
	go func() {
		metric.CpuStat = GetSNMPStat()
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
