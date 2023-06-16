package main

import (
	"fmt"
	"pkg"
)

func main() {
	cpu, err := pkg.GetCPUInfo()
	if err != nil { return }
	parsed, err := pkg.ParseCPUInfo(cpu)
	if err != nil { return }
	fmt.Println(parsed)
}
