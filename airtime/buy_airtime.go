package airtime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type buyAirtime struct {
	Phone   string `json:"phone"`
	Amount  int    `json:"amount"`
	Network int    `json:"network"`
}

type buyAirtimeRes struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Amount  string `json:"amount"`
}

func BuyAirtime(phone string, amount, network_id int) (*buyAirtimeRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/buy-airtime", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := buyAirtime{}
	payload.Phone = phone
	payload.Amount = amount
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
	var response buyAirtimeRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
