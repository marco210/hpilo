package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
)

type ControllerBoard struct {
	Status Status `json:"status"`
}

type Current struct {
	VersionString string `json:"versionstring"`
}

type FirmwareVersion struct {
	Current Current `json:"current"`
}

type UnconfiguredDrives struct {
	ODataID string `json:"@odata.id"`
}

type Link_smart_storage_arr struct {
	LogicalDrives LogicalDrives `json:"logicaldrives"`
}

type ArrayControllers struct {
	ODataID                             string                 `json:"@odata.id"`
	ODataType                           string                 `json:"@odata.type"`
	Id                                  string                 `json:"id"`
	AdapterType                         string                 `json:"adaptertype"`
	BackupPowerSourceStatus             string                 `json:"backuppowersourcestatus"`
	ControllerBoard                     ControllerBoard        `json:"controllerboard"`
	ControllerPartNumber                string                 `json:"controllerpartnumber"`
	CurrentOperatingMode                string                 `json:"currentoperatingmode"`
	Description                         string                 `json:"description"`
	DriveWriteCache                     string                 `json:"drivewritecache"`
	ExternalPortCount                   int                    `json:"externalportcount"`
	FirmwareVersion                     FirmwareVersion        `json:"firmwareversion"`
	HardwareRevision                    string                 `json:"hardwarerevision"`
	InternalPortCount                   int                    `json:"internalportcount"`
	Links                               Link_smart_storage_arr `json:"links"`
	Location                            string                 `json:"location"`
	LocationFormat                      string                 `json:"locationformat"`
	Model                               string                 `json:"model"`
	Name                                string                 `json:"name"`
	ReadCachePercent                    int                    `json:"readcachepercent"`
	SerialNumber                        string                 `json:"serialnumber"`
	Status                              Status                 `json:"status"`
	WriteCacheWithoutBackupPowerEnabled bool                   `json:"writecachewithoutbackuppowerenabled"`
}

func (arrayControllers *ArrayControllers) UnmarshalJson(str string) (*ArrayControllers, error) {
	t, _ := config.GOFISH.Get(str)
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp ArrayControllers

	err := json.Unmarshal(bodyBytes, arrayControllers)
	if err != nil {
		panic(err)
	}
	return arrayControllers, nil
}
