package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type buyData struct {
	Phone   string `json:"phone"`
	Plan    int    `json:"plan"`
	Network int    `json:"network"`
}

type buyDataRes struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func BuyData(phone string, plan, network_id int) (*buyDataRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/buy-data", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := buyData{}
	payload.Phone = phone
	payload.Plan = plan
	payload.Network = network_id

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
	var response buyDataRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
