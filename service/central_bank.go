package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quotation/model"
)

func FetchDataFromCentralBank() (*model.CBRResponse, error) {
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(model.CBRResponse)
	err = json.Unmarshal(respBody, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}