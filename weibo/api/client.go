package api

import (
	`fmt`
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

func Do(req *http.Request) (*http.Response, error) {
	fmt.Println("xxxxxxxxxxxx")
	return client.Do(req)
}
