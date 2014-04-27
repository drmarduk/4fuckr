package main

import (
	_ "fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(Url string) (string, error) {
	resp, err := http.Get(Url)
	if err != nil {
		return "", err
	}
	src, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	return string(src), nil
}
