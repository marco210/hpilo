package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
	"log"
)

type Member struct {
	MemberOID string `json:"@odata.id"`
}

type AllPhysicalDrives struct {
	ODataID       string   `json:"@odata.id"`
	ODataType     string   `json:"@odata.type"`
	Description   string   `json:"description"`
	Name          string   `json:"name"`
	Members       []Member `json:"members"`
	Members_count int      `json:"members@odata.count"`
}

func (allPhysicalDrives *AllPhysicalDrives) UnmarshalJson(str string) (*AllPhysicalDrives, error) {
	t, err_resp := config.GOFISH.Get(str)
	if err_resp != nil {
		log.Fatal("err:", err_resp)
	}
	defer t.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp AllPhysicalDrives

	err := json.Unmarshal(bodyBytes, allPhysicalDrives)
	if err != nil {
		log.Fatal("err:", err)
	}
	return allPhysicalDrives, nil
}
