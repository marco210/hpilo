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
}

func (chass Chassis) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	chassisArr, chassisErr := metric.Chassis()

	if nil != chassisErr {
		panic(chassisErr)
	}

	for _, chassis := range chassisArr {
		chass.collectPowerLineInputVoltage(ch, chassis)
	}
}

func (chasiss Chassis) collectPowerLineInputVoltage(ch chan<- prometheus.Metric, chass *redfish.Chassis) {
	powers, _ := chass.Power()

	if nil != powers {
		supplies := powers.PowerSupplies

		for _, supply := range supplies {
			ch <- prometheus.MustNewConstMetric(config.C_power_line_input_voltage,
				prometheus.GaugeValue,
				float64(supply.LineInputVoltage),
				supply.MemberID,
				fmt.Sprintf("%v", supply.LineInputVoltageType),
			)
		}

	}
}

func (chassis Chassis) collectTemperature(ch chan<- prometheus.Metric, chass *redfish.Chassis) {
	thermals, _ := chass.Thermal()

	if nil != thermals {
		temperatures := thermals.Temperatures

		for _, temp := range temperatures {
			ch <- prometheus.MustNewConstMetric(config.C_temperature_status,
				prometheus.GaugeValue,
				float64(temp.ReadingCelsius),
				temp.MemberID,
				fmt.Sprintf("%v", temp.SensorNumber),
				string(temp.Status.Health),
				string(temp.Status.State),
				fmt.Sprintf("%v", temp.UpperThresholdCritical),
				fmt.Sprintf("%v", temp.UpperThresholdFatal),
			)
		}
	}
}
