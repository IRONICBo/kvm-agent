package conn

import (
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"testing"
)

func TestDMConn(t *testing.T) {
	logConfig := config.App{
		LogFile: "kvm-agent.log",
	}
	log.InitLogger(logConfig)

	config := config.DM{
		Ip:       "127.0.0.1",
		Port:     5236,
		Username: "SYSDBA",
		Password: "SYSDBA",
	}

	InitDMDB(config, true)
}
