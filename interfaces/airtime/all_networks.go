package airtime

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go/interfaces"
)

type allNetworksRes struct {
	Error   bool
	Message string
	Data    []allNetworksResDataBody
}

type allNetworksResDataBody struct {
	ID   string
	Name string
	Note string
}

func AllNetworks() (*allNetworksRes, int, error) {
	client := interfaces.NewHttpClient()
	url := "https://bingpay.ng/api/v1/all-networks"
	method := "GET"
	token := client.Token
	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, respErr := client.Http.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response allNetworksRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
