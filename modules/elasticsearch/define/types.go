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

type ResultsResp struct {
	Results []any  `json:"results"`
	Err     string `json:"err"`
}

type ResultResp struct {
	Result map[string]any `json:"result"`
	Err    string         `json:"err"`
}
