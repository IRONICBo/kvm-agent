package main

import (
	"flag"
	"kvm-agent/internal/config"
	"kvm-agent/internal/conn"
	"kvm-agent/internal/log"
	"kvm-agent/internal/models"
	"kvm-agent/internal/utils"
)

// migrate tables.
func main() {
	configPath := flag.String("c", "../../config.yaml", "config file path")
	isDrop := flag.Bool("d", false, "drop table if exist")
	flag.Parse()

	config.ConfigInit(*configPath)
	utils.KVMAgentBanner()
	log.InitLogger(config.Config.App)
	conn.InitDMDB(config.Config.DM)

	// get db instance.
	db := conn.GetDMDB()

	// tables
	tables := []interface{}{
		models.GuestInfo{},
	}

	// drop tables if exist.
	if *isDrop {
		for i := 0; i < len(tables); i++ {
			if db.Migrator().HasTable(&tables[i]) && db.Migrator().DropTable(&tables[i]) != nil {
				log.Panicf("KVM-Agent Table Migration", "Drop table %T... failed", tables[i])
			}
			log.Infof("KVM-Agent Table Migration", "Drop table %T... ok", tables[i])
		}
	}

	// migrate tables.
	for i := 0; i < len(tables); i++ {
		if err := db.AutoMigrate(&tables[i]); err != nil {
			log.Panicf("KVM-Agent Table Migration", "Migrate table %v... failed. Err:%s", tables[i], err.Error())
		}
		log.Infof("KVM-Agent Table Migration", "Migrate table %T... ok", tables[i])
	}
}
