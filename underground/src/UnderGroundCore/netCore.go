/*
   ###   ###                                    ######
   ###   ###                                   #########
   ###   ###  $$      #####   %%%%%%  ####     ##     ##  ****     ####   @@  @@  ##      &&&&&
   ###   ###  $$$$$   ######  %%%%%%  ######   ##         ******  ######  @@  @@  #####   &&&&&&
   ###   ###  $$$$$$  ##  ##  %%      ##  ##   ##         **  **  ##  ##  @@  @@  ######  &&  &&
   ###   ###  $$  $$  ##  ##  %%%%%%  ######   ##  #####  ******  ##  ##  @@  @@  ##  ##  &&  &&
   ###   ###  $$  $$  ##  ##  %%      ####     ##     ##  ****    ##  ##  @@  @@  ##  ##  &&  &&
    #######   $$  $$  ######  %%%%%%  ## ##    #########  ** **   ######  @@@@@@  ##  ##  &&&&&&
     #####    $$  $$  #####   %%%%%%  ##  ##    ######    **  **   ####    @@@@   ##  ##  &&&&&

by PriestRussian

VK: http://vk.com/priestrussan
GMAIL: priestrussian@gmail.com
*/

package UnderGroundCore

import (
	"fmt"
	"errors"
	"net/http"
	"io/ioutil"
)

type Request struct {
	address string
	method string
}

func NewRequest(addr string, meth string) Request {
	return Request{addr, meth}
}

func (req *Request) GetAddr() string {
	return req.address
}

func (req *Request) GetMethod() string {
	return req.method
}

func (req *Request) SetAddress(addr string) {
	req.address = addr
}

func (req *Request) SetMethod(meth string) {
	req.method = meth
}

func (req *Request) Exec() (string, error) {
	var response string
	resp, err := http.Get(req.address)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			response = string(bytes)
		}
	}
	return response, err
}
