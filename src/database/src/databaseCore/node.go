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

VK: https://vk.com/priestrussian
GMAIL: priestrussian@gmail.com

08.05.2018
*/

package databaseCore

import (
	"errors"
	"reflect"
)

type node struct {
	nodeName      string
	nodeType      NodeType
	nodeData      interface{}
	nodeParent    *node
	nodeChildrens map[int]*node
}

func Node() node {
	return node{}
}

// Getters start
func (this *node) GetName() string {
	return this.nodeName
}

func (this *node) GetType() NodeType {
	return this.nodeType
}

func (this *node) GetData() interface{} {
	return this.nodeData
}

func (this *node) GetParent() *node {
	return this.nodeParent
}

func (this *node) GetChilds() map[int]*node {
	return this.nodeChildrens
}

// Getters end

// Setters start
func (this *node) SetName(name string) {
	this.nodeName = name
}

func (this *node) SetType(nType NodeType) {
	this.nodeType = nType
}

func (this *node) SetData(data interface{}) error {
	udType := reflect.TypeOf(data).String()
	if udType == "string" && this.nodeType == STRING {
		this.nodeData = data
		return nil
	} else if udType == "int" && this.nodeType == INTEGER {
		this.nodeData = data
		return nil
	} else if udType == "float64" && this.nodeType == FLOAT {
		this.nodeData = data
		return nil
	} else if udType == "string" && this.nodeType == TEXT {
		this.nodeData = data
		return nil
	} else {
		return errors.New("Uncnown data type")
	}
}

func (this *node) SetParent(parent *node) {
	this.nodeParent = parent
}

func (this *node) AddChild(child *node) {
	this.nodeChildrens[len(this.nodeChildrens)] = child
}

// Setters end
