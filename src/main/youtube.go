package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const REQUEST_URL = "https://www.googleapis.com/youtube/v3/search"

var cache = make(map[string]*searchResponse)

type searchResponse struct {
	Items []struct {
		Id struct {
			VideoId string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			ChannelId    string `json:"channelId"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			ChannelTitle string `json:"channelTitle"`
		} `json:"snippet"`
	} `json:"items"`
}

func search(url string) (*searchResponse, error) {
	if cachedValue, ok := cache[url]; ok {
		return cachedValue, nil
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var sResponse searchResponse
	err = json.NewDecoder(response.Body).Decode(&sResponse)
	if err != nil {
		return nil, err
	}
	return &sResponse, err
}

func getUrl(query string) (*string, error) {
	address, err := url.Parse(REQUEST_URL)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("q", query)
	params.Add("key", conf.Keys.Youtube)
	address.RawQuery = params.Encode()
	requestUrl := address.String()
	return &requestUrl, nil
}

func youtubeSearchRoute(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("search")
	requestUrl, err := getUrl(query)
	if err != nil {
		fmt.Println("Error getting url,", err)
		return
	}
	response, err := search(*requestUrl)
	if err != nil {
		fmt.Println("Error searching youtube,", err)
		return
	}
	json.NewEncoder(writer).Encode(abstractResponse{false, API_VERSION, response.Items})
}
