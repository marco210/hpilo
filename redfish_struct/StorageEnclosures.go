package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
	"log"
)

type StorageEnclosures struct {
	ODataID         string          `json:"@odata.id"`
	ODataType       string          `json:"@odata.type"`
	Id              string          `json:"id"`
	Description     string          `json:"description"`
	DriveBayCount   int             `json:"drivebaycount"`
	FirmwareVersion FirmwareVersion `json:"firmwareversion"`
	Location        string          `json:"location"`
	LocationFormat  string          `json:"locationformat"`
	Model           string          `json:"model"`
	Name            string          `json:"name"`
	SerialNumber    string          `json:"serialnumber"`
	Status          Status          `json:"status"`
}

func (storageEnclosures *StorageEnclosures) UnmarshalJson(str string) (*StorageEnclosures, error) {
	t, err_resp := config.GOFISH.Get(str)
	if err_resp != nil {
		log.Fatal("err:", err_resp)
	}
	defer t.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp StorageEnclosures

	err := json.Unmarshal(bodyBytes, storageEnclosures)
	if err != nil {
		log.Fatal("err:", err)
		return nil, err
	}
	return storageEnclosures, err
}
