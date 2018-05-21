package web

import "errors"

type site struct {
	title string
	indexPage page
	pages map[int]page
}

func Site(url string) site {
	p:=Page(url)
	s:=site{"", p, make(map[int]page)}
	return s
}

func (this *site) GetTitle() string {
	return this.title
}

func (this *site) GetIndexPage() page {
	return this.indexPage
}

func (this *site) GetPages() map[int]page {
	return this.pages
}

func (this *site) GetPage(position int) (page,error) {
	if position<len(this.pages) {
		return this.pages[position],nil
	}

	return page{}, errors.New("Index out of range")
}

func (this *site) SetTitle(title string) {
	this.title=title
}

func (this *site) SetIndexPage(indexPage page) {
	this.indexPage=indexPage
}

func (this *site) AddPage(p page) {
	this.pages[len(this.pages)]=p
}