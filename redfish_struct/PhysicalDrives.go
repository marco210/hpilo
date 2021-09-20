package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
)

type PhysicalDrives struct {
	ODataID                           string          `json:"@odata.id"`
	ODataType                         string          `json:"@odata.type"`
	Id                                string          `json:"id"`
	BlockSizeBytes                    int             `json:"blocksizebytes"`
	CapacityGB                        int             `json:"capacitygb"`
	CapacityLogicalBlocks             int             `json:"capacitylogicalblocks"`
	CapacityMiB                       int             `json:"capacitymib"`
	CarrierApplicationVersion         string          `json:"carrierapplicationversion"`
	CarrierAuthenticationStatus       string          `json:"carrierauthenticationstatus"`
	CurrentTemperatureCelsius         int             `json:"currenttemperaturecelsius"`
	Description                       string          `json:"description"`
	DiskDriveStatusReasons            []string        `json:"diskdriveStatusreasons"`
	DiskDriveUse                      string          `json:"diskdriveuse"`
	EncryptedDrive                    bool            `json:"encrypteddrive"`
	FirmwareVersion                   FirmwareVersion `json:"firmwareversion"`
	InterfaceSpeedMbps                int             `json:"interfacespeedmbps"`
	InterfaceType                     string          `json:"interfacetype"`
	LegacyBootPriority                string          `json:"legacybootpriority"`
	Location                          string          `json:"location"`
	LocationFormat                    string          `json:"locationformat"`
	MaximumTemperatureCelsius         int             `json:"maximumtemperaturecelsius"`
	MediaType                         string          `json:"mediatype"`
	Model                             string          `json:"model"`
	Name                              string          `json:"name"`
	PowerOnHours                      int             `json:"poweronhours"`
	SSDEnduranceUtilizationPercentage int             `json:"ssdenduranceutilizationpercentage"`
	SerialNumber                      string          `json:"serialnumber"`
	Status                            Status          `json:"status"`
	UncorrectedReadErrors             int             `json:"uncorrectedreaderrors"`
	UncorrectedWriteErrors            int             `json:"uncorrectedwriteerrors"`
}

func (physicalDrives *PhysicalDrives) UnmarshalJson(str string) (*PhysicalDrives, error) {
	t, _ := config.GOFISH.Get(str)
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp PhysicalDrives

	err := json.Unmarshal(bodyBytes, physicalDrives)
	if err != nil {
		panic(err)
	}
	return physicalDrives, nil
}
