package service

import (
	"context"
	"kvm-agent/internal/config"
	"kvm-agent/internal/dal/dao"
	"kvm-agent/internal/models"
	"kvm-agent/internal/monitor/info"

	"gorm.io/gorm"
)

type GuestService struct {
	Service

	GuestInfoDao *dao.GuestInfoDao
}

// NewGuestService return new service with context.
func NewGuestService(c context.Context) *GuestService {
	return &GuestService{
		Service: Service{
			ctx: c,
		},
		GuestInfoDao: dao.NewGuestInfoDao(),
	}
}

// CreateGuestInfo create guest info.
func (svc *GuestService) CreateGuestInfo(uuid string) (string, int, error) {
	online := true

	guestInfo := models.GuestInfo{
		UUID:     uuid,
		HostDesc: info.GetHostInfoJsonCompressed(),
		CpuDesc:  info.GetCpuInfoJsonCompressed(),
		MemDesc:  info.GetMemInfoJsonCompressed(),
		DiskDesc: info.GetDiskInfoJsonCompressed(),
		NetDesc:  info.GetNetInfoJsonCompressed(),
		Period:   config.Config.Agent.Period, // Period can not be 0
		UseGzip:  &config.Config.Agent.GZip,
		IsOnline: &online,
	}

	err := svc.GuestInfoDao.Create(&guestInfo)
	if err != nil {
		return "", 0, err
	}

	return uuid, 0, nil
}

// UpdateGuestInfo update guest info.
func (svc *GuestService) UpdateGuestInfo(uuid string) (string, int, error) {
	online := true

	// check guest info exist
	guestInfo, err := svc.GuestInfoDao.FindFirstByUUID(uuid)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", 0, err
	}

	if guestInfo == nil {
		return svc.CreateGuestInfo(uuid)
	}

	guestInfo.HostDesc = info.GetHostInfoJsonCompressed()
	guestInfo.CpuDesc = info.GetCpuInfoJsonCompressed()
	guestInfo.MemDesc = info.GetMemInfoJsonCompressed()
	guestInfo.DiskDesc = info.GetDiskInfoJsonCompressed()
	guestInfo.NetDesc = info.GetNetInfoJsonCompressed()
	guestInfo.UseGzip = &config.Config.Agent.GZip
	guestInfo.IsOnline = &online

	err = svc.GuestInfoDao.Update(guestInfo)
	if err != nil {
		return "", 0, err
	}

	return uuid, 0, nil
}

// UpdateGuestAgentOffline update guest agent offline.
func (svc *GuestService) UpdateGuestAgentOffline(uuid string) error {
	online := false

	guestInfo, err := svc.GuestInfoDao.FindFirstByUUID(uuid)
	if err != nil {
		return err
	}

	guestInfo.IsOnline = &online

	err = svc.GuestInfoDao.Update(guestInfo)
	if err != nil {
		return err
	}

	return nil
}
