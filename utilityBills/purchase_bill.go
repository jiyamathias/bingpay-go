package utilitybills

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type purchaseBill struct {
	Service_Id  string `json:"service_id"`
	Customer_Id string `json:"customer_id"`
	Variation   string `json:"variation"`
	Amount      string `json:"amount"`
}

type purchaseBillRes struct {
	Error          bool                    `json:"error"`
	Message        string                  `json:"message"`
	Data           purchaseBillResDataBody `json:"data"`
	Purchased_Code interface{}             `json:"purchased_code"`
}

type purchaseBillResDataBody struct {
	Status          string      `json:"status"`
	Product_Name    string      `json:"product_name"`
	Unique_Element  string      `json:"enique_element"`
	Unit_Price      int         `json:"unit_price"`
	Quantity        int         `json:"quantity"`
	Channel         string      `json:"channel"`
	Type            string      `json:"type"`
	Phone           string      `json:"phone"`
	Name            interface{} `json:"name"`
	Convinience_fee int         `json:"convinience_fee"`
	Amount          int         `json:"amount"`
	Platform        string      `json:"platform"`
}

func PurchaseBill(service_Id, customer_Id, variation, amount string) (*purchaseBillRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/purchase-bill", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := purchaseBill{}
	payload.Service_Id = service_Id
	payload.Customer_Id = customer_Id
	payload.Variation = variation
	payload.Amount = amount

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
	var response purchaseBillRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
