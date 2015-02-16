package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Artist_Location(t *testing.T) {
	a := assert.New(t)

	body := `{
    "response": {
      "artist": {
        "artist_location": {
          "city": "Abingdon",
          "country": "United Kingdom",
          "location": "Abingdon, England, GB",
          "region": "England"
        },
        "id": "ARH6W4X1187B99274F",
        "name": "Radiohead"
      },
      "status": {
        "code": 0,
        "message": "Success",
        "version": "4.2"
      }
    }
	}`

	FakeServer(body, func() {
		location, err := RetrieveLocation("radiohead")
		a.NoError(err)
		a.Equal(location.City, "Abingdon")
		a.Equal(location.Country, "United Kingdom")
		a.Equal(location.Location, "Abingdon, England, GB")
		a.Equal(location.Region, "England")
	})

}

func Test_Artist_Location_Not_Found(t *testing.T) {
	a := assert.New(t)

	body := `{
    "response": {
      "status": {
          "code": 5,
          "message": "Success",
          "version": "4.2"
      }
    }
	}`

	FakeServer(body, func() {
		_, err := RetrieveLocation("joajdsllskd")
		a.Error(err)
		a.Equal("Theres no band with this name: joajdsllskd. Please try again", err.Error())
	})

}
