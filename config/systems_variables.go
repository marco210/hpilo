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
		"hpilo_system_memory_status",
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
		"hpilo_system_processor_status",
		"System processor {0: OK, 1:Warning, 2: Critical}",
		[]string{
			"id",
			"instruction_set",
			"manufacturer",
			"max_speed_MHz",
			"model",
			"processor_architecture",
			"processor_type",
			"status_health",
			"status_state",
			"total_cores",
			"total_enabled_cores",
			"total_threads",
		},
		nil,
	)

	S_network_interface = prometheus.NewDesc(
		"hpilo_network_interface_status",
		"System processor {0: OK, 1:Warning, 2: Critical}",
		[]string{
			"id",
			"instruction_set",
			"manufacturer",
			"max_speed_MHz",
			"model",
			"processor_architecture",
			"processor_type",
			"status_health",
			"status_state",
			"total_cores",
			"total_enabled_cores",
			"total_threads",
		},
		nil,
	)

	S_storage_array_status = prometheus.NewDesc(
		"hpilo_storage_array_status",
		"hpilo storage array controller status",
		[]string{
			"id",
			"adapter_type",
			"controller_board_status_health",
			"controller_partnumber",
			"escription",
			"firmware_version",
			"internal_port_count",
			"location",
			"location_format",
			"model",
			"name",
			"status_health",
			"status_state",
		},
		nil,
	)
)
