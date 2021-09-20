package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
)

type AllBaseNetworkAdapter struct {
	ODataID       string   `json:"@odata.id"`
	ODataType     string   `json:"@odata.type"`
	Description   string   `json:"description"`
	Name          string   `json:"name"`
	Members       []Member `json:"members"`
	Members_count int      `json:"members@odata.count"`
}

func (allBaseNetworkAdapter *AllBaseNetworkAdapter) UnmarshalJson(str string) (*AllBaseNetworkAdapter, error) {
	t, _ := config.GOFISH.Get(str)
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp AllBaseNetworkAdapter

	err := json.Unmarshal(bodyBytes, allBaseNetworkAdapter)
	if err != nil {
		panic(err)
	}
	return allBaseNetworkAdapter, nil
}
