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

	// S_network_interface = prometheus.NewDesc(
	// 	"hpilo_network_interface_status",
	// 	"System processor {0: OK, 1:Warning, 2: Critical}",
	// 	[]string{
	// 		"id",
	// 		"instruction_set",
	// 		"manufacturer",
	// 		"max_speed_MHz",
	// 		"model",
	// 		"processor_architecture",
	// 		"processor_type",
	// 		"status_health",
	// 		"status_state",
	// 		"total_cores",
	// 		"total_enabled_cores",
	// 		"total_threads",
	// 	},
	// 	nil,
	// )

	S_network_interfaces_status = prometheus.NewDesc(
		"hpilo_network_interfaces_status",
		"hpilo network interfaces status",
		[]string{
			"id",
			"name",
			"part_number",
			"physical_port",
			"serial_number",
			"status_state",
		},
		nil,
	)

	// S_network_adapter_status => system network adapter status
	S_network_adapter_status = prometheus.NewDesc(
		"idrac_network_adapter_status",
		"System Controller {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"network_adapter_manufacture",
			"firmware_package_version",
			"network_device_functions_count",
			"network_ports_count",
		},
		nil,
	)

	// S_ethernetinterface => system's ethernet interface
	S_ethernetinterface = prometheus.NewDesc(
		"idrac_system_ethernet_interface_status",
		"System Ethernet Interface{0: OK, 1: Warning, 2: Critical}",
		[]string{
			"auto_negotiation",
			"description",
			"ethernet_interface_type",
			"fqdn",
			"full_duplex",
			"host_name",
			"mac_address",
			"mtu_size",
			"speed_Mbps",
		},
		nil,
	)

	// S_networkport => system's network port
	S_networkport = prometheus.NewDesc(
		"system_network_port_status",
		"System Network Port",
		[]string{
			"adapter_manufacturer",
			"link_status",
			"current_link_speed_mbps",
			"description",
			"max_frame_size",
			"number_discovered_remote_ports",
			"physical_port_number",
			"port_maximum_mtu",
		},
		nil,
	)
)
