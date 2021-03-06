package config

import "github.com/prometheus/client_golang/prometheus"

var (
	Chasis_status = prometheus.NewDesc(
		"hpilo_chasis_status",
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
		"hpilo_power_line_input_voltage",
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

	C_power_supply_status = prometheus.NewDesc(
		"hpilo_power_supply_status",
		"Power Supply Status",
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
		"hpilo_power_control",
		"Power Capacity watts",
		[]string{
			"member_id",
			"power_capacity_watts",
			"power_consumed_watts",
			"average_consumed_watts",
			"max_consumed_watts",
			"min_consumed_watts",
		},
		nil,
	)
	C_power_consume_by_all = prometheus.NewDesc(
		"hpilo_power_consumed_by_all",
		"The actual power being consumed by the server in Watts",
		[]string{
			"member_id",
			"power_capacity_watts",
			"power_consumed_watts",
		},
		nil,
	)

	C_power_consume_by_each = prometheus.NewDesc(
		"hpilo_power_consumed_by_each",
		"(LastPowerOutputWatts) The latest observed average power being drawn by the power supply (Watts)",
		[]string{
			"member_id",
			"power_capacity_watts",
			"LastPowerOutputWatts",
		},
		nil,
	)

	C_fans_status = prometheus.NewDesc(
		"hpilo_fan_status",
		"status of fans",
		[]string{
			"member_id",
			"name",
			"reading",
			"status",
			"state",
		},
		nil,
	)

	C_fans_reading = prometheus.NewDesc(
		"hpilo_fan_speed",
		"The current speed of the fan",
		[]string{
			"member_id",
			"name",
			"reading",
			"status",
			"state",
		},
		nil,
	)

	C_temperature_status = prometheus.NewDesc(
		"hpilo_temperature_status",
		"Chassis temperature {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"member_id",
			"name",
			"reading_celsius",
			"sensor_number",
			"status_health",
			"status_state",
			"upper_threshold_critical",
			"upper_threshold_fatal",
		},
		nil,
	)

	C_temperature_reading = prometheus.NewDesc(
		"hpilo_temperature_reading",
		"The current reading of the temperature sensor in Celsius",
		[]string{
			"member_id",
			"name",
			"reading_celsius",
			"sensor_number",
			"status_health",
			"status_state",
			"upper_threshold_critical",
			"upper_threshold_fatal",
		},
		nil,
	)
)
