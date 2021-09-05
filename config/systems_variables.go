package config

import "github.com/prometheus/client_golang/prometheus"

var (
	//System health metric
	S_health = prometheus.NewDesc(
		"hpilo_system_health_status",
		"hpilo_system_health {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"bios_version",
			"indicator_LED",
			"manufacturer",
			"model",
			"name",
			"power_state",
			"sku",
			"serial_number",
			"status_health",
			"status_state",
			"system_type",
			"uuid",
		},
		nil,
	)

	// S_memory => system's memory
	S_memory = prometheus.NewDesc(
		"idrac_system_memory_status",
		"System Memory {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"bus_width_bits",
			"cache_size_MiB",
			"capacity_MiB",
			"data_width_bits",
			"device_locator",
			"logical_size_MiB",
			"memory_location_channel",
			"memory_location_memorycontroller",
			"memory_location_slot",
			"memory_location_socket",
			"memory_type",
			"name",
			"status_health",
			"status_state",
		},
		nil,
	)

	// S_processor => system's processor
	S_processor = prometheus.NewDesc(
		"idrac_system_processor_status",
		"System processor {0: OK, 1:Warning, 2: Critical}",
		[]string{
			"actions",
			"description",
			"manufacturer",
			"max_speed_MHz",
			"max_td_watts",
			"model",
			"processor_type",
			"socket",
			"sub_processors",
			"td_watts",
			"total_cores",
			"total_enabled_cores",
			"total_threads",
			"uuid",
		},
		nil,
	)

	// S_bios => system's bios
	S_bios = prometheus.NewDesc(
		"idrac_system_bios",
		"System bios",
		[]string{
			"attribute_registry",
			"description",
		},
		nil,
	)
)
