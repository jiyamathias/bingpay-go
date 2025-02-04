package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type resloveAccount struct {
	Bank_Code      string `json:"bank_code"`
	Account_Number string `json:"account_number"`
}

type resloveAccountRes struct {
	Error bool                      `json:"error"`
	Data  resloveAccountResDataBody `json:"data"`
}
type resloveAccountResDataBody struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	ResponseCode  string `json:"response_code"`
	Message       string `json:"message"`
}

func ResolveAccount(bank_code, acct_num string) (*resloveAccountRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/resolve-account", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := resloveAccount{}
	payload.Bank_Code = bank_code
	payload.Account_Number = acct_num

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
	var response resloveAccountRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
