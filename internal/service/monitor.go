package service

import (
	"context"
	"fmt"
	"kvm-agent/internal/dal/dao"
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
func (s *MonitorService) GuestMonitorPush(uuid, data string, interval int) error {
	// log.Infof("GuestMonitorPush", "%s data: %s", uuid, data)

	// Push to redis.
	// check list length, if length > 20, wait forever.
	err := s.MonitorDao.PushListWithRetry(fmt.Sprintf("%s%s", MONITOR_PERFIX, uuid), data, 10, interval)
	if err != nil {
		return err
	}

	return nil
}
