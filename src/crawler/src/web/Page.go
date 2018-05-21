package web

import (
	"net/http"
)

type page struct {
	url      string
	response *http.Response
	source   string
}

func Page(url string) page {
	Page := page{}

	response, err := sendRequest(url)

	if err == nil {
		content, err := readResponse(response)
		if err == nil {
			Page.SetURL(url)
			Page.SetResponse(response)
			Page.SetSource(content)
		}
	}

	return Page
}

func (this *page) GetURL() string {
	return this.url
}

func (this *page) GetResponse() *http.Response {
	return this.response
}

func (this *page) GetSource() string {
	return this.source
}

func (this *page) SetURL(url string) {
	this.url = url
}

func (this *page) SetResponse(response *http.Response) {
	this.response = response
}

func (this *page) SetSource(source string) {
	this.source = source
}

func (this *page) UpdateResponse() {
	tmpResponse, err := sendRequest(this.url)

	if err == nil {
		this.response = tmpResponse
	}
}

func (this *page) UpdateSource() {
	this.UpdateResponse()

	tmpSource, err := readResponse(this.response)

	if err == nil {
		this.source = tmpSource
	}
}
