package models

type GuestInfo struct {
	Model

	UUID     string `json:"uuid" gorm:"index;not null;comment:UUID"`
	CpuDesc  string `json:"cpu_desc" gorm:"not null;comment:CpuDesc"`
	MemDesc  string `json:"mem_desc" gorm:"not null;comment:MemDesc"`
	DiskDesc string `json:"disk_desc" gorm:"not null;comment:DiskDesc"`
	NetDesc  string `json:"net_desc" gorm:"not null;comment:NetDesc"`
	UseGzip  bool   `json:"use_gzip" gorm:"not null;default:0;comment:UseGzip,1:enable,0:disable"`
	IsOnline bool   `json:"is_online" gorm:"not null;default:1;comment:IsOnline,1:enable,0:disable"`
}
