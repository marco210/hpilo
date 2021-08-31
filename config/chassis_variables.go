package config

import "github.com/prometheus/client_golang/prometheus"

var (
	C_power_line_input_voltage = prometheus.NewDesc(
		"ilo_power_line_input_voltage",
		"Power Line Input Voltage",
		[]string{
			"member_id",
			"line_input_voltage_type",
		},
		nil)

	C_temperature_status = prometheus.NewDesc(
		"ilo_temperature_status",
		"Chassis temperature {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"reading_celsius",
			"memberid",
			"sensor_number",
			"status_health",
			"status_state",
			"upper_threshold_critical",
			"upper_threshold_fatal",
		},
		nil,
	)

	// C_networkadapter => network adapter of the chassis
	C_networkadapter = prometheus.NewDesc(
		"ilo_chassis_network_adapter",
		"Chassis network adapter {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"description",
			"manufacturer",
			"model",
			"part_number",
			"sku",
			"serial_number",
		},
		nil,
	)
)
