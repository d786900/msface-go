package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

var uri = "http://localhost:5000/face/v1.0/detect"

func getEmotion(m io.Reader) (string, error) {
	client := &http.Client{}

	// only request emotion detection
	params := "*"
	request, err := http.NewRequest("POST", uri+"?returnFaceAttributes="+params, m)
	request.Header.Add("Content-Type", "application/octet-stream")

	// Send the request to the local web service
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	// read response
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response interface{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(response, "", "\t")

	return string(b), nil
}
