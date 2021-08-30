package config

import (
	"github.com/stmcginnis/gofish"
)

var (
	// GOFISH is global variable
	GOFISH *gofish.APIClient // gofish/client.go -> APIClient

	// Status map
	Status = map[string]float64{"OK": 0.0}

	// hpilo info
	hpilouser = "username"
	// hpilopassword info
	hpilopassword = "password"

	// Map status -> number
	State_dict = map[string]float64{
		"OK":       0.0,
		"WARNING":  1.0,
		"CRITICAL": 2.0,
	}
)
