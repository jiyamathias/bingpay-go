package airtimetocash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type airtimeToCash struct {
	Amount  int    `json:"amount"`
	Network int    `json:"network"`
	Phone   string `json:"phone"`
}

type airtimeToCashRes struct {
	Error   bool                     `json:"error"`
	Message string                   `json:"message"`
	Data    airtimeToCashResDataBody `json:"data"`
}

type airtimeToCashResDataBody struct {
	Amount    string `json:"amount"`
	Value     int    `json:"value"`
	Reference int    `json:"reference"`
}

func AirtimeToCash(amount, network_id int, phone string) (*airtimeToCashRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/airtime-cash/process", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := airtimeToCash{}
	payload.Amount = amount
	payload.Network = network_id
	payload.Phone = phone

	jsonReq, jsonErr := json.Marshal(&payload)
	if jsonErr != nil {
		return nil, 0, jsonErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
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
	var response airtimeToCashRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
