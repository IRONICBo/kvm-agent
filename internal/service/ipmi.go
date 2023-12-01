package service

import (
	"context"
	"fmt"
	"kvm-agent/internal/dal/dao"
	"kvm-agent/internal/log"
)

const IPMI_SENSOR_PERFIX = "ipmi:sensor:"
const IPMI_SEL_PERFIX = "ipmi:sel"

type IPMIService struct {
	Service

	MonitorDao *dao.MonitorDao
}

// NewIPMIService return new service with context.
func NewIPMIService(c context.Context) *IPMIService {
	return &IPMIService{
		Service: Service{
			ctx: c,
		},
		MonitorDao: dao.NewMonitorDao(),
	}
}

// IPMISensorMonitorPush update ipmi sensor monitor info.
func (s *IPMIService) IPMISensorMonitorPush(uuid, data string, interval int) error {
	log.Infof("IPMIMonitorPush", "%s data: %s", uuid, data)

	// Push to redis.
	// check list length, if length > 20, wait forever.
	err := s.MonitorDao.PushListWithRetry(fmt.Sprintf("%s%s", IPMI_SENSOR_PERFIX, uuid), data, 10, interval)
	if err != nil {
		return err
	}

	return nil
}

// IPMISelMonitorPush update ipmi sel monitor info.
func (s *IPMIService) IPMISelMonitorPush(uuid, data string, interval int) error {
	log.Infof("IPMIMonitorPush", "%s data: %s", uuid, data)

	// Push to redis.
	// check list length, if length > 20, wait forever.
	err := s.MonitorDao.PushListWithRetry(fmt.Sprintf("%s%s", IPMI_SEL_PERFIX, uuid), data, 10, interval)
	if err != nil {
		return err
	}

	return nil
}
