package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

type Urls struct {
	MyspaceUrl   string `json:"myspace_url,omitempty"`
	LastfmUrl    string `json:"lastfm_url,omitempty"`
	MbUrl        string `json:"mb_url,omitempty"`
	WikipediaUrl string `json:"wikipedia_url,omitempty"`
	OfficialUrl  string `json:"official_url,omitempty"`
}

type UrlsStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Version string `json:"version"`
}

type UrlResponse struct {
	Status UrlsStatus `json:"status"`
	Urls   Urls       `json:"urls"`
}

type UrlsSearchResults struct {
	Response UrlResponse `json:"response"`
}

func (u Urls) FormatUrls() string {

	st := reflect.ValueOf(u)
	fields_count := st.NumField()
	retrieved_urls := []string{}

	for i := 0; i < fields_count; i++ {
		if field_value := st.Field(i).String(); field_value != "" {
			retrieved_urls = append(retrieved_urls, field_value)
		}
	}

	if len(retrieved_urls) == 0 {
		return ""
	}

	formated_urls := ""
	for _, url := range retrieved_urls {
		formated_urls = fmt.Sprintf("%s%s\n", formated_urls, url)
	}

	return formated_urls

}

func FetchUrls(artist string) (Urls, error) {

	url := fmt.Sprintf("%s/urls?api_key=%s&name=%s&format=json", ApiRoot, ApiKey, url.QueryEscape(artist))
	results := UrlsSearchResults{}
	res, err := http.Get(url)

	eUrls := Urls{}
	if err != nil {
		return eUrls, err
	}

	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return eUrls, err
	}

	if results.Response.Status.Code == 5 {
		return eUrls, fmt.Errorf("Theres no band with this name: %s. Please try again", artist)
	}

	return results.Response.Urls, nil

}

func RunUrls(artist string) {
	urls, err := FetchUrls(artist)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
	fmt.Fprintf(os.Stdout, urls.FormatUrls())
}
