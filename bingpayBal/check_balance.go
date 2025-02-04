package bingpayBal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type checkBalanceRes struct {
	Error   bool                    `json:"error"`
	Message string                  `json:"message"`
	Data    checkBalanceResDataBody `json:"data"`
}

type checkBalanceResDataBody struct {
	Balance  string `json:"balance"`
	Currency string `json:"currency"`
}

func CheckBalance() (*checkBalanceRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/self/balance", client.BaseUrl)
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
	var response checkBalanceRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
