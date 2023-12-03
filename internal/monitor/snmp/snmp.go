package snmp

import (
	"errors"
	"fmt"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"kvm-agent/internal/monitor/metrics"
	"kvm-agent/internal/utils"
	"strings"
	"sync"
	"time"

	"github.com/gosnmp/gosnmp"
)

func GetSNMPStat(name, host, port, community string) *metrics.SNMPStat {
	resp := &metrics.SNMPStat{
		Name:      name,
		IP:        host,
		Port:      port,
		Timestamp: time.Now().Unix(),
	}

	// Get translate list
	translateList := &metrics.OIDTranslateList{}
	translateList.OIDTranslates = make([]metrics.OIDInstance, 0)
	translateString, err := utils.GetSNMPTranslateList(host, port, community, ".")
	if err != nil {
		log.Errorf("GetSNMPStat", "utils.GetSNMPWalkList error: %v", err)

		return resp
	}
	result := parseOutput(translateString)
	for _, data := range result {
		translate := parseTranslateData(string(data))
		// fmt.Printf("GetSNMPStat", "translate: %#v\n", translate)
		translateList.OIDTranslates = append(translateList.OIDTranslates, translate)
	}

	// fmt.Printf("GetSNMPStat", "translateString: %v", translateString)
	// log.Infof("GetSNMPStat", "translateList: %v", translateList)

	// Get oidInstanceMapList
	oidInstanceMapList := metrics.OIDInstanceMapList{}
	oidInstanceMapList.OIDInstanceMap = make([]metrics.OIDInstanceMap, 0)

	// Get oid values
	gosnmp.Default.Target = host
	gosnmp.Default.Port = uint16(utils.StringToInt(port))
	gosnmp.Default.Community = community
	gosnmp.Default.Timeout = time.Duration(10 * time.Second) // Timeout better suited to walking
	err = gosnmp.Default.Connect()
	if err != nil {
		log.Errorf("GetSNMPStat", "gosnmp.Default.Connect error: %v", err)

		return resp
	}
	defer gosnmp.Default.Conn.Close()

	for _, oids := range translateList.OIDTranslates {
		// BulkWalkAll may cause loop timeout
		resultPUDs, err := gosnmp.Default.BulkWalkAll(oids.OID)
		if err != nil {
			log.Errorf("GetSNMPStat", "gosnmp.Default.BulkWalkAll error: %v in oids: %v", err, oids)

			continue
		}

		// Add all results to oidInstanceMapList
		for _, pdu := range resultPUDs {
			pduValue, err := filterValue(pdu)
			if err != nil {
				log.Errorf("GetSNMPStat", "filterValue error: %v in pdu: %v", err, pdu)

				continue
			}

			oidInstanceMap := metrics.OIDInstanceMap{
				OIDInstance: metrics.OIDInstance{
					Name: oids.Name,
					OID:  pdu.Name,
				},
				Value: pduValue,
			}

			oidInstanceMapList.OIDInstanceMap = append(oidInstanceMapList.OIDInstanceMap, oidInstanceMap)
		}

	}

	resp.OIDInstanceMapList = oidInstanceMapList

	return resp
}

func filterValue(pdu gosnmp.SnmpPDU) (string, error) {
	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		// fmt.Printf("STRING: %s\n", string(b))
		return string(b), nil
	default:
		// fmt.Printf("TYPE %d: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
		return fmt.Sprintf("%d", gosnmp.ToBigInt(pdu.Value)), nil
	}
	return "", errors.New("not support type")
}

func parseTranslateData(data string) metrics.OIDInstance {
	// Waring!!! data may be with \t!!!
	// "snmpCommunityGroup"                    "1.3.6.1.6.3.18.2.2.1"
	// "snmpProxyTrapForwardGroup"                     "1.3.6.1.6.3.18.2.2.3"
	// "snmpv2tm"                      "1.3.6.1.6.3.19"
	// "zeroDotZero"                   "0.0"

	// Waring!!! data may be with \t!!!
	data = strings.ReplaceAll(data, "\t", "")
	data = strings.ReplaceAll(data, "\"\"", "-")
	data = strings.ReplaceAll(data, "\"", "")
	// change to snmpCommunityGroup-0.0

	pairs := strings.Split(data, "-")

	if len(pairs) != 2 {
		return metrics.OIDInstance{}
	}

	return metrics.OIDInstance{
		Name: strings.Trim(pairs[0], "\""),
		OID:  strings.Trim(pairs[1], "\""),
	}
}

func parseOutput(output string) []string {
	lines := strings.Split(output, "\n")
	data := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			data = append(data, line)
		}
	}

	return data
}

func GetAllSNMPStat(confs []config.SNMP, timeout int) []*metrics.SNMPStat {
	snmpCount := len(confs)
	snmpStats := make([]*metrics.SNMPStat, 0)

	var wg sync.WaitGroup
	wg.Add(snmpCount)

	for i, conf := range confs {
		go func(i int, c config.SNMP) {
			defer wg.Done()

			stat := GetSNMPStat(c.Name, c.IP, fmt.Sprintf("%d", c.Port), c.Community)
			snmpStats = append(snmpStats, stat)
		}(i, conf)
	}

	timeoutEvent := time.Duration(timeout) * time.Second
	select {
	case <-time.After(timeoutEvent):
		{
			log.Errorf("GetAllSNMPStat", "timeout :%s", timeoutEvent.String())
			return snmpStats
		}
	case <-waitGroupTimeout(&wg):
		{
			return snmpStats
		}
	}
}

// waitGroupTimeout returns a channel that will be closed when the waitgroup is done or the timeout is reached.
func waitGroupTimeout(wg *sync.WaitGroup) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	return ch
}
