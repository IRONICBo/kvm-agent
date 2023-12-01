package metrics

// IPMISensorStat ipmi sensor stat.
type IPMISensorStat struct {
	Name      string `json:"name"`
	IP        string `json:"ip"`
	Port      string `json:"port"`
	Timestamp int64  `json:"timestamp"`

	SensorDataList SensorDataList `json:"sensor_data_list"`
}

// SensorDataList sensor data list.
type SensorDataList struct {
	SensorData []*SensorData `json:"sensor_data"`
}

// SensorData sensor data.
type SensorData struct {
	Name                    string  `json:"name"`
	Value                   float32 `json:"value"`
	Unit                    string  `json:"unit"` // 单位
	Status                  string  `json:"status"`
	NonrecoverableAlarmLow  float32 `json:"nonrecoverable_alarm_low"`  // 低不可恢复
	CriticalAlarmLow        float32 `json:"critical_alarm_low"`        // 低临界
	WarningAlarmLow         float32 `json:"warning_alarm_low"`         // 低警告
	WarningAlarmHigh        float32 `json:"warning_alarm_high"`        // 高警告
	CriticalAlarmHigh       float32 `json:"critical_alarm_high"`       // 高临界
	NonrecoverableAlarmHigh float32 `json:"nonrecoverable_alarm_high"` // 高不可恢复
}

// IPMISelStat ipmi sel stat.
type IPMISelStat struct {
	Name      string `json:"name"`
	IP        string `json:"ip"`
	Port      string `json:"port"`
	Timestamp int64  `json:"timestamp"`

	SelDataList SelDataList `json:"sel_data_list"`
}

// SelDataList sel data list.
type SelDataList struct {
	IPMIName string     `json:"ipmi_name"`
	IPMIIP   string     `json:"ipmi_ip"`
	SelData  []*SelData `json:"sel_data"`
}

// SelData sel data.
type SelData struct {
	Id     string `json:"id"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	Event  string `json:"event"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}
