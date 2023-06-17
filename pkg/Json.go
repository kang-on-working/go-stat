package pkg

import ( 
	"encoding/json"
	"fmt"
	"strings"
)

func CPUInfotoJson(cpuInfoList interface{}) (string, error) {
	jsonData, err := json.Marshal(cpuInfoList)
	if err != nil { fmt.Println("Error", err) ; return "", err }
	var jsonString string = string(jsonData)
	jsonString = strings.ReplaceAll(jsonString, "[", "\"CPU Info\": {")
	jsonString = strings.ReplaceAll(jsonString, "]", "}, \n")

	return jsonString, nil
}


func MEMInfotoJson(memInfo MEMInfo) (string, error) {
	jsonData, err := json.Marshal(memInfo)
	if err != nil { fmt.Println("Error", err) ; return "", err }
	var jsonString string = string(jsonData)
	jsonString = strings.ReplaceAll(jsonString, "{", "\"MEM Info\": {")
	return jsonString, nil
}

