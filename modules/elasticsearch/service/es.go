package service

import (
	"net/http"
	"thing-tool/modules/elasticsearch/define"

	"github.com/go-resty/resty/v2"
)

const (
	HealthApi = "/_cluster/health"
)

type ESService struct {
	ConnectObj *define.Connect
	Client     *resty.Client
}

func (es *ESService) TestClient(host, username, password string) (string, error) {
	client := resty.New()
	if username != "" && password != "" {
		client.SetBasicAuth(username, password)
	}
	resp, err := client.R().Get(host + HealthApi)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() == http.StatusOK {
		return string(resp.Body()), nil
	}
	return "", nil
}
