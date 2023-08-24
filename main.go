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
	conn.InitDMDB(config.Config.DM, config.Config.App.Debug)
	conn.InitRedisDB(config.Config.Redis)

	defer func() {
		log.GetLogger().Sync()
	}()

	tasks.InitGuestInfo(config.Config.Agent)
	tasks.RegisterGuestAgentOffline(config.Config.Agent)
	go tasks.StartGuestMonitorTask(config.Config.Agent, config.Config.Agent.GZip)

	for {
	}
}
