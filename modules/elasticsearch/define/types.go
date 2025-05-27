package define

type Connect struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Host          string `json:"host"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	UseSSL        bool   `json:"useSSL"`
	SkipSSLVerify bool   `json:"skipSSLVerify"`
	CACert        string `json:"caCert"`
}

type Node struct {
	Ip                 string `json:"ip"`
	Name               string `json:"name"`
	HeapPercent        string `json:"heap.percent"`
	HeapCurrent        string `json:"heap.current"`
	HeapMax            string `json:"heap.max"`
	RamPercent         string `json:"ram.percent"`
	RamCurrent         string `json:"ram.current"`
	RamMax             string `json:"ram.max"`
	NodeRole           string `json:"node.role"`
	Master             string `json:"master"`
	Cpu                string `json:"cpu"`
	Load5m             string `json:"load_5m"`
	DiskUsedPercent    string `json:"disk.used_percent"`
	DiskUsed           string `json:"disk.used"`
	DiskTotal          string `json:"disk.total"`
	FielddataMemory    string `json:"fielddataMemory"`
	QueryCacheMemory   string `json:"queryCacheMemory"`
	RequestCacheMemory string `json:"requestCacheMemory"`
	SegmentsMemory     string `json:"segmentsMemory"`
	SegmentsCount      string `json:"segments.count"`
}

type ResultsResp struct {
	Results interface{}
	Err     string
}

type ResultResp struct {
	Result map[string]any `json:"result"`
	Err    string         `json:"err"`
}
