package main

import (
	"fmt"
	"pkg"
)

func main() {
	cpu, err := pkg.GetCPUInfo()
	if err != nil { return }
	fmt.Println(string(cpu))
	mem, err := pkg.GetMEMInfo()
	if err != nil { return }
	fmt.Println(string(mem))
}
