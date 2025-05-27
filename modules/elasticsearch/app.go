package elasticsearch

import (
	"context"
	"thing-tool/define"
	esdefine "thing-tool/modules/elasticsearch/define"
	"thing-tool/modules/elasticsearch/service"
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
		{ID: 2, Name: "测试环境", Host: "http://172.16.32.207:9200", Port: "9200", Username: "elastic", Password: "zg123456"},
	}
	return define.D{
		"code": 200,
		"data": d,
	}
}

func (m *ESModule) SetConnection(id int, host, username, password string) define.R {
	service.SetConnection(&esdefine.Connect{
		Id:       id,
		Host:     host,
		Username: username,
		Password: password,
	})
	return define.D{
		"code": 200,
		"data": "设置成功",
	}
}

func (m *ESModule) TestClient(id int, host string, username string, password string) define.R {
	es := service.NewESService(id)
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

func (m *ESModule) GetNodes(id int) define.R {
	es := service.NewESService(id)
	g := es.GetNodes()
	if g.Err != "" {
		return define.D{
			"code": 400,
			"data": g.Err,
		}
	}
	return define.D{
		"code": 200,
		"data": g.Results,
	}
}

func (m *ESModule) GetIndexes(id int, name string) define.R {
	es := service.NewESService(id)
	g := es.GetIndexes(name)
	if g.Err != "" {
		return define.D{
			"code": 400,
			"data": g.Err,
		}
	}
	return define.D{
		"code": 200,
		"data": g.Results,
	}
}
