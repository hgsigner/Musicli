package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Artist_With__All_Urls(t *testing.T) {
	assert := assert.New(t)

	body := `{
	  "response": {
	    "status": {
	      "code": 0,
	      "message": "Success",
	      "version": "4.2"
	    },
	    "urls": {
	        "myspace_url": "http://www.myspace.com/radiohead",
	        "lastfm_url": "http://www.last.fm/music/Radiohead",
	        "mb_url": "http://musicbrainz.org/artist/a74b1b7f-71a5-4011-9441-d0b5e4122711.html",
	        "wikipedia_url": "http://en.wikipedia.org/wiki/Radiohead",
	        "official_url": "http://radiohead.com"
	    }
	  }
	}`

	FakeServer(body, func() {
		urls, err := FetchUrls("radiohead")
		assert.NoError(err)
		assert.Equal("http://www.myspace.com/radiohead", urls.MyspaceUrl)
	})

}

func Test_Artist_Not_Found(t *testing.T) {
	assert := assert.New(t)

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
		_, err := FetchUrls("jahsfks")
		assert.Error(err)
		assert.Equal("Theres no band with this name: jahsfks. Please try again", err.Error())
	})

}

func Test_Artist_Partial_Urls(t *testing.T) {
	assert := assert.New(t)

	body := `{
	  "response": {
	    "status": {
	      "code": 0,
	      "message": "Success",
	      "version": "4.2"
	    },
	    "urls": {
	        "myspace_url": "http://www.myspace.com/radiohead",
	        "wikipedia_url": "http://en.wikipedia.org/wiki/Radiohead",
	        "official_url": "http://radiohead.com"
	    }
	  }
	}`

	FakeServer(body, func() {
		urls, err := FetchUrls("radiohead")
		assert.NoError(err)
		assert.Empty(urls.MbUrl)
		assert.Empty(urls.LastfmUrl)
	})

}

func FakeServer(body string, f func()) {
	root := ApiRoot

	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, body)
	}))

	defer ts.Close()

	ApiRoot = ts.URL

	f()

	ApiRoot = root
}
