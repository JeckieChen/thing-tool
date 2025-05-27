package test

import (
	"fmt"
	"testing"
	esdefine "thing-tool/modules/elasticsearch/define"
	"thing-tool/modules/elasticsearch/service"
)

func TestClient(t *testing.T) {

	service.SetConnection(&esdefine.Connect{Id: 1, Host: "http://172.16.32.207:9200", Username: "elastic", Password: "zg123456"})
	fmt.Println(service.ClientList)
	es := service.NewESService(1)

	es.TestClient("http://172.16.32.207:9200", "elastic", "zg123456")

	g := es.GetNodes()
	es.GetIndexes("")
	fmt.Println(g)
}

func TestIndexes(t *testing.T) {
	service.SetConnection(&esdefine.Connect{Id: 1, Host: "http://172.16.32.207:9200", Username: "elastic", Password: "zg123456"})
	fmt.Println(service.ClientList)
	es := service.NewESService(1)
	es.TestClient("http://172.16.32.207:9200", "elastic", "zg123456")
	g := es.GetIndexes("")
	fmt.Println(g.Results)
}
