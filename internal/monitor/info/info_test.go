package info

import (
	"fmt"
	"kvm-agent/internal/utils"
	"testing"
)

func TestGetInfos(t *testing.T) {
	cpuinfo, _ := utils.DecompressText([]byte(GetCpuInfoJsonCompressed()))
	fmt.Println(cpuinfo)
	meminfo, _ := utils.DecompressText([]byte(GetMemInfoJsonCompressed()))
	fmt.Println(meminfo)
	diskinfo, _ := utils.DecompressText([]byte(GetDiskInfoJsonCompressed()))
	fmt.Println(diskinfo)
	netinfo, _ := utils.DecompressText([]byte(GetNetInfoJsonCompressed()))
	fmt.Println(netinfo)
}

func TestGetOriginalData(t *testing.T) {
	data := `{"interface_infos":[{"index":0,"mtu":65536,"name":"lo","hardwareaddr":"","flags":["up","loopback"],"addrs":[{"addr":"127.0.0.1/8"},{"addr":"::1/128"}]},{"index":0,"mtu":1500,"name":"enaphyt4i0","hardwareaddr":"00:07:3e:a7:29:1d","flags":["up","broadcast","multicast"],"addrs":[]},{"index":0,"mtu":1500,"name":"vnet1","hardwareaddr":"fe:54:00:58:de:82","flags":["up","broadcast","multicast"],"addrs":[{"addr":"fe80::fc54:ff:fe58:de82/64"}]},{"index":0,"mtu":1500,"name":"virbr0","hardwareaddr":"52:54:00:59:ff:e4","flags":["up","broadcast","multicast"],"addrs":[{"addr":"192.168.122.1/24"}]},{"index":0,"mtu":1500,"name":"virbr0-nic","hardwareaddr":"52:54:00:59:ff:e4","flags":["broadcast","multicast"],"addrs":[]},{"index":0,"mtu":1500,"name":"br0","hardwareaddr":"00:07:3e:a7:29:1d","flags":["up","broadcast","multicast"],"addrs":[{"addr":"192.168.0.75/24"},{"addr":"fe80::8b0f:7103:a9cc:57af/64"}]},{"index":0,"mtu":1500,"name":"vnet0","hardwareaddr":"fe:54:00:a8:8c:28","flags":["up","broadcast","multicast"],"addrs":[{"a
	ddr":"fe80::fc54:ff:fea8:8c28/64"}]},{"index":0,"mtu":1500,"name":"vnet2","hardwareaddr":"fe:54:00:d9:32:f4","flags":["up","broadcast","multicast"],"addrs":[{"addr":"fe80::fc54:ff:fed9:32f4/64"}]},{"index":0,"mtu":1500,"name":"vnet3","hardwareaddr":"fe:54:00:0a:45:1d","flags":["up","broadcast","multicast"],"addrs":[{"addr":"fe80::fc54:ff:fe0a:451d/64"}]}]}`

	compressedBytes, _ := utils.CompressText(data)
	// b64 := "H4sIAAAAAAAA/6yS746bMBDE32U/u7A2GMy+ShVVi/80qAQiQtJWp7z7CS6HThDliMg3a7We/Wlm3qBqet8Ftv5X1YT2BPRzmDn/DwgFHPozUKZ1kglo+OCBoG5BwJ4795c7z851QAACQs2/h99wPoKAum2PJds/sBMw7Hzo3ralyiOMMJKxgauYxkQylsrAdTcMZwxSI04IvuHj/n+fVrhEQSTMKfHEOamCpFuwlV3LzvKpBwGHc91X4/sL6Lf3L43v5fJ08KRTQiRtyHky6vnTkxnBGyQKVqcUAgU/ShoVZ+kKfy5VV3Z3vNHqE7AYVH26AVAWKpKZiaRSkYzVeqwfTWWfRNue2F07XlCVhR0Y5Xp0Q8yiNCUGyiUmxIW1pHMOK8NsfH8HfiobGzKWlHlh2UZJZdbzqQd8rqBEUdjStTnfKBnS9XzJAz5kSvW28Od8o6R0N77d9T0AAP//vBwmtmYFAAA="
	b64 := utils.Base64Encode(compressedBytes)
	b64Decode, _ := utils.Base64Decode(b64)
	fmt.Println(string(b64Decode))

	original, _ := utils.DecompressText(b64Decode)
	fmt.Println(original)
}

func TestEncodeAndDecode(t *testing.T) {
	jsonData := `{"interface_infos":[{"index":0,"mtu":65536,"name":"lo","hardwareaddr":"","flags":["up","loopback"],"addrs":[{"addr":"127.0.0.1/8"},{"addr":"::1/128"}]},{"index":0,"mtu":1500,"name":"enaphyt4i0","hardwareaddr":"00:07:3e:a7:29:1d","flags":["up","broadcast","multicast"],"addrs":[]},{"index":0,"mtu":1500,"name":"vnet1","hardwareaddr":"fe:54:00:58:de:82","flags":["up","broadcast","multicast"],"addrs":[{"addr":"fe80::fc54:ff:fe58:de82/64"}]},{"index":0,"mtu":1500,"name":"virbr0","hardwareaddr":"52:54:00:59:ff:e4","flags":["up","broadcast","multicast"],"addrs":[{"addr":"192.168.122.1/24"}]},{"index":0,"mtu":1500,"name":"virbr0-nic","hardwareaddr":"52:54:00:59:ff:e4","flags":["broadcast","multicast"],"addrs":[]},{"index":0,"mtu":1500,"name":"br0","hardwareaddr":"00:07:3e:a7:29:1d","flags":["up","broadcast","multicast"],"addrs":[{"addr":"192.168.0.75/24"},{"addr":"fe80::8b0f:7103:a9cc:57af/64"}]},{"index":0,"mtu":1500,"name":"vnet0","hardwareaddr":"fe:54:00:a8:8c:28","flags":["up","broadcast","multicast"],"addrs":[{"addr":"fe80::fc54:ff:fea8:8c28/64"}]},{"index":0,"mtu":1500,"name":"vnet2","hardwareaddr":"fe:54:00:d9:32:f4","flags":["up","broadcast","multicast"],"addrs":[{"addr":"fe80::fc54:ff:fed9:32f4/64"}]},{"index":0,"mtu":1500,"name":"vnet3","hardwareaddr":"fe:54:00:0a:45:1d","flags":["up","broadcast","multicast"],"addrs":[{"addr":"fe80::fc54:ff:fe0a:451d/64"}]}]}`

	encodedString, err := utils.CompressAndEncodeBase64(jsonData)
	if err != nil {
		fmt.Println("CompressAndEncodeBase64 failed:", err)
		return
	}

	fmt.Println("CompressAndEncodeBase64:", encodedString)

	decodedString, err := utils.DecodeAndDecompressBase64("H4sIAAAAAAAA/1zQ0Y6rIBAG4HeZayVIkQIvY+YopkRkDKA57qbvvrFk02bvmO+Hn2S+4UG5DD7ONOSCBWyViKsDC4FGDNfMXqeJVvQRGti34q8LhpteCd7AP6IyVOuUvt2V7nvRwJZozGBvRjZA+Sr0cf8PDWwBy0xpBQvLGV6dvzTMuPpwgoVPPFzKnmLVxaXowodJ1hlmeNsLJgQ7hOB3tpwdZ4hpfCj5fnPNYOHth09lx+C/sHiKQz5zcWv95k+UKLgaXBvxE1hQTo9CC932UvFWGhStMWha5SbNDTo5dRM8nz8BAAD//8g04NxmAQAA")
	if err != nil {
		fmt.Println("DecodeAndDecompressBase64 failed:", err)
		return
	}

	fmt.Println("DecodeAndDecompressBase64:", decodedString)
}
