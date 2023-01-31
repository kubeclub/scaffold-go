package lib

import (
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
	"net/http"
)

func InitEs(path string) (*elastic.Client, error) {
	ConfigEs := &EsConfig{}
	err := ParseConfig(path, ConfigEs)
	if err != nil {
		return nil,err
	}

	timeout := cast.ToDuration(ConfigEs.Es.Timeout)
	httpclient := &http.Client{
		Timeout: timeout,
	}

	ESClient, err := elastic.NewClient(
		elastic.SetURL(ConfigEs.Es.Proxy),
		elastic.SetBasicAuth(ConfigEs.Es.Username,ConfigEs.Es.Password),
		elastic.SetHttpClient(httpclient),
	)
	if err != nil {
		return nil,err
	}
	return ESClient,err
}
