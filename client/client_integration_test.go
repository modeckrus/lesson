package client

import (
	"net/http"
	"testing"
)

var client = Client{
	Addr:   "http://localhost:3333",
	Client: http.DefaultClient,
}

func TestPint(t *testing.T) {
	result, err := client.Ping()
	if err != nil {
		t.Fatal(err)
	}
	if result != "pong" {
		t.Fail()
	}
}
