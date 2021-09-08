package redfishstruct

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
