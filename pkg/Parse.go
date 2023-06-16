package pkg


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
