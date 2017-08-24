package main

import (
	"net/http"
	"testing"
	"io/ioutil"
	"time"
)

func TestTribHandler(t *testing.T) {
	go startServer()
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	r, _ := http.NewRequest("GET", "http://localhost:8081/v1/trib/10", nil)
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"result":"44"}`
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	if bodyString != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			bodyString, expected)
	}
}
