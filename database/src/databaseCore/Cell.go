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

07.04.2018
*/

package databaseCore

import (
	"errors"
	"reflect"
)

type Cell struct {
	cellName  string
	cellType  CellType
	cellValue interface{}
}

func NewEmptyCell() Cell {
	return Cell{}
}

func NewCell(cname string, ctype CellType, cvalue interface{}) Cell {
	return Cell{cname, ctype, cvalue}
}

func (this *Cell) GetName() string {
	return this.cellName
}

func (this *Cell) GetType() CellType {
	return this.cellType
}

func (this *Cell) GetTypeAsString() string {
	if this.cellType == INT {
		return "int"
	} else if this.cellType == STRING {
		return "string"
	}

	return ""
}

func (this *Cell) GetValue() interface{} {
	return this.cellValue
}

func (this *Cell) SetName(cname string) {
	this.cellName = cname
}

func (this *Cell) SetType(ctype CellType) {
	this.cellType = ctype
}

func (this *Cell) SetValue(cvalue interface{}) error {
	if this.cellType == INT && reflect.TypeOf(cvalue).String() == "int" {
		this.cellValue = cvalue

		return nil
	} else if this.cellType == STRING && reflect.TypeOf(cvalue).String() == "string" {
		this.cellValue = cvalue

		return nil
	}

	return errors.New("Couldn't assign type {" + reflect.TypeOf(cvalue).String() + "} to {" + this.GetTypeAsString() + "}")
}
