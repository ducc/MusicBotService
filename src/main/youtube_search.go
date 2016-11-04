package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SEARCH_REQUEST_URL = "https://www.googleapis.com/youtube/v3/search"
	INFO_REQUEST_URL   = "https://www.googleapis.com/youtube/v3/videos?"
)

var searchCache = make(map[string]*searchResponse)
var infoCache = make(map[string]*infoResponse)

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
	if cachedValue, ok := searchCache[url]; ok {
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
	searchCache[url] = &sResponse
	return &sResponse, err
}

func getSearchUrl(query string) (*string, error) {
	address, err := url.Parse(SEARCH_REQUEST_URL)
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

type infoResponse struct {
	Items []struct {
		Id             string `json:"id"`
		ContentDetails struct {
			Duration string   `json:"duration"`
			Blocked  []string `json:"blocked"`
		} `json:"contentDetails"`
	}
}

func info(url string) (*infoResponse, error) {
	if cachedValue, ok := infoCache[url]; ok {
		return cachedValue, nil
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var iResponse infoResponse
	err = json.NewDecoder(response.Body).Decode(&iResponse)
	if err != nil {
		return nil, err
	}
	infoCache[url] = &iResponse
	return &iResponse, err
}

func getInfoUrl(ids []string) (*string, error) {
	address, err := url.Parse(INFO_REQUEST_URL)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("part", "contentDetails")
	params.Add("id", strings.Join(ids, ","))
	params.Add("key", conf.Keys.Youtube)
	params.Add("type", "video")
	address.RawQuery = params.Encode()
	requestUrl := address.String()
	return &requestUrl, nil
}

type response struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	ChannelId    string   `json:"channel_id"`
	ChannelTitle string   `json:"channel_title"`
	Duration     string   `json:"duration"`
	Blocked      []string `json:"blocked"`
}

func youtubeSearchRoute(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("search")
	searchUrl, err := getSearchUrl(query)
	if err != nil {
		fmt.Println("Error getting search url,", err)
		return
	}
	sResponse, err := search(*searchUrl)
	if err != nil {
		fmt.Println("Error searching youtube,", err)
		return
	}
	items := sResponse.Items
	ids := make([]string, len(items))
	for index, value := range items {
		ids[index] = value.Id.VideoId
	}
	infoUrl, err := getInfoUrl(ids)
	if err != nil {
		fmt.Println("Error getting info url,", err)
		return
	}
	iResponse, err := info(*infoUrl)
	if err != nil {
		fmt.Println("Error getting info,", err)
		return
	}
	responses := make([]response, len(iResponse.Items))
	for index, info := range iResponse.Items {
		searchResult := sResponse.Items[index]
		snippet := searchResult.Snippet
		contentDetails := info.ContentDetails
		responses[index] = response{
			info.Id, snippet.Title, snippet.Description, snippet.ChannelId, snippet.ChannelTitle,
			contentDetails.Duration, contentDetails.Blocked,
		}
	}
	json.NewEncoder(writer).Encode(abstractResponse{false, API_VERSION, responses})
}
