package bank

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type allBanksRes struct {
	Error   bool                  `json:"error"`
	Message string                `json:"message"`
	Data    []allBanksResDataBody `json:"data"`
}

type allBanksResDataBody struct {
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Code        string      `json:"code"`
	LongCode    string      `json:"longcode"`
	Gateway     interface{} `json:"gateway"`
	PayWithBank bool        `json:"pay_with_bank"`
	Active      bool        `json:"active"`
	IsDeleted   bool        `json:"is_deleted"`
	Country     string      `json:"country"`
	Currency    string      `json:"currency"`
	Type        string      `json:"type"`
	Id          int         `json:"id"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

func AllBanks() (*allBanksRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/all-banks", client.BaseUrl)
	method := "POST"
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
	var response allBanksRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
