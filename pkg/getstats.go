package pkg

func GetCPUInfo() (string ,error) {
	const Target string = "/proc/cpuinfo"
	data, err := ReadFile(Target)
	if err != nil { return "", err }
	return string(data), err
}

func GetMEMInfo() (string ,error) {
	const Target string = "/proc/meminfo"
	data, err := ReadFile(Target)
	if err != nil { return "", err }
	return string(data), err
}
