package test

import (
	"exexm.com/golang_common/lib"
	"fmt"
	"testing"
)
/**
  官方文档：https://github.com/olivere/elastic
 */
func Test_ES(t *testing.T) {
	SetUp()
	version,err :=lib.ESClient.ElasticsearchVersion("http://172.17.221.29:9200")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(version)
}
