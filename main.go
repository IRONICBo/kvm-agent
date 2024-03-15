package main

import (
	"flag"
	"kvm-agent/internal/config"
	"kvm-agent/internal/conn"
	"kvm-agent/internal/log"
	"kvm-agent/internal/tasks"
	"kvm-agent/internal/utils"
)

func main() {
	configPath := flag.String("c", "./config.yaml", "config file path")
	flag.Parse()

	// init
	config.ConfigInit(*configPath)
	utils.KVMAgentBanner()
	log.InitLogger(config.Config.App)
	defer func() {
		log.GetLogger().Sync()
	}()

	if !config.Config.App.BanMonitor {
		conn.InitDMDB(config.Config.DM, config.Config.App.Debug)
		conn.InitRedisDB(config.Config.Redis)
		tasks.InitGuestInfo(config.Config.Agent)
		tasks.RegisterGuestAgentOffline(config.Config.Agent)
		go tasks.StartGuestMonitorTask(config.Config.Agent, config.Config.Agent.GZip)
		go tasks.StartGuestTriggerTask(config.Config.Agent, config.Config.Agent.GZip)
		if config.Config.Hardware.IPMI_Enable {
			go tasks.StartIPMIMonitorTask(config.Config.Agent, config.Config.IPMI, config.Config.Agent.GZip)
		}
		if config.Config.Hardware.SNMP_Enable {
			go tasks.StartSNMPTask(config.Config.Agent, config.Config.SNMP, config.Config.Agent.GZip)
		}
	}

	go tasks.StartGuestPluginTask(config.Config.Server)

	for {
	}
}
