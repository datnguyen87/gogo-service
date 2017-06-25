package service

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/unrolled/render"
	"bytes"
	"io/ioutil"
	"fmt"
	"strings"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)
const (
	fakeMatchLocationResult = "/matches/5a003b78-409e-4452-b456-a6f0dcee05bd"
)
func TestCreateMatch(t *testing.T) {
	client := &http.Client{}

	repo := newInMemoryRepository()
	server := httptest.NewServer(
		http.HandlerFunc(createMatchHandler(formatter, repo)),
	)

	defer server.Close()

	body := []byte(`{
	  "gridsize": 19,
	  "playerWhite": "bob",
	  "playerBlack": "alfred"
	}`)

	req, err := http.NewRequest(http.MethodPost, server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createMatchHandler: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createMatchHandler: %v", err)
	}

	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if location, ok := res.Header["Location"]; !ok {
		t.Error("Location header is not set")
	} else {
		if !strings.Contains(location[0], "/matches/") {
			t.Error("Location header should contain '/matches/'")
		}

		if len(location[0]) != len(fakeMatchLocationResult) {
			t.Error("Location value does not contain guid of new match")
		}
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received %s", res.Status)
	}

	fmt.Printf("Payload: %s", string(payload))
}