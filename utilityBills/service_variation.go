package utilitybills

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type serviceVariationRes struct {
	Error   bool                     `json:"error"`
	Message string                   `json:"message"`
	Data    serviceVariationDataBody `json:"data"`
}

type serviceVariationDataBody struct {
	ServiceName     string               `json:"service_name"`
	ServiceId       string               `json:"service_id"`
	Convinience_Fee string               `json:"convinience_fee"`
	Variations      []variationsDataBody `json:"variations"`
}

type variationsDataBody struct {
	Variation_Code   string `json:"variation_code"`
	Name             string `json:"name"`
	Variation_Amount string `json:"variation_amount"`
	FixedPrice       string `json:"fixed_price"`
}

func ServiceVariation(service_id int) (*serviceVariationRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/service/%d", client.BaseUrl, service_id)
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
	var response serviceVariationRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
