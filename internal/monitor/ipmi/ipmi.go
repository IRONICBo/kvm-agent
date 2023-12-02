package ipmi

import (
	"fmt"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"kvm-agent/internal/monitor/metrics"
	"kvm-agent/internal/utils"
	"strings"
	"sync"
	"time"
)

func GetSensorStat(name, host, port, username, password string) *metrics.IPMISensorStat {
	resp := &metrics.IPMISensorStat{
		Name:      name,
		IP:        host,
		Port:      port,
		Timestamp: time.Now().Unix(),
	}

	sensorString, err := utils.GetIPMISensorList(host, port, username, password)
	if err != nil {
		log.Errorf("GetSensorStat", "utils.GetIPMISensorList error: %v", err)

		return resp
	}

	result := parseOutput(sensorString)

	sensorDataList := metrics.SensorDataList{}
	sensorDataList.SensorData = make([]*metrics.SensorData, 0)

	for _, data := range result {
		sensor := parseSensorData(string(data))
		sensorDataList.SensorData = append(sensorDataList.SensorData, &sensor)
	}

	resp.SensorDataList = sensorDataList

	return resp
}

func parseSensorData(data string) metrics.SensorData {
	var sensor metrics.SensorData

	// CPU1_VDD         | 0.865      | Volts      | ok    | 0.730     | 0.745     | 0.765     | 0.935     | 0.950     | 0.980
	fmt.Sscanf(data, "%s | %f | %s | %s | %f | %f | %f | %f | %f | %f",
		&sensor.Name,
		&sensor.Value,
		&sensor.Unit,
		&sensor.Status,
		&sensor.NonrecoverableAlarmLow,
		&sensor.CriticalAlarmLow,
		&sensor.WarningAlarmLow,
		&sensor.WarningAlarmHigh,
		&sensor.CriticalAlarmHigh,
		&sensor.NonrecoverableAlarmHigh,
	)

	return sensor
}

func GetAllSensorStat(confs []config.IPMI, timeout int) []*metrics.IPMISensorStat {
	ipmiCount := len(confs)
	sensorStats := make([]*metrics.IPMISensorStat, 0)

	var wg sync.WaitGroup
	wg.Add(ipmiCount)

	for i, conf := range confs {
		go func(i int, c config.IPMI) {
			defer wg.Done()

			stat := GetSensorStat(c.Name, c.IP, fmt.Sprintf("%d", c.Port), c.Username, c.Password)
			sensorStats = append(sensorStats, stat)
		}(i, conf)
	}

	timeoutEvent := time.Duration(timeout) * time.Second
	select {
	case <-time.After(timeoutEvent):
		log.Errorf("GetAllSensorStat", "timeout :%s", timeoutEvent.String())

		return sensorStats
	}
}

func GetSelStat(name, host, port, username, password string) *metrics.IPMISelStat {
	resp := &metrics.IPMISelStat{
		Name:      name,
		IP:        host,
		Port:      port,
		Timestamp: time.Now().Unix(),
	}

	selString, err := utils.GetIPMISelList(host, port, username, password)
	if err != nil {
		log.Errorf("GetSelStat", "utils.GetIPMISelList error: %v", err)

		return resp
	}

	result := parseOutput(selString)

	selDataList := metrics.SelDataList{}
	selDataList.SelData = make([]*metrics.SelData, 0)

	for _, data := range result {
		sel := parseSelData(string(data))
		selDataList.SelData = append(selDataList.SelData, &sel)
	}

	resp.SelDataList = selDataList

	return resp
}

func parseSelData(data string) metrics.SelData {
	var sel metrics.SelData

	// 1 | 2021/07/12 | 01时20分05秒 CST | Event Logging Disabled #0x01 | Log area reset/cleared | Asserted
	lines := strings.Split(data, "|")
	if len(lines) != 6 {
		return sel
	}

	sel.Id = strings.TrimSpace(lines[0])
	sel.Date = strings.TrimSpace(lines[1])
	sel.Time = strings.TrimSpace(lines[2])
	sel.Event = strings.TrimSpace(lines[3])
	sel.Desc = strings.TrimSpace(lines[4])
	sel.Status = strings.TrimSpace(lines[5])

	return sel
}

func GetAllSelStat(confs []config.IPMI, timeout int) []*metrics.IPMISelStat {
	ipmiCount := len(confs)
	selStats := make([]*metrics.IPMISelStat, 0)

	var wg sync.WaitGroup
	wg.Add(ipmiCount)

	for i, conf := range confs {
		go func(i int, c config.IPMI) {
			defer wg.Done()

			stat := GetSelStat(c.Name, c.IP, fmt.Sprintf("%d", c.Port), c.Username, c.Password)
			selStats = append(selStats, stat)
		}(i, conf)
	}

	timeoutEvent := time.Duration(timeout) * time.Second
	select {
	case <-time.After(timeoutEvent):
		log.Errorf("GetAllSelStat", "timeout :%s", timeoutEvent.String())

		return selStats
	}
}

// parseOutput parses the output of the ipmi-sensors command.
func parseOutput(output string) []string {
	lines := strings.Split(output, "\n")
	data := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "Get Channel Cipher Suites") {
			data = append(data, line)
		}
	}

	return data
}
