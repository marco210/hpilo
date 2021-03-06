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
		"hpilo_memory_status",
		"System Memory {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"bus_width_bits",
			"cache_size_MiB",
			"capacity_MiB",
			"data_width_bits",
			"device_locator",
			"logical_size_MiB",
			"memory_device_type",
			"memory_location_channel",
			"memory_location_memorycontroller",
			"memory_location_slot",
			"memory_location_socket",
			"memory_type",
			"name",
			"part_number",
			"status_health",
			"status_state",
		},
		nil,
	)

	// S_processor => system's processor
	S_processor = prometheus.NewDesc(
		"hpilo_processor_status",
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

	// // S_networkport => system's network port
	// S_networkport = prometheus.NewDesc(
	// 	"hpilo_system_network_port_status",
	// 	"System Network Port",
	// 	[]string{
	// 		"adapter_manufacturer",
	// 		"link_status",
	// 		"current_link_speed_mbps",
	// 		"description",
	// 		"max_frame_size",
	// 		"number_discovered_remote_ports",
	// 		"physical_port_number",
	// 		"port_maximum_mtu",
	// 	},
	// 	nil,
	// )

	// // S_network_adapter_status => system network adapter status
	// S_network_adapter_status = prometheus.NewDesc(
	// 	"hpilo_network_adapter_status",
	// 	"System Controller {0: OK, 1: Warning, 2: Critical}",
	// 	[]string{
	// 		"network_adapter_manufacture",
	// 		"firmware_package_version",
	// 		"network_device_functions_count",
	// 		"network_ports_count",
	// 	},
	// 	nil,
	// )
	// S_ethernetinterface => system's ethernet interface
	S_ethernetinterface = prometheus.NewDesc(
		"hpilo_system_ethernet_interface_status",
		"System Ethernet Interface{0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"full_duplex",
			"ipv4addresses",
			"ipv4_static_addresses",
			"ipv6_addresses",
			"ipv6_static_addresses",
			"mac_address",
			"link_status",
			"speed_Mbps",
			"status_health",
			"status_state",
		},
		nil,
	)

	S_base_network_adapter_status = prometheus.NewDesc(
		"hpilo_network_adapter_status",
		"System base network adapter status{0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"firmware",
			"name",
			"part_number",
			"physicalport_linkstatus",
			"physicalport_mac",
			"serial_number",
			"status_health",
			"status_state",
		},
		nil,
	)

	//status
	S_base_network_port_adapter_status = prometheus.NewDesc(
		"hpilo_network_port_status",
		"System network port{0: OK, 1: Warning, 2: Critical, None}",
		[]string{
			"id_adapter",
			"port_number",
			"ipv4address",
			"macaddress",
			"link_status",
			"bad_receives",
			"bad_transmits",
			"good_receives",
			"good_transmits",
			"speed_mbps",
			"status_health",
		},
		nil,
	)
	//hpilo_network_port_bad_receives
	S_base_network_port_bad_receives = prometheus.NewDesc(
		"hpilo_network_port_bad_receives",
		"A count of frames that were received by the adapter but which had an error",
		[]string{
			"id_adapter",
			"port_number",
			"ipv4address",
			"macaddress",
			"link_status",
			"bad_receives",
			"bad_transmits",
			"good_receives",
			"good_transmits",
			"speed_mbps",
			"status_health",
		},
		nil,
	)
	//hpilo_network_port_good_receives
	S_base_network_port_good_receives = prometheus.NewDesc(
		"hpilo_network_port_good_receives",
		"A count of frames successfully received by the physical adapter",
		[]string{
			"id_adapter",
			"port_number",
			"ipv4address",
			"macaddress",
			"link_status",
			"bad_receives",
			"bad_transmits",
			"good_receives",
			"good_transmits",
			"speed_mbps",
			"status_health",
		},
		nil,
	)
	//hpilo_network_port_bad_transmits
	S_base_network_port_bad_transmits = prometheus.NewDesc(
		"hpilo_network_port_bad_transmits",
		"A count of frames that were not transmitted by the adapter because of an error",
		[]string{
			"id_adapter",
			"port_number",
			"ipv4address",
			"macaddress",
			"link_status",
			"bad_receives",
			"bad_transmits",
			"good_receives",
			"good_transmits",
			"speed_mbps",
			"status_health",
		},
		nil,
	)

	//hpilo_network_port_good_transmits
	S_base_network_port_good_transmits = prometheus.NewDesc(
		"hpilo_network_port_good_transmits",
		"A count of frames successfully transmitted by the physical adapter",
		[]string{
			"id_adapter",
			"port_number",
			"ipv4address",
			"macaddress",
			"link_status",
			"bad_receives",
			"bad_transmits",
			"good_receives",
			"good_transmits",
			"speed_mbps",
			"status_health",
		},
		nil,
	)
	S_ilo_status = prometheus.NewDesc("hpilo_ilo_port_status",
		"{0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"full_duplex_ilo",
			"hostname_ilo",
			"address_ilo",
			"address_origin",
			"gateway_ilo",
			"subnetmask_ilo",
			"speed_mbps",
			"status_health",
			"status_state",
			"vlan_enable",
			"vlan_id",
		},
		nil,
	)

	// S_storage => systems' storage
	S_storage_physical_drive_status = prometheus.NewDesc(
		//"hpiloc_system_physic_drive_status",
		"hpilo_physical_drive_status",
		"System storage physic_drive {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"block_Size_bytes",
			"capicity_gb",
			"capacity_logical_blocks",
			"capacity_MiB",
			"carrier_authentication_status",
			"current_temperature_celsius",
			"description",
			"disk_drive_use",
			"interface_speedMbps",
			"interface_type",
			"maximum_temperature_celsius",
			"media_type",
			"model",
			"name",
			"power_on_hours",
			"serial_number",
			"status_health",
			"status_state",
		},
		nil,
	)

	S_storage_array_controller_status = prometheus.NewDesc(
		//"hpilo_system_storage_array_controller_status",
		"hpilo_array_controller_status",
		"System storage array controller {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"adapter_type",
			"controller_board_status_health",
			"controller_partnumber",
			"description",
			"firmware_version",
			"external_port_count",
			"hardware_revision",
			"internal_port_count",
			"location",
			"location_format",
			"model",
			"name",
			"read_cache_percent",
			"serial_number",
			"status_health",
			"status_state",
		},
		nil,
	)

	S_storage_logical_drive_status = prometheus.NewDesc(
		//"hpilo_system_storage_logical_drive_status",
		"hpilo_logical_drive_status",
		"System storage logical drive {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"acceleration_method",
			"capacity_MiB",
			"description",
			"interface_type",
			"legacy_boot_priority",
			"logical_drive_encryption",
			"logical_drive_name",
			"logical_drive_number",
			"logical_drive_status_reasons",
			"logical_drive_type",
			"media_type",
			"name",
			"raid",
			"stripe_size_bytes",
			"volume_unique_identifier",
			"status_health",
			"status_state",
		},
		nil,
	)

	S_storage_enclosures_status = prometheus.NewDesc(
		//"hpilo_system_storage_enclosures_status",
		"hpilo_storage_enclosure_status",
		"System storage enclosures  {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"id",
			"description",
			"drive_bay_count",
			"firmware_version",
			"location",
			"location_format",
			"model",
			"name",
			"serial_number",
			"status_health",
			"status_state",
		},
		nil,
	)
)
