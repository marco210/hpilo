package config

import "github.com/prometheus/client_golang/prometheus"

var (
	S_health = prometheus.NewDesc(
		"hpilo_system_health_status",
		"hpilo_system_health {0:OK, 1: Warning, 2: Critical}",
		[]string{
			"BiosVersion",
			"HostName",
			"Manufacturer",
			"Model",
			"Name",
			"Power_state",
			"SKU",
			"Serial_Number",
			"Status",
			"UUID",
		},
		nil,
	)

	S_bios = prometheus.NewDesc(
		"ilo_system_bios",
		"System bios",
		[]string{
			"attribute_registry",
			"description",
		},
		nil,
	)
)
