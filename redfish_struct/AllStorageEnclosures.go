package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
)

type AllStorageEnclosures struct {
	ODataID       string   `json:"@odata.id"`
	ODataType     string   `json:"@odata.type"`
	Description   string   `json:"description"`
	Name          string   `json:"name"`
	Members       []Member `json:"members"`
	Members_count int      `json:"members@odata.count"`
}

func (allStorageEnclosures *AllStorageEnclosures) UnmarshalJson(str string) (*AllStorageEnclosures, error) {
	t, _ := config.GOFISH.Get(str)
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp AllStorageEnclosures

	err := json.Unmarshal(bodyBytes, allStorageEnclosures)
	if err != nil {
		panic(err)
	}
	return allStorageEnclosures, nil
}
