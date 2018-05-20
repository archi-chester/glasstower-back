package main

const (
	LISTENER_PORT  = "22022"
	LISTENER_PORT2 = "22024"
)

type PcInfoStruct struct {
	Hostname         string
	KernelName       string
	KernelRelease    string
	KernelVersion    string
	HardwareName     string
	Processor        string
	HardwarePlatform string
	Os               string
	//	hardware
	NetworkInterfaces []NetworkInterfaces
}

//type NetworkInfo struct {
//	CpuName string
//	CpuName string
//	CpuArch string
//	CpuName string
//	CpuName string
//	CpuName string
//	CpuName string
//	CpuName string
//
//}

//	Настройки сетевых интерфейсов
type NetworkInterfaces struct {
	Name  string
	MAC   string
	State string
}
