package metrics

// TriggerInfo
type TriggerInfo struct {
	UUID      string `json:"uuid"`
	Timestamp int64  `json:"timestamp"`

	Key   string `json:"key"`
	Value string `json:"value"`
}
