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

type Table struct {
	tableName  string
	tableCells map[int]map[string]CellType
	tableRows  map[int]Row
}

func NewEmptyTable() Table {
	return Table{}
}

func NewTable(tname string, tcells map[int]map[string]CellType) Table {
	return Table{tname, tcells, make(map[int]Row)}
}

func (this *Table) GetName() string {
	return this.tableName
}

func (this *Table) GetCells() map[int]map[string]CellType {
	return this.tableCells
}

func (this *Table) GetRows() map[int]Row {
	return this.tableRows
}
