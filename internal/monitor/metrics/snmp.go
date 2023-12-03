package metrics

// SNMPStat snmp stat.
type SNMPStat struct {
	Name      string `json:"name"`
	IP        string `json:"ip"`
	Port      string `json:"port"`
	Timestamp int64  `json:"timestamp"`

	OIDInstanceMapList OIDInstanceMapList `json:"oid_instance_map_list"`
}

// OIDTranslateList
type OIDTranslateList struct {
	OIDTranslates []OIDInstance `json:"oid_translates"`
}

// OIDInstanceMapList
type OIDInstanceMapList struct {
	OIDInstanceMap []OIDInstanceMap `json:"oid_instance_map"`
}

// OIDInstanceMap
type OIDInstanceMap struct {
	OIDInstance OIDInstance `json:"oid_instance"`
	Value       string      `json:"value"`
}

// OIDInstance
type OIDInstance struct {
	Name string `json:"name"`
	OID  string `json:"oid"`
}
