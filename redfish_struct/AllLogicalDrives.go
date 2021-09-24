package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
	"log"
)

type AllLogicalDrives struct {
	ODataID       string   `json:"@odata.id"`
	ODataType     string   `json:"@odata.type"`
	Description   string   `json:"description"`
	Name          string   `json:"name"`
	Members       []Member `json:"members"`
	Members_count int      `json:"members@odata.count"`
}

func (allLogicalDrives *AllLogicalDrives) UnmarshalJson(str string) (*AllLogicalDrives, error) {
	t, err_resp := config.GOFISH.Get(str)
	if err_resp != nil {
		log.Fatal("err:", err_resp)
	}
	defer t.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp AllLogicalDrives

	err := json.Unmarshal(bodyBytes, allLogicalDrives)
	if err != nil {
		log.Fatal("err:", err)
		return nil, err
	}
	return allLogicalDrives, err
}
