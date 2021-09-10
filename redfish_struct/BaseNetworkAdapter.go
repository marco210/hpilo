package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
)

type HpePhysicalPort struct {
	ODataContext  string `json:"@odata.context"`
	ODataType     string `json:"@odata.type"`
	BadReceives   int    `json:"badreceives"`
	BadTransmits  int    `json:"badtransmits"`
	GoodReceives  int    `json:"goodreceives"`
	GoodTransmits int    `json:"goodtransmits"`
	Team          string `json:"team"`
}

type OemPhysicalPort struct {
	HpePhysicalPort HpePhysicalPort `json:"hpe"`
}

type PhysicalPort struct {
	FullDuplex      bool            `json:"fullduplex"`
	IPv4Addresses   []string        `json:"ipv4address"`
	IPv6Addresses   []string        `json:"ipv6addresses"`
	LinkStatus      string          `json:"linkstatus"`
	MacAddress      string          `json:"macaddress"`
	OemPhysicalPort OemPhysicalPort `json:"oem"`
	SpeedMbps       int             `json:"speedmbps"`
	Status          Status          `json:"status"`
}

type BaseNetworkAdapter struct {
	ODataID         string          `json:"@odata.id"`
	ODataType       string          `json:"@odata.type"`
	ID              string          `json:"id"`
	FirmwareVersion FirmwareVersion `json:"firmware"`
	Name            string          `json:"name"`
	PartNumber      string          `json:"partnumber"`
	PhysicalPorts   []PhysicalPort  `json:"physicalports"`
	SerialNumber    string          `json:"serialnumber"`
	Status          Status          `json:"status"`
}

func (baseNetworkAdapter *BaseNetworkAdapter) UnmarshalJson(str string) (error, *BaseNetworkAdapter) {
	t, _ := config.GOFISH.Get(str)
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	var temp BaseNetworkAdapter

	err := json.Unmarshal(bodyBytes, &temp)
	if err != nil {
		panic(err)
	}
	return nil, &temp
}
