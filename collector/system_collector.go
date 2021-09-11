package collector

import (
	"fmt"
	"hpilo_exporter/config"
	redfishstruct "hpilo_exporter/redfish_struct"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
	//	"github.com/stmcginnis/gofish/redfish"
)

type SystemCollector struct{}

func (collector SystemCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.S_health
	ch <- config.S_memory
	ch <- config.S_processor
	//
	ch <- config.S_ethernetinterface
	ch <- config.S_base_network_adapter_status
	ch <- config.S_storage_physical_drive_status
	ch <- config.S_storage_array_controller_status
	ch <- config.S_storage_logical_drive_status
	ch <- config.S_storage_enclosures_status

	ch <- config.S_base_network_port_adapter_status
	ch <- config.S_ilo_status
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
		//sys_collector.collectorNetworks(ch, system)

	}
	sys_collector.collectPhysicalDriveStatus(ch, &redfishstruct.PhysicalDrives{})
	sys_collector.collectArrayControllerStatus(ch, &redfishstruct.ArrayControllers{})
	sys_collector.collectLogicalDriveStatus(ch, &redfishstruct.LogicalDrives{})
	sys_collector.collectEnclosureStatus(ch, &redfishstruct.StorageEnclosures{})
	sys_collector.collectBaseNetworkAdapterStatus(ch, &redfishstruct.BaseNetworkAdapter{})
	//sys_collector.collectPortNetworkAdapterStatus(ch, &redfishstruct.BaseNetworkAdapter{})
	sys_collector.collectILOPortStatus(ch, &redfishstruct.ILOPort{})
}

func (collector SystemCollector) collectSystemHealth(ch chan<- prometheus.Metric, v *redfish.ComputerSystem) {
	fmt.Println("system connected")

	status := config.State_dict[string(v.Status.Health)]
	ch <- prometheus.MustNewConstMetric(
		config.S_health,
		prometheus.GaugeValue,
		float64(status),
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
			//status := config.State_dict[string(memory.Status.Health)]
			ch <- prometheus.MustNewConstMetric(
				config.S_memory,
				prometheus.GaugeValue,
				float64(memory_temp1),
				fmt.Sprintf("%v", memory.ID),
				fmt.Sprintf("%v", memory.BusWidthBits),
				fmt.Sprintf("%v", memory.CacheSizeMiB),
				fmt.Sprintf("%v", memory.CapacityMiB),
				fmt.Sprintf("%v", memory.DataWidthBits),
				fmt.Sprintf("%v", memory.DeviceLocator),
				fmt.Sprintf("%v", memory.LogicalSizeMiB),
				fmt.Sprintf("%v", memory.MemoryDeviceType), //
				fmt.Sprintf("%v", memory.MemoryLocation.Channel),
				fmt.Sprintf("%v", memory.MemoryLocation.MemoryController),
				fmt.Sprintf("%v", memory.MemoryLocation.Slot),
				fmt.Sprintf("%v", memory.MemoryLocation.Socket),
				fmt.Sprintf("%v", memory.MemoryType),
				fmt.Sprintf("%v", memory.Name),
				fmt.Sprintf("%v", memory.PartNumber), //
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
			//status := config.State_dict[string(processor.Status.Health)]

			ch <- prometheus.MustNewConstMetric(
				config.S_processor,
				prometheus.GaugeValue,
				float64(processor_temp1),
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
			ethernet_temp := string(ethernetInterface.Status.Health)
			ethernet_temp1 := 0.0
			if ethernet_temp == "OK" {
				ethernet_temp1 = 0
			} else if ethernet_temp == "WARNING" {
				ethernet_temp1 = 1
			} else {
				ethernet_temp1 = 2
			}
			//status := config.State_dict[string(ethernetInterface.Status.Health)]
			ch <- prometheus.MustNewConstMetric(config.S_ethernetinterface,
				prometheus.GaugeValue,
				float64(ethernet_temp1),
				fmt.Sprintf("%v", ethernetInterface.ID), //
				fmt.Sprintf("%v", ethernetInterface.FullDuplex),
				fmt.Sprintf("%v", ethernetInterface.IPv4Addresses),
				fmt.Sprintf("%v", ethernetInterface.IPv4StaticAddresses),
				fmt.Sprintf("%v", ethernetInterface.IPv6Addresses),
				fmt.Sprintf("%v", ethernetInterface.IPv6StaticAddresses),
				fmt.Sprintf("%v", ethernetInterface.MACAddress),
				fmt.Sprintf("%v", ethernetInterface.LinkStatus),
				fmt.Sprintf("%v", ethernetInterface.SpeedMbps),
				fmt.Sprintf("%v", ethernetInterface.Status.State),
				fmt.Sprintf("%v", ethernetInterface.Status.Health),
			)
		}
	}
}

// func (collector SystemCollector) collectorNetworks(ch chan<- prometheus.Metric, system *redfish.ComputerSystem) {
// 	interfaces, err := system.NetworkInterfaces()

// 	if nil != err {
// 		panic(err)
// 	}

// 	if 0 != len(interfaces) {
// 		collector.makeNetworkPortMetricFromNetworkInterfaces(ch, interfaces)
// 	}
// }

// func (collector SystemCollector) makeNetworkPortMetricFromNetworkInterfaces(ch chan<- prometheus.Metric,
// 	interfaces []*redfish.NetworkInterface) {
// 	for _, netInterface := range interfaces {
// 		adapter, err := netInterface.NetworkAdapter()

// 		if nil != err {
// 			panic(err)
// 		}

// 		if nil != adapter {
// 			collector.collectNetworkPortMetricFromNetworkAdapter(ch, adapter)
// 			collector.collectNetworkAdapterStatus(ch, adapter)
// 		}
// 	}
// }

// func (collector SystemCollector) collectNetworkPortMetricFromNetworkAdapter(ch chan<- prometheus.Metric,
// 	adapter *redfish.NetworkAdapter) {
// 	networkPorts, err := adapter.NetworkPorts()
// 	netState := map[string]float64{"Up": 0.0, "Down": 1.0}

// 	if nil != err {
// 		panic(err)
// 	}

// 	for _, networkPort := range networkPorts {
// 		stateString := fmt.Sprintf("%v", networkPort.LinkStatus)
// 		status := netState[stateString]
// 		ch <- prometheus.MustNewConstMetric(config.S_networkport,
// 			prometheus.GaugeValue,
// 			status,
// 			adapter.Manufacturer,
// 			fmt.Sprintf("%v", networkPort.LinkStatus),
// 			fmt.Sprintf("%v", networkPort.CurrentLinkSpeedMbps),
// 			networkPort.Description,
// 			fmt.Sprintf("%v", networkPort.MaxFrameSize),
// 			fmt.Sprintf("%v", networkPort.NumberDiscoveredRemotePorts),
// 			networkPort.PhysicalPortNumber,
// 			fmt.Sprintf("%v", networkPort.PortMaximumMTU),
// 		)
// 	}
// }

// func (collector SystemCollector) collectNetworkAdapterStatus(ch chan<- prometheus.Metric,
// 	adapter *redfish.NetworkAdapter) {
// 	controllers := adapter.Controllers

// 	if 0 != len(controllers) {
// 		for _, control := range controllers {
// 			ch <- prometheus.MustNewConstMetric(config.S_network_adapter_status,
// 				prometheus.GaugeValue,
// 				float64(0),
// 				adapter.Manufacturer,
// 				control.FirmwarePackageVersion,
// 				fmt.Sprintf("%v", control.NetworkDeviceFunctionsCount),
// 				fmt.Sprintf("%v", control.NetworkPortsCount),
// 			)
// 		}
// 	}
// }

//storage
// func (collector SystemCollector) collectAllPhysicalDrive(ch chan<- prometheus.Metric,
// 	pd *redfishstruct.AllPhysicalDrives) {
// 	err, physic := pd.UnmarshalJson("/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/DiskDrives")

// 	if err != nil {
// 		panic(err)
// 	}
// 	var physic_detail redfishstruct.PhysicalDrives
// 	for _, physicdrive := range physic.Members {
// 		fmt.Println(physicdrive.MemberOID)
// 		physic_detail.UnmarshalJson(physicdrive.MemberOID)

// 	}
// }

func (collector SystemCollector) collectPhysicalDriveStatus(ch chan<- prometheus.Metric, pd *redfishstruct.PhysicalDrives) {
	var pds redfishstruct.AllPhysicalDrives
	err, physic := pds.UnmarshalJson("/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/DiskDrives")

	if err != nil {
		panic(err)
	}

	var physic_detail redfishstruct.PhysicalDrives
	for _, physicdrive := range physic.Members {
		fmt.Println(physicdrive.MemberOID)
		_, detail := physic_detail.UnmarshalJson(physicdrive.MemberOID)

		physic_temp := string(detail.Status.Health)
		physic_temp1 := 0.0
		if physic_temp == "OK" {
			physic_temp1 = 0
		} else if physic_temp == "WARNING" {
			physic_temp1 = 1
		} else {
			physic_temp1 = 2
		}
		//status := config.State_dict[string(detail.Status.Health)]
		ch <- prometheus.MustNewConstMetric(config.S_storage_physical_drive_status,
			prometheus.GaugeValue,
			float64(physic_temp1),
			fmt.Sprintf("%v", detail.Id),
			fmt.Sprintf("%v", detail.BlockSizeBytes),
			fmt.Sprintf("%v", detail.CapacityGB),
			fmt.Sprintf("%v", detail.CapacityLogicalBlocks),
			fmt.Sprintf("%v", detail.CapacityMiB),
			fmt.Sprintf("%v", detail.CarrierAuthenticationStatus),
			fmt.Sprintf("%v", detail.CurrentTemperatureCelsius),
			fmt.Sprintf("%v", detail.Description),
			fmt.Sprintf("%v", detail.DiskDriveUse),
			fmt.Sprintf("%v", detail.InterfaceSpeedMbps),
			fmt.Sprintf("%v", detail.InterfaceType),
			fmt.Sprintf("%v", detail.MaximumTemperatureCelsius),
			fmt.Sprintf("%v", detail.MediaType),
			fmt.Sprintf("%v", detail.Model),
			fmt.Sprintf("%v", detail.Name),
			fmt.Sprintf("%v", detail.PowerOnHours),
			fmt.Sprintf("%v", detail.SerialNumber),
			fmt.Sprintf("%v", detail.Status.Health),
			fmt.Sprintf("%v", detail.Status.State),
		)

	}

}

func (collector SystemCollector) collectArrayControllerStatus(ch chan<- prometheus.Metric, pd *redfishstruct.ArrayControllers) {
	var pds redfishstruct.AllArrayController
	err, physic := pds.UnmarshalJson("/redfish/v1/Systems/1/SmartStorage/ArrayControllers")

	if err != nil {
		panic(err)
	}

	var physic_detail redfishstruct.ArrayControllers
	for _, physicdrive := range physic.Members {
		fmt.Println(physicdrive.MemberOID)
		_, detail := physic_detail.UnmarshalJson(physicdrive.MemberOID)

		arr_temp := string(detail.Status.Health)
		arr_temp1 := 0.0
		if arr_temp == "OK" {
			arr_temp1 = 0
		} else if arr_temp == "WARNING" {
			arr_temp1 = 1
		} else {
			arr_temp1 = 2
		}
		//status := config.State_dict[string(detail.Status.Health)]
		ch <- prometheus.MustNewConstMetric(config.S_storage_array_controller_status,
			prometheus.GaugeValue,
			float64(arr_temp1),
			fmt.Sprintf("%v", detail.Id),
			fmt.Sprintf("%v", detail.AdapterType),
			fmt.Sprintf("%v", detail.ControllerBoard.Status.Health),
			fmt.Sprintf("%v", detail.ControllerPartNumber),
			fmt.Sprintf("%v", detail.Description),
			fmt.Sprintf("%v", detail.FirmwareVersion.Current.VersionString),
			fmt.Sprintf("%v", detail.ExternalPortCount),
			fmt.Sprintf("%v", detail.HardwareRevision),
			fmt.Sprintf("%v", detail.InternalPortCount),
			fmt.Sprintf("%v", detail.Location),
			fmt.Sprintf("%v", detail.LocationFormat),
			fmt.Sprintf("%v", detail.Model),
			fmt.Sprintf("%v", detail.Name),
			fmt.Sprintf("%v", detail.ReadCachePercent),
			fmt.Sprintf("%v", detail.SerialNumber),
			fmt.Sprintf("%v", detail.Status.Health),
			fmt.Sprintf("%v", detail.Status.State),
		)
	}
}

func (collector SystemCollector) collectLogicalDriveStatus(ch chan<- prometheus.Metric, pd *redfishstruct.LogicalDrives) {
	var pds redfishstruct.AllLogicalDrives
	err, physic := pds.UnmarshalJson("/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/LogicalDrives")

	if err != nil {
		panic(err)
	}

	var physic_detail redfishstruct.LogicalDrives
	for _, physicdrive := range physic.Members {
		fmt.Println(physicdrive.MemberOID)
		_, detail := physic_detail.UnmarshalJson(physicdrive.MemberOID)

		logical_temp := string(detail.Status.Health)
		logical_temp1 := 0.0
		if logical_temp == "OK" {
			logical_temp1 = 0
		} else if logical_temp == "WARNING" {
			logical_temp1 = 1
		} else {
			logical_temp1 = 2
		}
		//status := config.State_dict[string(detail.Status.Health)]
		ch <- prometheus.MustNewConstMetric(config.S_storage_logical_drive_status,
			prometheus.GaugeValue,
			float64(logical_temp1),
			fmt.Sprintf("%v", detail.Id),
			fmt.Sprintf("%v", detail.AccelerationMethod),
			fmt.Sprintf("%v", detail.CapacityMiB),
			fmt.Sprintf("%v", detail.Description),
			fmt.Sprintf("%v", detail.InterfaceType),
			fmt.Sprintf("%v", detail.LegacyBootPriority),
			fmt.Sprintf("%v", detail.LogicalDriveEncryption),
			fmt.Sprintf("%v", detail.LogicalDriveName),
			fmt.Sprintf("%v", detail.LogicalDriveNumber),
			fmt.Sprintf("%v", detail.LogicalDriveStatusReasons),
			fmt.Sprintf("%v", detail.LogicalDriveType),
			fmt.Sprintf("%v", detail.MediaType),
			fmt.Sprintf("%v", detail.Name),
			fmt.Sprintf("%v", detail.Raid),
			fmt.Sprintf("%v", detail.StripeSizeBytes),
			fmt.Sprintf("%v", detail.VolumeUniqueIdentifier),
			fmt.Sprintf("%v", detail.Status.Health),
			fmt.Sprintf("%v", detail.Status.State),
		)
	}
}

func (collector SystemCollector) collectEnclosureStatus(ch chan<- prometheus.Metric, pd *redfishstruct.StorageEnclosures) {
	var pds redfishstruct.AllStorageEnclosures
	err, physic := pds.UnmarshalJson("/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/StorageEnclosures")

	if err != nil {
		panic(err)
	}

	var physic_detail redfishstruct.StorageEnclosures
	for _, physicdrive := range physic.Members {
		fmt.Println(physicdrive.MemberOID)
		_, detail := physic_detail.UnmarshalJson(physicdrive.MemberOID)

		enclosure_temp := string(detail.Status.Health)
		enclosure_temp1 := 0.0
		if enclosure_temp == "OK" {
			enclosure_temp1 = 0
		} else if enclosure_temp == "WARNING" {
			enclosure_temp1 = 1
		} else {
			enclosure_temp1 = 2
		}
		//status := config.State_dict[string(detail.Status.Health)]
		ch <- prometheus.MustNewConstMetric(config.S_storage_enclosures_status,
			prometheus.GaugeValue,
			float64(enclosure_temp1),
			fmt.Sprintf("%v", detail.Id),
			fmt.Sprintf("%v", detail.Description),
			fmt.Sprintf("%v", detail.DriveBayCount),
			fmt.Sprintf("%v", detail.FirmwareVersion.Current.VersionString),
			fmt.Sprintf("%v", detail.Location),
			fmt.Sprintf("%v", detail.LocationFormat),
			fmt.Sprintf("%v", detail.Model),
			fmt.Sprintf("%v", detail.Name),
			fmt.Sprintf("%v", detail.SerialNumber),
			fmt.Sprintf("%v", detail.Status.Health),
			fmt.Sprintf("%v", detail.Status.State),
		)
	}
}

func (collector SystemCollector) collectBaseNetworkAdapterStatus(ch chan<- prometheus.Metric, pd *redfishstruct.BaseNetworkAdapter) {
	var pds redfishstruct.AllBaseNetworkAdapter
	err, physic := pds.UnmarshalJson("/redfish/v1/Systems/1/BaseNetworkAdapters")

	if err != nil {
		panic(err)
	}

	var physic_detail redfishstruct.BaseNetworkAdapter
	for _, physicdrive := range physic.Members {
		fmt.Println(physicdrive.MemberOID)
		_, detail := physic_detail.UnmarshalJson(physicdrive.MemberOID)

		nw_temp := string(detail.Status.Health)
		nw_temp1 := 0.0
		if nw_temp == "OK" {
			nw_temp1 = 0
		} else if nw_temp == "WARNING" {
			nw_temp1 = 1
		} else {
			nw_temp1 = 2
		}
		//status := config.State_dict[string(detail.Status.Health)]
		for i, v := range detail.PhysicalPorts {
			ch <- prometheus.MustNewConstMetric(config.S_base_network_adapter_status,
				prometheus.GaugeValue,
				float64(nw_temp1),
				fmt.Sprintf("%v", detail.ID),
				fmt.Sprintf("%v", detail.FirmwareVersion.Current.VersionString),
				fmt.Sprintf("%v", detail.Name),
				fmt.Sprintf("%v", detail.PartNumber),
				fmt.Sprintf("%v", v.LinkStatus),
				fmt.Sprintf("%v", v.MacAddress),
				fmt.Sprintf("%v", detail.SerialNumber),
				fmt.Sprintf("%v", detail.Status.Health),
				fmt.Sprintf("%v", detail.Status.State),
			)

			ch <- prometheus.MustNewConstMetric(config.S_base_network_port_adapter_status,
				prometheus.GaugeValue,
				float64(nw_temp1),
				fmt.Sprintf("%v", detail.ID),
				fmt.Sprintf("%v", i+1),
				fmt.Sprintf("%v", v.IPv4Addresses),
				fmt.Sprintf("%v", v.MacAddress),
				fmt.Sprintf("%v", v.LinkStatus),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadTransmits),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodTransmits),
				fmt.Sprintf("%v", v.SpeedMbps),
				fmt.Sprintf("%v", v.Status.Health),
			)

			ch <- prometheus.MustNewConstMetric(config.S_base_network_port_good_transmits,
				prometheus.GaugeValue,
				float64(v.OemPhysicalPort.HpePhysicalPort.GoodTransmits),
				fmt.Sprintf("%v", detail.ID),
				fmt.Sprintf("%v", i+1),
				fmt.Sprintf("%v", v.IPv4Addresses),
				fmt.Sprintf("%v", v.MacAddress),
				fmt.Sprintf("%v", v.LinkStatus),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadTransmits),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodTransmits),
				fmt.Sprintf("%v", v.SpeedMbps),
				fmt.Sprintf("%v", v.Status.Health),
			)
			ch <- prometheus.MustNewConstMetric(config.S_base_network_port_bad_transmits,
				prometheus.GaugeValue,
				float64(v.OemPhysicalPort.HpePhysicalPort.BadTransmits),
				fmt.Sprintf("%v", detail.ID),
				fmt.Sprintf("%v", i+1),
				fmt.Sprintf("%v", v.IPv4Addresses),
				fmt.Sprintf("%v", v.MacAddress),
				fmt.Sprintf("%v", v.LinkStatus),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadTransmits),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodTransmits),
				fmt.Sprintf("%v", v.SpeedMbps),
				fmt.Sprintf("%v", v.Status.Health),
			)
			ch <- prometheus.MustNewConstMetric(config.S_base_network_port_good_receives,
				prometheus.GaugeValue,
				float64(v.OemPhysicalPort.HpePhysicalPort.GoodReceives),
				fmt.Sprintf("%v", detail.ID),
				fmt.Sprintf("%v", i+1),
				fmt.Sprintf("%v", v.IPv4Addresses),
				fmt.Sprintf("%v", v.MacAddress),
				fmt.Sprintf("%v", v.LinkStatus),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadTransmits),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodTransmits),
				fmt.Sprintf("%v", v.SpeedMbps),
				fmt.Sprintf("%v", v.Status.Health),
			)
			ch <- prometheus.MustNewConstMetric(config.S_base_network_port_bad_receives,
				prometheus.GaugeValue,
				float64(v.OemPhysicalPort.HpePhysicalPort.BadReceives),
				fmt.Sprintf("%v", detail.ID),
				fmt.Sprintf("%v", i+1),
				fmt.Sprintf("%v", v.IPv4Addresses),
				fmt.Sprintf("%v", v.MacAddress),
				fmt.Sprintf("%v", v.LinkStatus),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadTransmits),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodReceives),
				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodTransmits),
				fmt.Sprintf("%v", v.SpeedMbps),
				fmt.Sprintf("%v", v.Status.Health),
			)
		}

	}
}

func (collector SystemCollector) collectILOPortStatus(ch chan<- prometheus.Metric, pd *redfishstruct.ILOPort) {
	var pds redfishstruct.ILOPort
	err, iloports := pds.UnmarshalJson("/redfish/v1/Managers/1/EthernetInterfaces/1/")

	if err != nil {
		panic(err)
	}
	ilo_temp := string(iloports.Status.Health)
	ilo_temp1 := 0.0
	if ilo_temp == "OK" {
		ilo_temp1 = 0
	} else if ilo_temp == "WARNING" {
		ilo_temp1 = 1
	} else {
		ilo_temp1 = 2
	}
	//status := config.State_dict[string(iloports.Status.Health)]
	ch <- prometheus.MustNewConstMetric(config.S_ilo_status,
		prometheus.GaugeValue,
		float64(ilo_temp1),
		fmt.Sprintf("%v", iloports.ID),
		fmt.Sprintf("%v", iloports.FullDuplex),
		fmt.Sprintf("%v", iloports.HostName),
		fmt.Sprintf("%v", iloports.IPv4Addresses[0].Address),
		fmt.Sprintf("%v", iloports.IPv4Addresses[0].AddressOrigin),
		fmt.Sprintf("%v", iloports.IPv4Addresses[0].Gateway),
		fmt.Sprintf("%v", iloports.IPv4Addresses[0].SubnetMask),
		fmt.Sprintf("%v", iloports.SpeedMbps),
		fmt.Sprintf("%v", iloports.Status.Health),
		fmt.Sprintf("%v", iloports.Status.State),
		fmt.Sprintf("%v", iloports.VLAN.VLANEnable),
		fmt.Sprintf("%v", iloports.VLAN.VLANId),
	)

}

// func (collector SystemCollector) collectPortNetworkAdapterStatus(ch chan<- prometheus.Metric, pd *redfishstruct.BaseNetworkAdapter) {
// 	var pds redfishstruct.AllBaseNetworkAdapter
// 	err, physic := pds.UnmarshalJson("/redfish/v1/Systems/1/BaseNetworkAdapters")

// 	if err != nil {
// 		panic(err)
// 	}

// 	var physic_detail redfishstruct.BaseNetworkAdapter
// 	for _, physicdrive := range physic.Members {
// 		fmt.Println(physicdrive.MemberOID)
// 		_, detail := physic_detail.UnmarshalJson(physicdrive.MemberOID)

// 		for i, v := range detail.PhysicalPorts {
// 			status := config.State_dict[string(v.Status.Health)]
// 			ch <- prometheus.MustNewConstMetric(config.S_base_network_port_adapter_status,
// 				prometheus.GaugeValue,
// 				float64(status),
// 				fmt.Sprintf("%v", detail.ID),
// 				fmt.Sprintf("%v", i+1),
// 				fmt.Sprintf("%v", v.IPv4Addresses),
// 				fmt.Sprintf("%v", v.MacAddress),
// 				fmt.Sprintf("%v", v.LinkStatus),
// 				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadReceives),
// 				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.BadTransmits),
// 				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodReceives),
// 				fmt.Sprintf("%v", v.OemPhysicalPort.HpePhysicalPort.GoodTransmits),
// 				fmt.Sprintf("%v", v.SpeedMbps),
// 				fmt.Sprintf("%v", v.Status.Health),
// 			)
// 		}

// 	}
// }
