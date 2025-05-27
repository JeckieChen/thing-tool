package test

import (
	"fmt"
	"testing"
	esdefine "thing-tool/modules/elasticsearch/define"
	"thing-tool/modules/elasticsearch/service"

	"github.com/go-resty/resty/v2"
)

func TestClient(t *testing.T) {
	es := service.ESService{
		Client:     resty.New(),
		ConnectObj: &esdefine.Connect{Host: "http://172.16.32.207:9200", Username: "elastic", Password: "zg123456"},
	}
	d, _ := es.TestClient("http://172.16.32.207:9200", "elastic", "zg123456")

	g := es.GetNodes()
	fmt.Println(d)
	fmt.Println(g.Results[0])
}
