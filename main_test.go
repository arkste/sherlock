package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestStatusCodeUserIsAvailable(t *testing.T) {
	socialNetwork := socialNetwork{
		ErrorType: errorStatusCode,
	}
	response := http.Response{
		StatusCode: 404,
	}
	if !isAvailable(&socialNetwork, &response) {
		t.Error("Username should be available")
	}
}

func TestStatusCodeUserIsNotAvailable(t *testing.T) {
	socialNetwork := socialNetwork{
		ErrorType: errorStatusCode,
	}
	response := http.Response{
		StatusCode: 200,
	}
	if isAvailable(&socialNetwork, &response) {
		t.Error("Username should not be available")
	}
}

func TestMessageUserIsAvailable(t *testing.T) {
	socialNetwork := socialNetwork{
		ErrorType: errorMessage,
		ErrorMsg:  "Not Found",
	}
	response := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("<html><body>Not Found</body></html>")),
	}
	if !isAvailable(&socialNetwork, &response) {
		t.Error("Username should be available")
	}
}

func TestMessageUserIsNotAvailable(t *testing.T) {
	socialNetwork := socialNetwork{
		ErrorType: errorMessage,
		ErrorMsg:  "Not Found",
	}
	response := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("<html><body>Foo</body></html>")),
	}
	if isAvailable(&socialNetwork, &response) {
		t.Error("Username should not be available")
	}
}

func TestResponseUrlUserIsAvailable(t *testing.T) {
	url, _ := url.Parse("http://foobar.com/index")
	request := http.Request{
		URL: url,
	}
	socialNetwork := socialNetwork{
		ErrorType: errorResponseURL,
		ErrorURL:  "http://foobar.com/index",
	}
	response := http.Response{
		Request: &request,
	}
	if !isAvailable(&socialNetwork, &response) {
		t.Error("Username should be available")
	}
}

func TestResponseUrlUserIsNotAvailable(t *testing.T) {
	url, _ := url.Parse("http://foobar.com/")
	request := http.Request{
		URL: url,
	}
	socialNetwork := socialNetwork{
		ErrorType: errorResponseURL,
		ErrorURL:  "http://foobar.com/index",
	}
	response := http.Response{
		Request: &request,
	}
	if isAvailable(&socialNetwork, &response) {
		t.Error("Username should not be available")
	}
}
