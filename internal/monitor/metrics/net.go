package metrics

type NetInfo struct {
	InterfaceInfos []InterfaceInfo `json:"interface_infos"`
}

// NetInterfaceAddr net interface addr.
type InterfaceAddr struct {
	Addr string `json:"addr"`
}

// InterfaceInfo net interface info.
type InterfaceInfo struct {
	Index        int             `json:"index"`
	MTU          int             `json:"mtu"`          // maximum transmission unit
	Name         string          `json:"name"`         // e.g., "en0", "lo0", "eth0.100"
	HardwareAddr string          `json:"hardwareaddr"` // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        []string        `json:"flags"`        // e.g., FlagUp, FlagLoopback, FlagMulticast
	Addrs        []InterfaceAddr `json:"addrs"`
}

// NetStat net stat.
type NetStat struct {
	ConnectionStats    []ConnectionStat       `json:"connection_stats"`
	ConntrackStats     []ConntrackStat        `json:"conntrack_stats"`
	FilterStats        []FilterStat           `json:"filter_stats"`
	NetIOCountersStats map[string]interface{} `json:"net_io_counters_stats"` // key is interface name
	ProtoCountersStats map[string]interface{} `json:"proto_counters_stats"`  // key is protocol name
}

// Addr net address
type Addr struct {
	IP   string `json:"ip"`
	Port uint32 `json:"port"`
}

// ConnectionStat connection stat.
type ConnectionStat struct {
	Fd     uint32  `json:"fd"`
	Family uint32  `json:"family"`
	Type   uint32  `json:"type"`
	Laddr  Addr    `json:"localaddr"`
	Raddr  Addr    `json:"remoteaddr"`
	Status string  `json:"status"`
	Uids   []int32 `json:"uids"`
	Pid    int32   `json:"pid"`
}

// ConntrackStat has conntrack summary info
type ConntrackStat struct {
	Entries       uint32 `json:"entries"`        // Number of entries in the conntrack table
	Searched      uint32 `json:"searched"`       // Number of conntrack table lookups performed
	Found         uint32 `json:"found"`          // Number of searched entries which were successful
	New           uint32 `json:"new"`            // Number of entries added which were not expected before
	Invalid       uint32 `json:"invalid"`        // Number of packets seen which can not be tracked
	Ignore        uint32 `json:"ignore"`         // Packets seen which are already connected to an entry
	Delete        uint32 `json:"delete"`         // Number of entries which were removed
	DeleteList    uint32 `json:"delete_list"`    // Number of entries which were put to dying list
	Insert        uint32 `json:"insert"`         // Number of entries inserted into the list
	InsertFailed  uint32 `json:"insert_failed"`  // # insertion attempted but failed (same entry exists)
	Drop          uint32 `json:"drop"`           // Number of packets dropped due to conntrack failure.
	EarlyDrop     uint32 `json:"early_drop"`     // Dropped entries to make room for new ones, if maxsize reached
	IcmpError     uint32 `json:"icmp_error"`     // Subset of invalid. Packets that can't be tracked d/t error
	ExpectNew     uint32 `json:"expect_new"`     // Entries added after an expectation was already present
	ExpectCreate  uint32 `json:"expect_create"`  // Expectations added
	ExpectDelete  uint32 `json:"expect_delete"`  // Expectations deleted
	SearchRestart uint32 `json:"search_restart"` // Conntrack table lookups restarted due to hashtable resizes
}

// FilterStat filter stat.
type FilterStat struct {
	ConnTrackCount int64 `json:"conntrackCount"`
	ConnTrackMax   int64 `json:"conntrackMax"`
}

// NetIOCountersStat net io counters stat.
type NetIOCountersStat struct {
	// Name        string `json:"name"`         // interface name
	BytesSent   uint64 `json:"bytes_sent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytes_recv"`   // number of bytes received
	PacketsSent uint64 `json:"packets_sent"` // number of packets sent
	PacketsRecv uint64 `json:"packets_recv"` // number of packets received
	Errin       uint64 `json:"errin"`        // total number of errors while receiving
	Errout      uint64 `json:"errout"`       // total number of errors while sending
	Dropin      uint64 `json:"dropin"`       // total number of incoming packets which were dropped
	Dropout     uint64 `json:"dropout"`      // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
	Fifoin      uint64 `json:"fifoin"`       // total number of FIFO buffers errors while receiving
	Fifoout     uint64 `json:"fifoout"`      // total number of FIFO buffers errors while sending
}

// ProtoCountersStat proto counters stat.
type ProtoCountersStat struct {
	Protocol string           `json:"protocol"`
	Stats    map[string]int64 `json:"stats"`
}
