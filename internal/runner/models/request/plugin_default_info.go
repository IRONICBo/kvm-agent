package request

type PingPlug struct {
	Host     string `json:"host"`
	Count    int    `json:"count"`
	Size     int    `json:"size"`
	Interval int    `json:"interval"`
}

type FIOPlug struct {
	Dir   string `json:"dir"`
	Batch string `json:"batch"`
	Size  string `json:"size"`
}

type DBTestPlug struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Count    int    `json:"count"`
}

type PTP4LPlug struct {
	Interface string `json:"interface"`
}
