package collector

import (
	"fmt"
	"hpilo_exporter/config"
	"math"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
	//	"github.com/stmcginnis/gofish/redfish"
)

type SystemCollector struct{}

func (collector SystemCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.S_health
	ch <- config.S_memory
	ch <- config.S_processor
	ch <- config.S_network_adapter_status
	ch <- config.S_ethernetinterface
	ch <- config.S_networkport
	ch <- config.S_network_interfaces_status
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
		//
		sys_collector.collectEthernetInterfaces(ch, system)
		sys_collector.collectorNetworks(ch, system)
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

func (collector SystemCollector) collectEthernetInterfaces(ch chan<- prometheus.Metric, system *redfish.ComputerSystem) {
	ethernetInterfaces, ethernetErr := system.EthernetInterfaces()
	if nil != ethernetErr {
		panic(ethernetErr)
	}

	if 0 != len(ethernetInterfaces) {
		for _, ethernetInterface := range ethernetInterfaces {
			status := config.State_dict[string(ethernetInterface.Status.Health)]
			ch <- prometheus.MustNewConstMetric(config.S_ethernetinterface,
				prometheus.GaugeValue,
				float64(status),
				fmt.Sprintf("%v", ethernetInterface.AutoNeg),
				ethernetInterface.Description,
				fmt.Sprintf("%v", ethernetInterface.EthernetInterfaceType),
				ethernetInterface.FQDN,
				fmt.Sprintf("%v", ethernetInterface.FullDuplex),
				ethernetInterface.HostName,
				ethernetInterface.MACAddress,
				fmt.Sprintf("%v", ethernetInterface.MTUSize),
				fmt.Sprintf("%v", ethernetInterface.SpeedMbps),
			)
		}
	}
}

func (collector SystemCollector) collectProcessors(ch chan<- prometheus.Metric, system *redfish.ComputerSystem) {
	processors, proErr := system.Processors()

	if nil != proErr {
		panic(proErr)
	}

	for _, processor := range processors {
		status := config.State_dict[string(processor.Status.Health)]
		ch <- prometheus.MustNewConstMetric(config.S_processor,
			prometheus.GaugeValue,
			float64(status),
			processor.Actions,
			processor.Description,
			processor.Manufacturer,
			fmt.Sprintf("%v", processor.MaxSpeedMHz),
			fmt.Sprintf("%v", processor.MaxTDPWatts),
			processor.Model,
			fmt.Sprintf("%v", processor.ProcessorType),
			processor.Socket,
			processor.SubProcessors,
			fmt.Sprintf("%v", processor.TDPWatts),
			fmt.Sprintf("%v", processor.TotalCores),
			fmt.Sprintf("%v", processor.TotalEnabledCores),
			fmt.Sprintf("%v", processor.TotalThreads),
			processor.UUID,
		)
	}
}

func (collector SystemCollector) collectorNetworks(ch chan<- prometheus.Metric, system *redfish.ComputerSystem) {
	interfaces, err := system.NetworkInterfaces()

	if nil != err {
		panic(err)
	}

	if 0 != len(interfaces) {
		collector.makeNetworkPortMetricFromNetworkInterfaces(ch, interfaces)
	}
}

func (collector SystemCollector) makeNetworkPortMetricFromNetworkInterfaces(ch chan<- prometheus.Metric,
	interfaces []*redfish.NetworkInterface) {
	for _, netInterface := range interfaces {
		adapter, err := netInterface.NetworkAdapter()

		if nil != err {
			panic(err)
		}

		if nil != adapter {
			collector.collectNetworkPortMetricFromNetworkAdapter(ch, adapter)
			collector.collectNetworkAdapterStatus(ch, adapter)
		}
	}
}

func (collector SystemCollector) collectNetworkPortMetricFromNetworkAdapter(ch chan<- prometheus.Metric,
	adapter *redfish.NetworkAdapter) {
	networkPorts, err := adapter.NetworkPorts()
	netState := map[string]float64{"Up": 0.0, "Down": 1.0}

	if nil != err {
		panic(err)
	}

	for _, networkPort := range networkPorts {
		stateString := fmt.Sprintf("%v", networkPort.LinkStatus)
		status := netState[stateString]
		ch <- prometheus.MustNewConstMetric(config.S_networkport,
			prometheus.GaugeValue,
			status,
			adapter.Manufacturer,
			fmt.Sprintf("%v", networkPort.LinkStatus),
			fmt.Sprintf("%v", networkPort.CurrentLinkSpeedMbps),
			networkPort.Description,
			fmt.Sprintf("%v", networkPort.MaxFrameSize),
			fmt.Sprintf("%v", networkPort.NumberDiscoveredRemotePorts),
			networkPort.PhysicalPortNumber,
			fmt.Sprintf("%v", networkPort.PortMaximumMTU),
		)
	}
}

func (collector SystemCollector) collectNetworkAdapterStatus(ch chan<- prometheus.Metric,
	adapter *redfish.NetworkAdapter) {
	controllers := adapter.Controllers

	if 0 != len(controllers) {
		for _, control := range controllers {
			ch <- prometheus.MustNewConstMetric(config.S_network_adapter_status,
				prometheus.GaugeValue,
				float64(0),
				adapter.Manufacturer,
				control.FirmwarePackageVersion,
				fmt.Sprintf("%v", control.NetworkDeviceFunctionsCount),
				fmt.Sprintf("%v", control.NetworkPortsCount),
			)
		}
	}
}

//storage
func (collector SystemCollector) collectStorage(ch chan<- prometheus.Metric, system *redfish.ComputerSystem) {
	storages, storageErr := system.Storage()

	if nil != storageErr {
		panic(storageErr)
	}

	if 0 != len(storages) {
		for _, storage := range storages {
			status := config.State_dict[string(storage.Status.Health)]
			ch <- prometheus.MustNewConstMetric(config.S_storage,
				prometheus.GaugeValue,
				float64(status),
				storage.Description,
				fmt.Sprintf("%v", storage.DrivesCount),
				fmt.Sprintf("%v", storage.RedundancyCount),
				fmt.Sprintf("%v", storage.EnclosuresCount),
			)

			collector.collectDrives(ch, storage)
		}
	}
}

func (collector SystemCollector) associatedDriveIds(volume *redfish.Volume) []string {
	drives, _ := volume.Drives()
	driveId := make([]string, 0)

	if 0 != len(drives) {
		for _, drive := range drives {
			words := strings.Split(drive.Description, " ")
			driveId = append(driveId, words[len(words)-1])
		}
	}

	return driveId
}

func (collector SystemCollector) collectDrives(ch chan<- prometheus.Metric, storage *redfish.Storage) {
	drives, driveErr := storage.Drives()

	if nil != driveErr {
		panic(driveErr)
	}

	for _, drive := range drives {
		status := config.State_dict[string(drive.Status.Health)]
		ch <- prometheus.MustNewConstMetric(config.S_storage_drive,
			prometheus.GaugeValue,
			float64(status),
			fmt.Sprintf("%v", drive.BlockSizeBytes),
			fmt.Sprintf("%v", drive.CapableSpeedGbs),
			collector.convertCapacity(float64(drive.CapacityBytes)),
			drive.Description,
			fmt.Sprintf("%v", drive.IndicatorLED),
			drive.Manufacturer,
			fmt.Sprintf("%v", drive.MediaType),
			drive.Model,
			drive.PartNumber,
			fmt.Sprintf("%v", drive.Protocol),
			drive.Revision,
			drive.SerialNumber,
		)

		if "SSD" == fmt.Sprintf("%v", drive.MediaType) {
			collector.collectSSDDrives(ch, drive)
		}
	}
}

func (collector SystemCollector) collectSSDDrives(ch chan<- prometheus.Metric, drive *redfish.Drive) {
	ch <- prometheus.MustNewConstMetric(config.S_storage_drive_predicted_media_life_left_percent,
		prometheus.GaugeValue,
		float64(drive.PredictedMediaLifeLeftPercent),
		fmt.Sprintf("%v", drive.BlockSizeBytes),
		fmt.Sprintf("%v", drive.CapableSpeedGbs),
		collector.convertCapacity(float64(drive.CapacityBytes)),
		drive.Description,
		drive.Manufacturer,
		fmt.Sprintf("%v", drive.MediaType),
		drive.Model,
		drive.PartNumber,
		fmt.Sprintf("%v", drive.Protocol),
		drive.Revision,
		drive.SerialNumber,
	)
}

func (collector SystemCollector) convertCapacity(num float64) string {
	units := []string{"TB", "GB", "MB", "KB", "B"}
	idx := len(units) - 1

	for idx > -1 && num >= 1000 {
		idx -= 1
		num = num / 1000
	}

	return fmt.Sprintf("%v", math.RoundToEven(num)) + units[idx]
}
