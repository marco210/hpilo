package collector

import (
	"fmt"
	"hpilo_exporter/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
)

type Chassis struct{}

func (chassis Chassis) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.C_temperature_status
	ch <- config.C_power_line_input_voltage
	ch <- config.C_power_control
	ch <- config.Chasis_status
	ch <- config.C_fans_status
}

func (chass Chassis) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	chassisArr, chassisErr := metric.Chassis()

	if nil != chassisErr {
		panic(chassisErr)
	}

	for _, chassis := range chassisArr {
		chass.collectPowerLineInputVoltage(ch, chassis)
		chass.collectTemperature(ch, chassis)
		chass.collectChasissStatus(ch, chassis)
		chass.collectFansStatus(ch, chassis)
	}
}

func (chasiss Chassis) collectChasissStatus(ch chan<- prometheus.Metric, chass *redfish.Chassis) {
	chass_temp := string(chass.Status.Health)
	chass_temp1 := 0.0
	if chass_temp == "OK" {
		chass_temp1 = 0
	} else if chass_temp == "WARNING" {
		chass_temp1 = 1
	} else {
		chass_temp1 = 2
	}

	ch <- prometheus.MustNewConstMetric(
		config.Chasis_status,
		prometheus.GaugeValue,
		chass_temp1,
		fmt.Sprintf("%v", chass.ID),
		fmt.Sprintf("%v", chass.IndicatorLED),
		fmt.Sprintf("%v", chass.Manufacturer),
		fmt.Sprintf("%v", chass.Model),
		fmt.Sprintf("%v", chass.Name),
		fmt.Sprintf("%v", chass.PowerState),
		fmt.Sprintf("%v", chass.SKU),
		fmt.Sprintf("%v", chass.SerialNumber),
		fmt.Sprintf("%v", chass.Status.Health),
	)

}

func (chasiss Chassis) collectPowerLineInputVoltage(ch chan<- prometheus.Metric, chass *redfish.Chassis) {
	powers, _ := chass.Power()

	if nil != powers {
		supplies := powers.PowerSupplies

		for _, supply := range supplies {
			// "member_id",
			// "line_input_voltage_type",
			// "firmware_ersion",
			// "last_power_outputWatts",
			// "manufacturer",
			// "model",
			// "power_capacity_Watts",
			// "serial_number",
			// "status",
			ch <- prometheus.MustNewConstMetric(config.C_power_line_input_voltage,
				prometheus.GaugeValue,
				float64(supply.LineInputVoltage),
				supply.MemberID,
				fmt.Sprintf("%v", supply.LineInputVoltageType),
				fmt.Sprintf("%v", supply.FirmwareVersion),
				fmt.Sprintf("%v", supply.LastPowerOutputWatts),
				fmt.Sprintf("%v", supply.Manufacturer),
				fmt.Sprintf("%v", supply.Model),
				fmt.Sprintf("%v", supply.PowerCapacityWatts),
				fmt.Sprintf("%v", supply.SerialNumber),
				fmt.Sprintf("%v", supply.Status.Health),
			)
		}

		pw_controls := powers.PowerControl

		for _, pw_control := range pw_controls {
			ch <- prometheus.MustNewConstMetric(config.C_power_control,
				prometheus.GaugeValue,
				float64(pw_control.PowerCapacityWatts),
				pw_control.MemberID,
				fmt.Sprintf("%v", pw_control.PowerCapacityWatts),
				fmt.Sprintf("%v", pw_control.PowerConsumedWatts),
				fmt.Sprintf("%v", pw_control.PowerMetrics.AverageConsumedWatts),
				fmt.Sprintf("%v", pw_control.PowerMetrics.MaxConsumedWatts),
				fmt.Sprintf("%v", pw_control.PowerMetrics.MinConsumedWatts),
			)
		}

	}
}

func (chassis Chassis) collectFansStatus(ch chan<- prometheus.Metric, chass *redfish.Chassis) {
	thermals, _ := chass.Thermal()
	// 		"member_id",
	// 		"name",
	// 		"reading",
	// 		"status",
	// 		"state",
	if nil != thermals {
		fans := thermals.Fans
		for _, fan := range fans {
			fan_temp := fan.Status.Health
			fan_temp1 := 0.0
			if fan_temp == "OK" {
				fan_temp1 = 0
			} else if fan_temp == "WARNING" {
				fan_temp1 = 1
			} else {
				fan_temp1 = 2
			}
			ch <- prometheus.MustNewConstMetric(
				config.C_fans_status,
				prometheus.GaugeValue,
				fan_temp1,
				fmt.Sprintf("%v", fan.MemberID),
				fmt.Sprintf("%v", fan.Name),
				fmt.Sprintf("%v", fan.Reading),
				fmt.Sprintf("%v", fan.Status.Health),
				fmt.Sprintf("%v", fan.Status.State),
			)
		}
	}
}

func (chassis Chassis) collectTemperature(ch chan<- prometheus.Metric, chass *redfish.Chassis) {
	thermals, _ := chass.Thermal()

	if nil != thermals {
		temperatures := thermals.Temperatures

		for _, temp := range temperatures {
			temperature_temp := temp.Status.Health
			temperature_temp1 := 0.0
			if temperature_temp == "OK" {
				temperature_temp1 = 0
			} else if temperature_temp == "WARNING" {
				temperature_temp1 = 1
			} else {
				temperature_temp1 = 2
			}
			ch <- prometheus.MustNewConstMetric(config.C_temperature_status,
				prometheus.GaugeValue,
				temperature_temp1,
				fmt.Sprintf("%v", temp.MemberID),
				fmt.Sprintf("%v", temp.Name),
				fmt.Sprintf("%v", temp.ReadingCelsius),
				fmt.Sprintf("%v", temp.SensorNumber),
				string(temp.Status.Health),
				string(temp.Status.State),
				fmt.Sprintf("%v", temp.UpperThresholdCritical),
				fmt.Sprintf("%v", temp.UpperThresholdFatal),
			)
		}
	}
}
