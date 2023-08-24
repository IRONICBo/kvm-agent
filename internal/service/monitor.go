package service

import (
	"context"
	"fmt"
	"kvm-agent/internal/dal/dao"
	"kvm-agent/internal/log"
)

const MONITOR_PERFIX = "monitor:"

type MonitorService struct {
	Service

	MonitorDao *dao.MonitorDao
}

// NewMonitorService return new service with context.
func NewMonitorService(c context.Context) *MonitorService {
	return &MonitorService{
		Service: Service{
			ctx: c,
		},
		MonitorDao: dao.NewMonitorDao(),
	}
}

// GuestMonitorPush update guest monitor info.
func (s *MonitorService) GuestMonitorPush(uuid, data string, retry int) error {
	log.Infof("GuestMonitorPush", "%s data: %s", uuid, data)

	// Push to redis.
	// check list length, if length > 100, wait forever.
	err := s.MonitorDao.PushListWithRetry(fmt.Sprintf("%s%s", MONITOR_PERFIX, uuid), data, retry)
	if err != nil {
		return err
	}

	return nil
}
