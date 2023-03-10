package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-pluspoint/utils"
	"net/http"
)

type Kairosdb struct {
	Server      string
	Port        string
	QueryUrl    string
	DeleteUrl   string
	PushUrl     string
	DelUrl      string
	Headersjson map[string]string
	Headersgzip map[string]string
}

func NewKairosdb() Kairosdb {
	var k Kairosdb
	k.Server = utils.KairosdbHost
	k.Port = utils.KairosdbPort
	k.QueryUrl = fmt.Sprintf("http://%s:%s/api/v1/datapoints/query", k.Server, k.Port)
	k.DeleteUrl = fmt.Sprintf("http://%s:%s/api/v1/datapoints/delete", k.Server, k.Port)
	k.PushUrl = fmt.Sprintf("http://%s:%s/api/v1/datapoints", k.Server, k.Port)
	k.DelUrl = fmt.Sprintf("http://%s:%s/api/v1/metric/", k.Server, k.Port)
	k.Headersjson = map[string]string{"content-type": "application/json"}
	k.Headersgzip = map[string]string{"content-type": "application/gzip"}

	return k
}
func SendRequest(url string, bodyText interface{}, headers map[string]string) (*http.Response, error) {
	jsonBody, err := json.Marshal(bodyText)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
