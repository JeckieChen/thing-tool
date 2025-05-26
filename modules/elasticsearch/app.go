package elasticsearch

import (
	"context"
	"thing-tool/define"
	esdefine "thing-tool/modules/elasticsearch/define"
	"thing-tool/modules/elasticsearch/service"

	"github.com/go-resty/resty/v2"
)

type ESModule struct {
	ctx context.Context
}

type Connection struct {
	ID       int    `json:"id"` // 必须添加 json 标签
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewESModule() *ESModule {
	return &ESModule{}
}

func (m *ESModule) Init(ctx context.Context) {
	m.ctx = ctx
}

func (m *ESModule) GeTest() define.R {
	d := []Connection{
		{ID: 1, Name: "生产环境", Host: "10.0.0.1"},
		{ID: 2, Name: "测试环境", Host: "172.16.32.207", Port: "9200", Username: "elastic", Password: "zg123456"},
	}
	return define.D{
		"code": 200,
		"data": d,
	}
}

func (m *ESModule) TestClient(host string, username string, password string) define.R {
	es := service.ESService{
		Client:     resty.New(),
		ConnectObj: &esdefine.Connect{Host: host, Username: username, Password: password},
	}
	s, err := es.TestClient(host, username, password)
	if s == "" || err != nil {
		return define.D{
			"code": 500,
			"data": "连接失败",
		}
	}
	return define.D{
		"code": 200,
		"data": s,
	}
}
