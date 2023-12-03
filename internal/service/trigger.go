package service

import (
	"context"
	"kvm-agent/internal/dal/dao"
	"kvm-agent/internal/log"
)

const TRIGGER_PERFIX = "trigger:"

type TriggerService struct {
	Service

	TriggerDao *dao.TriggerDao
}

// NewTriggerrService return new service with context.
func NewTriggerrService(c context.Context) *TriggerService {
	return &TriggerService{
		Service: Service{
			ctx: c,
		},
		TriggerDao: dao.NewTriggerDao(),
	}
}

// GuestTriggerPush update guest trigger info.
func (s *TriggerService) GuestTriggerPush(uuid, value string) error {
	// log.Infof("GuestTriggerPush", "%s value: %s", uuid, value)

	// Push to redis.
	// check list length, if length > 20, wait forever.
	err := s.TriggerDao.PushListWithRetry(TRIGGER_PERFIX+uuid, value, 10, 60)
	if err != nil {
		log.Errorf("GuestTriggerPush", "PushListWithRetry error: %s", err.Error())

		return err
	}

	return nil
}
