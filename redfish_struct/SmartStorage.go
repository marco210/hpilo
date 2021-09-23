package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
	"log"
)

type Status struct {
	Health string `json:"health"`
	State  string `json:"state"`
}

type HostBusAdapters struct {
	ODataID string `json:"@odata.id"`
}

type Link_smart_storage struct {
	ArrayControllers ArrayControllers `json:"arraycontrollers"`
	HostBusAdapters  HostBusAdapters  `json:"hostbusadapters"`
}

type SmartStorage struct {
	ODataContext string             `json:"@odata.context"`
	ODataType    string             `json:"@odata.type"`
	Id           string             `json:"id"`
	Description  string             `json:"description"`
	Links        Link_smart_storage `json:"links"`
	Name         string             `json:"name"`
	Status       Status             `json:"status"`
}

func (smartStorage *SmartStorage) UnmarshalJson(str string) (*SmartStorage, error) {
	t, err_resp := config.GOFISH.Get(str)
	if err_resp != nil {
		log.Fatal("err:", err_resp)
	}
	defer t.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp SmartStorage

	err := json.Unmarshal(bodyBytes, smartStorage)
	if err != nil {
		log.Fatal("err:", err)
	}
	return smartStorage, nil
}
