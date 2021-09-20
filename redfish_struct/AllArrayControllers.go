package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
)

type AllArrayController struct {
	ODataID       string   `json:"@odata.id"`
	ODataType     string   `json:"@odata.type"`
	Description   string   `json:"description"`
	Name          string   `json:"name"`
	Members       []Member `json:"members"`
	Members_count int      `json:"members@odata.count"`
}

func (allArrayController *AllArrayController) UnmarshalJson(str string) (*AllArrayController, error) {
	t, _ := config.GOFISH.Get(str)
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp AllArrayController

	err := json.Unmarshal(bodyBytes, allArrayController)
	if err != nil {
		panic(err)
	}
	return allArrayController, nil
}
