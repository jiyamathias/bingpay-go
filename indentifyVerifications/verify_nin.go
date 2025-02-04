package indentifyverifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jiyamathias/bingpay-go"
)

type verifyNin struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Nin       string `json:"nin"`
}

type verifyNinRes struct {
	Error   bool                 `json:"error"`
	Message string               `json:"message"`
	Data    verifyNinResDataBody `json:"data"`
}

type verifyNinResDataBody struct {
	Nin              string                           `json:"nin"`
	Title            string                           `json:"title"`
	FirstName        string                           `json:"first_name"`
	LastName         string                           `json:"last_name"`
	MiddleName       string                           `json:"middle_name"`
	Phone            string                           `json:"phone"`
	BirthDate        string                           `json:"birth_date"`
	Nationality      string                           `json:"nationality"`
	Gender           string                           `json:"gender"`
	Profession       string                           `json:"profession"`
	StateOfOrigin    string                           `json:"state_of_origin"`
	LgaOfOrigin      string                           `json:"lga_of_origin"`
	PlaceOfOrigin    string                           `json:"place_of_origin"`
	Photo            string                           `json:"photo"`
	MaritalStatus    string                           `json:"marital_status"`
	Height           string                           `json:"height"`
	Email            string                           `json:"email"`
	EmploymentStatus string                           `json:"employment_status"`
	BirthState       string                           `json:"birth_state"`
	BirthCountry     string                           `json:"birth_country"`
	NextOfKin        verifyNinResDataBodyNextOfKin    `json:"next_of_kin"`
	NspokenLang      string                           `json:"n_spoken_lang"`
	OspokenLang      string                           `json:"o_spoken_lang"`
	ParentLastName   string                           `json:"parent_last_name"`
	Religion         string                           `json:"religion"`
	Residence        verifyNinResDataBodyResidence    `json:"residence"`
	Signature        string                           `json:"signature"`
	FieldMatches     verifyNinResDataBodyFieldMatches `json:"field_matches"`
}
type verifyNinResDataBodyNextOfKin struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	Lga        string `json:"lga"`
	State      string `json:"state"`
	Town       string `json:"town"`
}

type verifyNinResDataBodyResidence struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Town     string `json:"town"`
	Lga      string `json:"lga"`
	State    string `json:"state"`
	Status   string `json:"status"`
}

type verifyNinResDataBodyFieldMatches struct {
	LastName  bool `json:"last_name"`
	FirstName bool `json:"first_name"`
}

func VerifyNin(firstName, lastName, phone, nin string) (*verifyNinRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/verify/nin", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := verifyNin{}
	payload.FirstName = firstName
	payload.LastName = lastName
	payload.Phone = phone
	payload.Nin = nin

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
	var response verifyNinRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
