package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type EndpointResponse struct {
	Data    *DataFromResponse    `json:"data"`
	Support *SupportFromResponse `json:"support"`
}

type DataFromResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type SupportFromResponse struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

func CompareTwoStrings(str1, str2 string) (string, string) {
	op1 := getUniqueValuesFromStrings(str1, str2)
	op2 := getUniqueValuesFromStrings(str2, str1)
	return op1, op2
}

func getUniqueValuesFromStrings(str1, str2 string) string {
	str2Arr := strings.Split(str2, "")
	for _, val := range str2Arr {
		str1 = strings.ReplaceAll(str1, val, "")
	}
	return str1
}

func GetEmailFromTheResponseOfAPI(endpoint string) (string, error) {
	response, resErr := http.Get(endpoint)
	if resErr != nil {
		return "", resErr
	}
	decodedData, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return "", readErr
	}
	if decodedData == nil {
		return "", nil
	}

	endpointRes := EndpointResponse{}
	decodeErr := json.Unmarshal(decodedData, &endpointRes)
	if decodeErr != nil {
		return "", decodeErr
	}
	if endpointRes.Data != nil {
		data := endpointRes.Data
		return data.Email, nil
	}

	return "", nil
}
