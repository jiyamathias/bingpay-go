package utilitybills

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type verifyCustomer struct {
	Service_Id  string `json:"service_id"`
	Customer_Id string `json:"customer_id"` //customer meter number, smart card number...
	Type        string `json:"type"`        //meter type (prepaid or postpaid) required for electricity bills only
}

type verifyCustomerRes struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func VerifyCustomer(service_id, customer_id, meter_type string) (*verifyCustomerRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/validate-service", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := verifyCustomer{}
	payload.Service_Id = service_id
	payload.Customer_Id = customer_id
	payload.Type = meter_type

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
	var response verifyCustomerRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
