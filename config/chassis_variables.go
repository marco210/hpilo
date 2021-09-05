package config

import "github.com/prometheus/client_golang/prometheus"

var (
	Chasis_status = prometheus.NewDesc(
		"chasis_status",
		"chasis status",
		[]string{
			"id",
			"indicator_LED",
			"manufacturer",
			"model",
			"name",
			"power_state",
			"sku",
			"serial_number",
			"status",
		},
		nil,
	)
	C_power_line_input_voltage = prometheus.NewDesc(
		"ilo_power_line_input_voltage",
		"Power Line Input Voltage",
		[]string{
			"member_id",
			"line_input_voltage_type",
			"firmware_ersion",
			"last_power_outputWatts",
			"manufacturer",
			"model",
			"power_capacity_Watts",
			"serial_number",
			"status",
		},
		nil,
	)
	C_power_control = prometheus.NewDesc(
		"ilo_power_control",
		"Power Control",
		[]string{
			"member_id",
			"power_capacity_watts",
			"power_consumed_watts",
			"power_metrics",
		},
		nil,
	)

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
