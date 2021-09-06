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
	ch <- config.S_processor
	ch <- config.S_storage_array_status
}

func (sys_collector SystemCollector) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, sysErr := metric.Systems() //Systems get the system instances from the service

	if nil != sysErr {
		panic(sysErr)
	}
	for _, system := range systems {
		sys_collector.collectSystemHealth(ch, system)
		sys_collector.collectMemories(ch, system)
		sys_collector.collectProcessor(ch, system)
		sys_collector.collectStorageArrayController(ch, system)
	}
}

func (collector SystemCollector) collectSystemHealth(ch chan<- prometheus.Metric, v *redfish.ComputerSystem) {
	fmt.Println("system connected")

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
	fmt.Println("memory connected")

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

func (collector SystemCollector) collectProcessor(ch chan<- prometheus.Metric, v *redfish.ComputerSystem) {

	processors, err := v.Processors()

	if nil != err {
		panic(err)
	}

	if err == nil {
		for _, processor := range processors {
			processor_temp := string(processor.Status.Health)
			processor_temp1 := 0.0
			if processor_temp == "OK" {
				processor_temp1 = 0
			} else if processor_temp == "WARNING" {
				processor_temp1 = 1
			} else {
				processor_temp1 = 2
			}

			ch <- prometheus.MustNewConstMetric(
				config.S_processor,
				prometheus.GaugeValue,
				processor_temp1,
				fmt.Sprintf("%v", processor.ID),
				fmt.Sprintf("%v", processor.InstructionSet),
				fmt.Sprintf("%v", processor.Manufacturer),
				fmt.Sprintf("%v", processor.MaxSpeedMHz),
				fmt.Sprintf("%v", processor.Model),
				fmt.Sprintf("%v", processor.ProcessorArchitecture),
				fmt.Sprintf("%v", processor.ProcessorType),
				fmt.Sprintf("%v", processor.Status.Health),
				fmt.Sprintf("%v", processor.Status.State),
				fmt.Sprintf("%v", processor.TotalCores),
				fmt.Sprintf("%v", processor.TotalEnabledCores),
				fmt.Sprintf("%v", processor.TotalThreads),
			)
		}

	}
}

// /redfish/v1/Systems/1/SmartStorage/ArrayControllers/x/
func (collector SystemCollector) collectStorageArrayController(ch chan<- prometheus.Metric, v *redfish.ComputerSystem) {
	storages, err := v.Storage()

	if nil != err {
		panic(err)
	}

	if err == nil {
		for _, storage := range storages {
			storage_temp := string(storage.Status.Health)
			storage_temp1 := 0.0
			if storage_temp == "OK" {
				storage_temp1 = 0
			} else if storage_temp == "WARNING" {
				storage_temp1 = 1
			} else {
				storage_temp1 = 2
			}
			fmt.Println(storage.ODataContext)
			ch <- prometheus.MustNewConstMetric(
				config.S_storage_array_status,
				prometheus.GaugeValue,
				storage_temp1,
				fmt.Sprintf("%v", storage.Description),
				fmt.Sprintf("%v", storage.DrivesCount),
				fmt.Sprintf("%v", storage.EnclosuresCount),
				fmt.Sprintf("%v", storage.StorageControllersCount),
				fmt.Sprintf("%v", storage.StorageControllers),
				fmt.Sprintf("%v", storage.ID),
				fmt.Sprintf("%v", storage.Name),
				fmt.Sprintf("%v", storage.ODataID),
			)

			// storage_encloses, err := storage.Enclosures()
			// if err != nil {
			// 	panic(err)
			// }

			// if err == nil {
			// 	for _, enclose := range storage_encloses {

			// 	}
			// }
		}
	}

}
