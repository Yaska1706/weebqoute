package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRequestHandler(url string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}

func GetSingleRequest(url string) QouteAnime {
	bodyBytes := GetRequestHandler(url)

	var responseObject QouteAnime
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject

}
func GetAllRequest(url string) []QouteAnime {
	bodyBytes := GetRequestHandler(url)

	var responseObject []QouteAnime
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject
}

func GetAllAnimeRequest(url string) []string {
	bodyBytes := GetRequestHandler(url)
	var responseObject []string
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject
}
