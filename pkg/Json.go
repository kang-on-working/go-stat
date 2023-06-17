package pkg

import ( 
	"encoding/json"
	"fmt"
)

func CPUInfotoJson(cpuInfoList interface{}) (string, error) {
	jsonData, err := json.Marshal(cpuInfoList)
	if err != nil { fmt.Println("Error", err) ; return "", err }
	return string(jsonData), nil
}


func MEMInfotoJson(memInfo MEMInfo) (string, error) {
	jsonData, err := json.Marshal(memInfo)
	if err != nil { fmt.Println("Error", err) ; return "", err }
	return string(jsonData), nil
}

