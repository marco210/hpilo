package config

import "github.com/prometheus/client_golang/prometheus"

var (
	//System health metric
	S_health = prometheus.NewDesc(
		"idrac_system_health_status",
		"idrac_system_health {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"bios_version",
			"description",
			"hostname",
			"hosted_services",
			"manufacturer",
			"model",
			"name",
			"part_number",
			"power_restore_policy",
			"power_state",
			"sku",
			"serial_number",
			"submodel",
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
			"allocation_alignment_MiB",
			"allocation_increment_MiB",
			"base_module_type",
			"bus_width_bits",
			"cache_size_MiB",
			"capacity_MiB",
			"configuration_locked",
			"data_width_bits",
			"description",
			"device_locator",
			"error_correction",
			"firmware_api_version",
			"firmware_revision",
			"is_rank_square_enabled",
			"is_square_device_enabled",
			"logical_size_MiB",
			"manufacturer",
			"memory_device_type",
			"memory_type",
			"operating_speed_Mhz",
			"part_number",
			"rank_count",
			"serial_number",
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
