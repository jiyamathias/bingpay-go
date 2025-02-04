package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type allDataPlansRes struct {
	Error   bool                      `json:"error"`
	Message string                    `json:"message"`
	Data    []allDataPlansResDataBody `json:"data"`
}

type allDataPlansResDataBody struct {
	Id         string `json:"id"`
	Network_Id string `json:"network_id"`
	Name       string `json:"name"`
	Price      string `json:"price"`
	Uniq_Id    string `json:"unique_id"`
}

func AllDataPlans() (*allDataPlansRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/all-data-plans", client.BaseUrl)
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

	resp_body, _ := io.ReadAll(resp.Body)
	var response allDataPlansRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
