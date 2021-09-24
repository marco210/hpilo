package redfishstruct

import (
	"encoding/json"
	"hpilo_exporter/config"
	"io/ioutil"
	"log"
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
	t, err_resp := config.GOFISH.Get(str)
	if err_resp != nil {
		log.Fatal("err:", err_resp)
		//panic(err_resp)
	}
	defer t.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(t.Body)

	//var temp AllBaseNetworkAdapter

	err := json.Unmarshal(bodyBytes, allBaseNetworkAdapter)
	if err != nil {
		log.Fatal("err:", err)
		return nil, err
	}
	return allBaseNetworkAdapter, err
}
