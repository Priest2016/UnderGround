package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	return response, err
}

func readResponse(response *http.Response) (string, error) {
	bs, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(bs), err
}
