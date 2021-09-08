package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
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
	CapacityMiB               int                `json:"capacitymib"`
	Description               string             `json:"description"`
	InterfaceType             string             `json:"interfacetype"`
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
	StripeSizeBytes           string             `json:"stripesizebytes"`
	VolumeUniqueIdentifier    string             `json:"volumeuniqueidentifier"`
}

func (logicalDrives *LogicalDrives) UnmarshalJson(str string) (error, *LogicalDrives) {
	t, _ := config.GOFISH.Get(str)
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	var temp LogicalDrives

	err := json.Unmarshal(bodyBytes, &temp)
	if err != nil {
		panic(err)
	}
	return nil, &temp
}
