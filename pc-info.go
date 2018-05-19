package main

import (
	"fmt"
	"os/exec"
	"strings"
)

//	отработка uname
func commandUname(currentPc *PcInfoStruct) {
	// заполняем
	//	network node hostname
	targetB, err := exec.Command("uname", "-n").Output()
	if err != nil {
		currentPc.Hostname = "unknown"
	} else {
		currentPc.Hostname = string(targetB)
	}
	//	kernel name
	targetB, err = exec.Command("uname", "-s").Output()
	if err != nil {
		fmt.Println("Error")
		currentPc.KernelName = "unknown"
	} else {
		currentPc.KernelName = string(targetB)
	}
	//	kernel release
	targetB, err = exec.Command("uname", "-r").Output()
	if err != nil {
		fmt.Println("Error")
		currentPc.KernelRelease = "unknown"
	} else {
		currentPc.KernelRelease = string(targetB)
	}
	//	kernel version
	targetB, err = exec.Command("uname", "-v").Output()
	if err != nil {
		fmt.Println("Error")
		currentPc.KernelVersion = "unknown"
	} else {
		currentPc.KernelVersion = string(targetB)
	}
	//	kernel name
	targetB, err = exec.Command("uname", "-s").Output()
	if err != nil {
		fmt.Println("Error")
		currentPc.Hostname = "unknown"
	} else {
		currentPc.Hostname = string(targetB)
	}
	//	hardware name
	targetB, err = exec.Command("uname", "-m").Output()
	if err != nil {
		fmt.Println("Error")
		currentPc.HardwareName = "unknown"
	} else {
		currentPc.HardwareName = string(targetB)
	}
	//	processor type
	targetB, err = exec.Command("uname", "-p").Output()
	if err != nil {
		fmt.Println("Error")
		currentPc.Processor = "unknown"
	} else {
		currentPc.Processor = string(targetB)
	}
	//	hardware platform
	targetB, err = exec.Command("uname", "-i").Output()
	if err != nil {
		fmt.Println("Error")
		currentPc.HardwarePlatform = "unknown"
	} else {
		currentPc.HardwarePlatform = string(targetB)
	}
	//	os
	targetB, err = exec.Command("uname", "-o").Output()
	if err != nil {
		fmt.Println("Error")
		currentPc.Os = "unknown"
	} else {
		currentPc.Os = string(targetB)
	}

}

//	отработка Ip
func commandIp(currentPc *PcInfoStruct) {
	// заполняем
	commandIpLink(currentPc)
}

//	отработка Ip Link
func commandIpLink(currentPc *PcInfoStruct) {

	fmt.Println("отработка Ip Link")
	// заполняем
	coupleB1, err := exec.Command("bash", "-c", "ip link | awk 'NR % 2 == 1 {print $2}'").Output()
	var netNames, netState []string
	var netMAC []string

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		//currentPc.Os = "unknown"
	} else {
		netNames = strings.Split(string(coupleB1), ":\n")
		fmt.Println(netNames)
	}

	coupleB2, err := exec.Command("bash", "-c", "ip link | awk 'NR % 2 == 1 {print $9}'").Output()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		//currentPc.Os = "unknown"
	} else {
		netMAC = strings.Split(string(coupleB2), "\n")
		fmt.Println(netMAC)
	}

	coupleB3, err := exec.Command("bash", "-c", " ip link | awk 'NR % 2 == 0 {print $2}'").Output()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		//currentPc.Os = "unknown"
	} else {
		netState = strings.Split(string(coupleB3), "\n")
		fmt.Println(netState)
	}
	//	перебираем
	for i, c := range netNames {
		fmt.Println(c, i)
		var networkInterface NetworkInterfaces
		networkInterface.Name = netNames[i]
		networkInterface.MAC = netMAC[i]
		networkInterface.State = netState[i]
		currentPc.NetworkInterfaces = append(currentPc.NetworkInterfaces, networkInterface)
	}
	//var i = 0
	//for c := range coupleB2 {
	//	tempNetworkInterface[i].MAC = string(c)
	//	fmt.Println(c)
	//}
	//var i = 0
	//for c := range coupleB3 {
	//	tempNetworkInterface[i].State = string(c)
	//	fmt.Println(c)
	//}
	//fmt.Println(tempNetworkInterface)

	////	network node hostname
	//targetB, err = exec.Command("uname", "-n").Output()
	//if err != nil {
	//	currentPc.Hostname = "unknown"
	//} else {
	//	currentPc.Hostname = string(targetB)
	//}
	////	kernel name
	//targetB, err = exec.Command("uname", "-s").Output()
	//if err != nil {
	//	fmt.Println("Error")
	//	currentPc.KernelName = "unknown"
	//} else {
	//	currentPc.KernelName = string(targetB)
	//}
	////	kernel release
	//targetB, err = exec.Command("uname", "-r").Output()
	//if err != nil {
	//	fmt.Println("Error")
	//	currentPc.KernelRelease = "unknown"
	//} else {
	//	currentPc.KernelRelease = string(targetB)
	//}
	////	kernel version
	//targetB, err = exec.Command("uname", "-v").Output()
	//if err != nil {
	//	fmt.Println("Error")
	//	currentPc.KernelVersion = "unknown"
	//} else {
	//	currentPc.KernelVersion = string(targetB)
	//}
	////	kernel name
	//targetB, err = exec.Command("uname", "-s").Output()
	//if err != nil {
	//	fmt.Println("Error")
	//	currentPc.Hostname = "unknown"
	//} else {
	//	currentPc.Hostname = string(targetB)
	//}
	////	hardware name
	//targetB, err = exec.Command("uname", "-m").Output()
	//if err != nil {
	//	fmt.Println("Error")
	//	currentPc.HardwareName = "unknown"
	//} else {
	//	currentPc.HardwareName = string(targetB)
	//}
	////	processor type
	//targetB, err = exec.Command("uname", "-p").Output()
	//if err != nil {
	//	fmt.Println("Error")
	//	currentPc.Processor = "unknown"
	//} else {
	//	currentPc.Processor = string(targetB)
	//}
	////	hardware platform
	//targetB, err = exec.Command("uname", "-i").Output()
	//if err != nil {
	//	fmt.Println("Error")
	//	currentPc.HardwarePlatform = "unknown"
	//} else {
	//	currentPc.HardwarePlatform = string(targetB)
	//}
	////	os
	//targetB, err = exec.Command("uname", "-o").Output()
	//if err != nil {
	//	fmt.Println("Error")
	//	currentPc.Os = "unknown"
	//} else {
	//	currentPc.Os = string(targetB)
	//}

}
