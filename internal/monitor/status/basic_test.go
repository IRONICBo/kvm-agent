package status

import (
	"encoding/json"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"testing"
)

func TestBasicMonitor(t *testing.T) {
	logConfig := config.App{
		LogFile: "kvm-agent.log",
	}
	log.InitLogger(logConfig)

	stat := GetAllStat("6e8c2828-5460-49a2-99a9-6ed809ae4d1d", 5)
	jsonString, _ := json.Marshal(stat)

	log.Infof("GetAllStat", "GetAllStat: %v", string(jsonString))
}
