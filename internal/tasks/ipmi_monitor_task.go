package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"kvm-agent/internal/monitor/ipmi"
	"kvm-agent/internal/service"
	"kvm-agent/internal/utils"
	"time"
)

// IPMI collect interval multiplier.
const IPMI_SENSOR_MULTIPLIER = 1
const IPMI_SEL_MULTIPLIER = 2

func StartIPMIMonitorTask(config config.Agent, ipmisConfig []config.IPMI, gzip bool) {
	interval := time.Duration(config.Period) * time.Second
	sensor_ticker := time.NewTicker(interval * IPMI_SENSOR_MULTIPLIER)
	defer sensor_ticker.Stop()

	sel_ticker := time.NewTicker(interval * IPMI_SEL_MULTIPLIER)
	defer sel_ticker.Stop()

	svc := service.NewIPMIService(context.Background())

	for {
		select {
		case <-sensor_ticker.C:
			{
				log.Debugf("StartIPMIMonitorTask", "sensor_ticker.C")

				// Default set timeout to config.Period
				metric := ipmi.GetAllSensorStat(ipmisConfig, config.Period*IPMI_SENSOR_MULTIPLIER*len(ipmisConfig))
				metricString, err := json.Marshal(metric)
				if err != nil {
					log.Errorf("StartIPMIMonitorTask", "json.Marshal error: %v", err)
				}

				fmt.Printf("metricString: %s", metricString)

				log.Debugf("StartIPMIMonitorTask", "metricString: %s", metricString)

				if gzip {
					metricString, err = utils.CompressText(string(metricString))
					if err != nil {
						log.Errorf("StartIPMIMonitorTask", "utils.CompressText error: %v", err)
					}
				}

				if err = svc.IPMISensorMonitorPush(config.UUID, string(metricString), config.Period); err != nil {
					log.Errorf("StartIPMIMonitorTask", "svc.IPMISensorMonitorPush error: %v", err)
				}
			}
		case <-sel_ticker.C:
			{
				log.Debugf("StartIPMIMonitorTask", "sel_ticker.C")

				// Default set timeout to config.Period
				metric := ipmi.GetAllSelStat(ipmisConfig, config.Period*IPMI_SEL_MULTIPLIER*len(ipmisConfig))
				metricString, err := json.Marshal(metric)
				if err != nil {
					log.Errorf("StartIPMIMonitorTask", "json.Marshal error: %v", err)
				}

				fmt.Printf("metricString: %s", metricString)

				if gzip {
					metricString, err = utils.CompressText(string(metricString))
					if err != nil {
						log.Errorf("StartIPMIMonitorTask", "utils.CompressText error: %v", err)
					}
				}

				if err = svc.IPMISelMonitorPush(config.UUID, string(metricString), config.Period); err != nil {
					log.Errorf("StartIPMIMonitorTask", "svc.GuestMonitorPush error: %v", err)
				}
			}
		}
	}
}
