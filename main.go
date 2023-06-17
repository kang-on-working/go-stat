package main

import (
	"fmt"
	"pkg"
)

func main() {
	target := "./info.json"

	cpu, err := pkg.GetCPUInfo()
	if err != nil { return }
	cpuinfo, err := pkg.ParseCPUInfo(cpu)
	if err != nil { return }
	cpujsonData, err := pkg.CPUInfotoJson(cpuinfo)
	if err != nil { return }
	fmt.Println(cpujsonData)

	mem, err := pkg.GetMEMInfo()
	if err != nil { return }
	meminfo, err := pkg.ParseMEMInfo(mem)
	if err != nil { return }
	memjsonData, err := pkg.MEMInfotoJson(meminfo)
	if err != nil { return }
	fmt.Println(memjsonData)

	jsonData := cpujsonData + memjsonData

	pkg.WriteFile(target, []byte(jsonData))
}
