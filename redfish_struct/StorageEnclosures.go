package redfishstruct

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
