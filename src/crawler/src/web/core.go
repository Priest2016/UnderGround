package web

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type page struct {
	url string
	response *http.Response
	source string
}

func Page(url string) page{
	Page:=page{}

	response,err:=http.Get(url)

	if err!=nil{
		fmt.Println(err.Error())
	} else {
		defer response.Body.Close()

		bs,err:=ioutil.ReadAll(response.Body)

		if err!=nil{
			fmt.Println(err.Error())
		} else {
			content:=string(bs)

			Page.SetURL(url)
			Page.SetResponse(response)
			Page.SetSource(content)
		}
	}

	return Page
}

func (this *page) GetURL()string{
	return this.url
}

func (this *page) GetResponse() *http.Response {
	return this.response
}

func (this *page) GetSource()string{
	return this.source
}

func (this *page) SetURL(url string) {
	this.url=url
}

func (this *page) SetResponse(response *http.Response) {
	this.response=response
}

func (this *page) SetSource(source string) {
	this.source = source
}

func (this *page) UpdateResponse() {
	var err error
	var tmpReponse *http.Response
	tmpReponse,err=http.Get(this.url)

	if err!=nil{
		fmt.Println(err.Error())
	} else {
		this.response=tmpReponse
	}
}

func (this *page) UpdateSource() {
	this.UpdateResponse()

	bs,err:=ioutil.ReadAll(this.response.Body)

	if err!=nil{
		fmt.Println(err.Error())
	} else {
		this.source = string(bs)
	}
}