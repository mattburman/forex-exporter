package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const urlf = "https://openexchangerates.org/api/latest.json?app_id=%s"

// requestFixture loads fixture.json and returns the response.
func requestFixture(appId string) (res response, err error) {
	file, err := ioutil.ReadFile("fixture.json")
	if err != nil {
		return res, errors.Wrap(err, "unable to read fixture.json file")
	}

	err = json.Unmarshal(file, &res)
	if err != nil {
		return res, errors.Wrap(err, "unable to unmarshal fixture.json")
	}

	return res, nil
}

// request makes a http request to the api using the api key.
// and returns the data.
// currently hardcoded to return a json file.
func request(appId string) (res response, err error) {
	resp, err := http.Get(fmt.Sprintf(urlf, appId))
	if err != nil {
		return res, errors.Wrap(err, "http error")
	}
	if resp.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("HTTP %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, errors.Wrap(err, "error reading response body")
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, errors.Wrap(err, "unable to unmarshal response body")
	}

	return res, nil
}

type response struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int                `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}
