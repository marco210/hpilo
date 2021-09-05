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
	ch <- config.S_health
	ch <- config.S_memory
}

func (collector SystemCollector) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, sysErr := metric.Systems() //Systems get the system instances from the service

	if nil != sysErr {
		panic(sysErr)
	}
	for _, system := range systems {
		collector.collectSystemHealth(ch, system)
		collector.collectMemories(ch, system)
	}
}

func (collector SystemCollector) collectSystemHealth(ch chan<- prometheus.Metric, v *redfish.ComputerSystem) {
	system_temp := string(v.Status.Health)
	system_temp1 := 0.0
	if system_temp == "OK" {
		system_temp1 = 0
	} else if system_temp == "WARNING" {
		system_temp1 = 1
	} else {
		system_temp1 = 2
	}
	ch <- prometheus.MustNewConstMetric(
		config.S_health,
		prometheus.GaugeValue,
		system_temp1,
		fmt.Sprintf("%v", v.BIOSVersion),
		string(v.IndicatorLED),
		v.Manufacturer,
		v.Model,
		v.Name,
		string(v.PowerState),
		v.SKU,
		fmt.Sprintf("%v", v.SerialNumber),
		fmt.Sprintf("%v", v.Status.Health),
		fmt.Sprintf("%v", v.Status.State),
		string(v.SystemType),
		v.UUID,
	)
}

func (collector SystemCollector) collectMemories(ch chan<- prometheus.Metric, v *redfish.ComputerSystem) {
	memories, err := v.Memory()

	if nil != err {
		panic(err)
	}

	if err == nil {
		for _, memory := range memories {
			memory_temp := string(memory.Status.Health)
			memory_temp1 := 0.0
			if memory_temp == "OK" {
				memory_temp1 = 0
			} else if memory_temp == "WARNING" {
				memory_temp1 = 1
			} else {
				memory_temp1 = 2
			}

			ch <- prometheus.MustNewConstMetric(
				config.S_memory,
				prometheus.GaugeValue,
				memory_temp1,
				fmt.Sprintf("%v", memory.ID),
				fmt.Sprintf("%v", memory.BusWidthBits),
				fmt.Sprintf("%v", memory.CacheSizeMiB),
				fmt.Sprintf("%v", memory.CapacityMiB),
				fmt.Sprintf("%v", memory.DataWidthBits),
				fmt.Sprintf("%v", memory.DeviceLocator),
				fmt.Sprintf("%v", memory.LogicalSizeMiB),
				fmt.Sprintf("%v", memory.MemoryLocation.Channel),
				fmt.Sprintf("%v", memory.MemoryLocation.MemoryController),
				fmt.Sprintf("%v", memory.MemoryLocation.Slot),
				fmt.Sprintf("%v", memory.MemoryLocation.Socket),
				fmt.Sprintf("%v", memory.MemoryType),
				fmt.Sprintf("%v", memory.Name),
				fmt.Sprintf("%v", memory.Status.Health),
				fmt.Sprintf("%v", memory.Status.State),
			)
		}
	}

}
