package redfishstruct

type AllPhysicalDrives struct {
	ODataID       string   `json:"@odata.id"`
	ODataType     string   `json:"@odata.type"`
	Description   string   `json:"description"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	Members_count int      `json:"members@odata.count"`
}
