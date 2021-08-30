package config

import "github.com/prometheus/client_golang/prometheus"

var (
	C_power_line_input_voltage = prometheus.NewDesc(
		"idrac_power_line_input_voltage",
		"Power Line Input Voltage",
		[]string{
			"member_id",
			"line_input_voltage_type",
		},
		nil)
)
