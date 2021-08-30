package collector

import (
	"fmt"
	"hpilo_exporter/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
	//	"github.com/stmcginnis/gofish/redfish"
)

type SystemCollector struct{}

func (collector SystemCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.S_bios
}

func (collector SystemCollector) collectSystemHealth(ch chan<- prometheus.Metric, v *redfish.ComputerSystem) {
	status := config.State_dict[string(v.Status.Health)]
	ch <- prometheus.MustNewConstMetric(config.S_health, prometheus.GaugeValue, status,
		v.BIOSVersion,
		v.Description,
		v.HostName,
		v.Manufacturer,
		v.Model,
		v.Name,
		v.PartNumber,
		fmt.Sprintf("%v", v.PowerRestorePolicy),
		fmt.Sprintf("%v", v.PowerState),
		v.SKU,
		v.SerialNumber,
		v.SubModel,
		fmt.Sprintf("%v", v.SystemType),
		v.UUID,
	)
}

func (collector SystemCollector) collectBios(ch chan<- prometheus.Metric, system *redfish.ComputerSystem) {
	val, biosErr := system.Bios()

	if nil != biosErr {
		panic(biosErr)
	}

	ch <- prometheus.MustNewConstMetric(config.S_bios,
		prometheus.GaugeValue,
		float64(0),
		val.AttributeRegistry,
		val.Description,
	)
}
