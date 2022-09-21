package chgdns

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"os"
)

// DnsServers represents the DNS servers
type DnsServers struct {
	DnsServers []DnsServer `json:"servers"`
}

// DnsServer represents a DNS server
type DnsServer struct {
	Name         string  `json:"name"`
	PrimaryDNS   *net.IP `json:"primaryDns"`
	SecondaryDNS *net.IP `json:"secondaryDNS"`
}

// GetName returns the name
func (data DnsServer) GetName() string {
	return data.Name
}

// GetPrimaryDNS returns the primary DNS
func (data DnsServer) GetPrimaryDNS() *net.IP {
	return data.PrimaryDNS
}

// GetSecondaryDNS returns the secondary DNS
func (data DnsServer) GetSecondaryDNS() *net.IP {
	return data.SecondaryDNS
}

// GetDnsServers returns the DNS servers
func (data DnsServers) GetDnsServers() []DnsServer {
	return data.DnsServers
}

// GetDnsServer returns the DNS server
/*func (data DnsServer) GetDnsServer() []DnsServer {
	return []DnsServer{data}
}*/

// ReadDnsServers reads the DNS servers from the file
func ReadDnsServers() (DnsServers, error) {
	var err error
	var servers DnsServers

	jsonFile, err := os.Open("dns_servers.json")
	if err != nil {
		panic(err)
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
		}
	}(jsonFile)

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &servers)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(servers.DnsServers); i++ {
		servers.GetDnsServers()[i].GetName()
		servers.GetDnsServers()[i].GetPrimaryDNS()
		servers.GetDnsServers()[i].GetSecondaryDNS()
	}

	return servers, nil
}
