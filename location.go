package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ArtistLocation struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	Location string `json:"location"`
	Region   string `json:"region"`
}

type Artist struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	ArtistLocation ArtistLocation `json:"artist_location"`
}

type LocationStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Version string `json:"-"`
}

type LocationResponse struct {
	Artist Artist         `json:"artist"`
	Status LocationStatus `json:"status"`
}

type LocationResults struct {
	Response LocationResponse `json:"response"`
}

func RetrieveLocation(artist string) (ArtistLocation, error) {

	url := fmt.Sprintf("%s/profile?api_key=%s&name=%s&format=json&bucket=artist_location", ApiRoot, ApiKey, url.QueryEscape(artist))
	results := LocationResults{}
	res, err := http.Get(url)

	al := ArtistLocation{}
	if err != nil {
		return al, err
	}

	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return al, err
	}

	if results.Response.Status.Code == 5 {
		return al, fmt.Errorf("Theres no band with this name: %s. Please try again", artist)
	}

	return results.Response.Artist.ArtistLocation, nil

}

func RunLocation(art string, out io.Writer) {
	location, err := RetrieveLocation(art)
	if err != nil {
		fmt.Fprintln(out, err)
	}
	fmt.Fprintf(out, FormatStructToText(location))
}
