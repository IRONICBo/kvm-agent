package service

import (
	"context"
	"fmt"
	"kvm-agent/internal/dal/dao"
	"kvm-agent/internal/log"
)

const SNMP_PERFIX = "snmp:"

type SNMPService struct {
	Service

	MonitorDao *dao.MonitorDao
}

// NewSNMPService return new service with context.
func NewSNMPService(c context.Context) *SNMPService {
	return &SNMPService{
		Service: Service{
			ctx: c,
		},
		MonitorDao: dao.NewMonitorDao(),
	}
}

// GuestMonitorPush update guest monitor info.
func (s *MonitorService) GuestSNMPPush(uuid, data string, interval int) error {
	log.Infof("GuestSNMPPush", "%s data: %s", uuid, data)

	// Push to redis.
	// check list length, if length > 20, wait forever.
	err := s.MonitorDao.PushListWithRetry(fmt.Sprintf("%s%s", SNMP_PERFIX, uuid), data, 10, interval)
	if err != nil {
		return err
	}

	return nil
}
