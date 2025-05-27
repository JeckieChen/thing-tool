package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thing-tool/modules/elasticsearch/define"

	"github.com/go-resty/resty/v2"
)

const (
	HealthApi = "/_cluster/health"
	NodesApi  = "/_cat/nodes?format=json&pretty&h=ip,name,heap.percent,heap.current,heap.max,ram.percent,ram.current,ram.max,node.role,master,cpu,load_5m,disk.used_percent,disk.used,disk.total,fielddataMemory,queryCacheMemory,requestCacheMemory,segmentsMemory,segments.count"
)

type ESService struct {
	ConnectObj *define.Connect
	Client     *resty.Client
}

var ClientList = make(map[int]*ESService)

func NewESService(id int) *ESService {
	return ClientList[id]
}

func SetConnection(connectObj *define.Connect) error {
	v := ClientList[connectObj.Id]
	if v != nil {
		delete(ClientList, connectObj.Id)
	}
	ClientList[connectObj.Id] = &ESService{
		Client:     resty.New(),
		ConnectObj: connectObj,
	}
	return nil
}

// 连接测试
func (es *ESService) TestClient(host, username, password string) (string, error) {
	client := es.Client
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

func (es *ESService) GetNodes() *define.ResultsResp {
	if es.ConnectObj.Host == "" {
		return &define.ResultsResp{Err: "请先选择一个集群"}
	}
	var result []any
	resp, err := es.Client.R().SetResult(&result).Get(es.ConnectObj.Host + NodesApi)
	fmt.Println("nodes====================")
	fmt.Println(resp)
	fmt.Println("nodes====================")
	if err != nil {
		return &define.ResultsResp{Err: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return &define.ResultsResp{Err: string(resp.Body())}
	}
	var nodes []define.Node
	json.Unmarshal(resp.Body(), &nodes)
	return &define.ResultsResp{Results: nodes}
}
