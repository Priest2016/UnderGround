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

package main

import (
	// dbc "databaseCore"
	"fmt"
)

type Node struct {
	nodeName   string
	nodeParent *Node
}

func CreateNode(name string, parent *Node) Node {
	return Node{name, parent}
}

func main() {
	node := CreateNode("1", &Node{})
	nodee := CreateNode("2", &node)
	fmt.Println(nodee)
	fmt.Println(*nodee.nodeParent)
}
