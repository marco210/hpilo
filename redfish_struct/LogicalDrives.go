package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
	"log"
)

type DataDrives struct {
	ODataID string `json:"@odata.id"`
}

type Link_logical_drive struct {
	DataDrives DataDrives `json:"datadrives"`
}

type LogicalDrives struct {
	ODataID                   string             `json:"@odata.id"`
	ODataType                 string             `json:"@odata.type"`
	Id                        string             `json:"id"`
	AccelerationMethod        string             `json:"accelerationmethod"`
	CapacityMiB               int                `json:"capacitymib"`
	Description               string             `json:"description"`
	InterfaceType             string             `json:"interfacetype"`
	LegacyBootPriority        string             `json:"legacybootpriority"`
	Links                     Link_logical_drive `json:"links"`
	LogicalDriveEncryption    bool               `json:"logicaldriveencryption"`
	LogicalDriveName          string             `json:"logicaldrivename"`
	LogicalDriveNumber        int                `json:"logicaldrivenumber"`
	LogicalDriveStatusReasons []string           `json:"logicaldrivestatusreasons"`
	LogicalDriveType          string             `json:"logicaldrivetype"`
	MediaType                 string             `json:"mediatype"`
	Name                      string             `json:"name"`
	Raid                      string             `json:"raid"`
	Status                    Status             `json:"status"`
	StripeSizeBytes           int                `json:"stripesizebytes"`
	VolumeUniqueIdentifier    string             `json:"volumeuniqueidentifier"`
}

func (logicalDrives *LogicalDrives) UnmarshalJson(str string) (*LogicalDrives, error) {
	t, err_resp := config.GOFISH.Get(str)
	if err_resp != nil {
		log.Fatal("err:", err_resp)
	}
	defer t.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp LogicalDrives

	err := json.Unmarshal(bodyBytes, logicalDrives)
	if err != nil {
		log.Fatal("err:", err)
	}
	return logicalDrives, nil
}
