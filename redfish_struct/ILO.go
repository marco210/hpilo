package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
)

type Ipv4Address struct {
	Address       string `json:"address"`
	AddressOrigin string `json:"addressorigin"`
	Gateway       string `json:"gateway"`
	SubnetMask    string `json:"subnetmask"`
}

type VLAN struct {
	VLANEnable bool   `json:"vlanenable"`
	VLANId     string `json:"vlanid"`
}

type ILOPort struct {
	ODataID       string        `json:"@odata.id"`
	ODataType     string        `json:"@odata.type"`
	Description   string        `json:"description"`
	ID            string        `json:"id"`
	FullDuplex    bool          `json:"fullduplex"`
	HostName      string        `json:"hostname"`
	IPv4Addresses []Ipv4Address `json:"ipv4addresses"`
	LinkStatus    string        `json:"linkstatus"`
	MACAddress    string        `json:"macaddress"`
	Name          string        `json:"name"`
	SpeedMbps     int           `json:"speedmbps"`
	Status        Status        `json:"status"`
	VLAN          VLAN          `json:"vlan"`
}

func (iloport *ILOPort) UnmarshalJson(str string) (error, *ILOPort) {
	t, _ := config.GOFISH.Get(str)
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	var temp ILOPort

	err := json.Unmarshal(bodyBytes, &temp)
	if err != nil {
		panic(err)
	}
	return nil, &temp
}