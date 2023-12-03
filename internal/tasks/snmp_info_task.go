package tasks

import (
	"context"
	"encoding/json"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"kvm-agent/internal/monitor/status"
	"kvm-agent/internal/service"
	"kvm-agent/internal/utils"
	"time"
)

// SNMP collect interval multiplier.
const SNMP_MULTIPLIER = 10

// StartSNMPTask start snmp task.
func StartSNMPTask(config config.Agent, gzip bool) {
	interval := time.Duration(config.Period) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	svc := service.NewSNMPService(context.Background())

	for {
		select {
		case <-ticker.C:
			// Default set timeout to config.Period
			metric := status.GetAllStat(config.UUID, config.Period)
			metricString, err := json.Marshal(metric)
			if err != nil {
				log.Errorf("StartGuestMonitorTask", "json.Marshal error: %v", err)
			}

			if gzip {
				metricString, err = utils.CompressText(string(metricString))
				if err != nil {
					log.Errorf("StartGuestMonitorTask", "utils.CompressText error: %v", err)
				}
			}

			if err = svc.GuestMonitorPush(config.UUID, string(metricString), config.Period); err != nil {
				log.Errorf("StartGuestMonitorTask", "svc.GuestMonitorPush error: %v", err)
			}
		}
	}
}
