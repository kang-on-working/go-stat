package main

import (
	"fmt"
	"pkg"
)

func main() {
	mem, err := pkg.GetMEMInfo()
	if err != nil { return }
	parsed, err := pkg.ParseMEMInfo(mem)
	if err != nil { return }
	jsonData, err := pkg.MEMInfotoJson(parsed)
	if err != nil { return }
	fmt.Println(jsonData)
}
