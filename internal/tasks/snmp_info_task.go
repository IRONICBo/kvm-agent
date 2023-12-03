package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"kvm-agent/internal/monitor/snmp"
	"kvm-agent/internal/service"
	"kvm-agent/internal/utils"
	"time"
)

// SNMP collect interval multiplier.
const SNMP_MULTIPLIER = 1

// StartSNMPTask start snmp task.
func StartSNMPTask(config config.Agent, snmpConfig []config.SNMP, gzip bool) {
	interval := time.Duration(config.Period) * time.Second
	ticker := time.NewTicker(interval * SNMP_MULTIPLIER)
	defer ticker.Stop()

	svc := service.NewSNMPService(context.Background())

	for {
		select {
		case <-ticker.C:
			// Default set timeout to config.Period
			metric := snmp.GetAllSNMPStat(snmpConfig, config.Period*IPMI_SEL_MULTIPLIER*10)
			metricString, err := json.Marshal(metric)
			if err != nil {
				log.Errorf("StartSNMPTask", "json.Marshal error: %v", err)
			}
			fmt.Printf("StartSNMPTask metricString: %v", string(metricString))

			if gzip {
				metricString, err = utils.CompressText(string(metricString))
				if err != nil {
					log.Errorf("StartSNMPTask", "utils.CompressText error: %v", err)
				}
			}

			if err = svc.GuestSNMPPush(config.UUID, string(metricString), config.Period); err != nil {
				log.Errorf("StartSNMPTask", "svc.GuestSNMPPush error: %v", err)
			}
		}
	}
}
