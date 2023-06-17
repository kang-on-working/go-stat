package pkg

import (
	// "fmt"
	"strings"
)


type CPUInfo struct {
	CPUProcessorNo		string	`json:"cpu_proc_no"`
	CPUFamily			string 	`json:"cpu_family"`
	CPUArchitecture		string	`json:"cpu_arch"`
	CPUModel			string	`json:"cpu_model"`
	CPUModelName		string	`json:"cpu_model_name"`
	CPUMHz				string	`json:"cpu_freq"`
	CPUCacheSize		string	`json:"cpu_cache_size"`
	CPUPhysicalID		string	`json:"cpu_physic_id"`
	CPUCoreID			string	`json:"cpu_core_id"`
	CPUCores			string	`json:"cpu_co"`
}


func ParseCPUInfo(cpuInfoStr string) ([]CPUInfo, error) {
	cpuInfoList := []CPUInfo{}
	cpuInfo := CPUInfo{}

	lines := strings.Split(cpuInfoStr, "\n")
	for _, line := range lines {
		if line == "" {
			// 빈 줄을 만나면 현재 CPU 정보를 리스트에 추가하고 새로운 CPU 정보 생성
			cpuInfoList = append(cpuInfoList, cpuInfo)
			cpuInfo = CPUInfo{}
			continue
		}
		fields := strings.Split(line, ":")
		if len(fields) < 2 { continue }
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])

		switch key {
		case "processor":
			cpuInfo.CPUProcessorNo = value
		case "cpu family":
			cpuInfo.CPUFamily = value
		case "model":
			cpuInfo.CPUModel = value
		case "model name":
			cpuInfo.CPUModelName = value
		case "cpu MHz":
			cpuInfo.CPUMHz = value
		case "cache size":
			cpuInfo.CPUCacheSize = value
		case "physical id":
			cpuInfo.CPUPhysicalID = value
		case "core id":
			cpuInfo.CPUCoreID = value
		case "cpu cores":
			cpuInfo.CPUCores = value
		}

		cpuInfo.CPUArchitecture = GetCPUArch(cpuInfo.CPUFamily)
	}

	if len(cpuInfo.CPUProcessorNo) > 0 {
		// 마지막 CPU 정보 추가
		cpuInfoList = append(cpuInfoList, cpuInfo)
	}

	return cpuInfoList, nil
}

func GetCPUArch(cpu_family string) string {
	// mapping cpu family into ISA Architecture (CPU Architecture)
	switch cpu_family {
	case "1":
		return "8086"
	case "2":
		return "80286"
	case "3":
		return "80386"
	case "4":
		return "80486"
	case "5":
		return "Pentium"
	case "6":
		return "x86"
	case "15":
		return "x86-64 (AMD64)"
	case "42":
		return "ARM"
	case "62":
		return "PowerPC"
	case "86":
		return "IA-64 (Itanium)"
	case "180":
		return "ARM64"
	case "220":
		return "RISC-V"
	// 추가적인 패밀리에 대한 매핑을 수행합니다.
	default:
		return "Unknown"
	}
}

type MEMInfo struct {
	MemTotal      string `json:"mem_total"`
	MemFree       string `json:"mem_free"`
	MemAvailable  string `json:"mem_available`
	MemBuffer     string `json:"mem_buf"`
	MemCached     string `json:"mem_cached"`
	MemSwapCached string `json:"mem_swap_cached"`
	MemActive     string `json:"mem_active"`
	MemInactive   string `json:"mem_inactive"`
	MemSwapTotal  string `json:"mem_swap_total"`
	MemSwapFree   string `json:"mem_swap_free"`
}

func ParseMEMInfo(memInfoStr string) (MEMInfo, error) {

	memInfo := MEMInfo{}

	lines := strings.Split(memInfoStr, "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) <2 { continue }

		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		
		switch key {
		case "MemTotal":
			memInfo.MemTotal = value
		case "MemFree":
			memInfo.MemFree = value
		case "MemAvailable":
			memInfo.MemAvailable = value
		case "Buffers":
			memInfo.MemBuffer = value
		case "Cached":
			memInfo.MemCached = value
		case "SwapCached":
			memInfo.MemSwapCached = value
		case "Active":
			memInfo.MemActive = value
		case "Inactive":
			memInfo.MemInactive = value
		case "SwapTotal":
			memInfo.MemSwapTotal = value
		case "SwapFree":
			memInfo.MemSwapFree = value

		}
	}

	return memInfo, nil
}
