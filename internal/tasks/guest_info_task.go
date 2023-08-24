package tasks

import (
	"context"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"kvm-agent/internal/service"
	"os"
	"os/signal"
	"syscall"
)

// InitGuestInfo Start agent and update guest info.
func InitGuestInfo(config config.Agent) {
	svc := service.NewGuestService(context.Background())

	uuid, uid, err := svc.UpdateGuestInfo(config.UUID)
	if err != nil {

		log.Panicf("InitGuestInfo", "Can not Create/Update guest info, err: %s", err.Error())
	}

	log.Infof("InitGuestInfo", "Create/Update guest info success, uuid: %s, uid: %d", uuid, uid)
}

// RegisterGuestAgentOffline Register guest agent offline event and update status.
func RegisterGuestAgentOffline(config config.Agent) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// update guest agent offline status.
	go func() {
		<-c
		svc := service.NewGuestService(context.Background())
		if err := svc.UpdateGuestAgentOffline(config.UUID); err != nil {
			log.Errorf("RegisterGuestAgentOffline", "Update guest agent offline failed, err: %s", err.Error())
		}
		os.Exit(0)
	}()
}
