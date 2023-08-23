package main

import (
	"flag"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
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
}
