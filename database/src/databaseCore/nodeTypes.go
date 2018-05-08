package databaseCore

type NodeType int

const (
	INTEGER NodeType = iota
	FLOAT
	STRING
	TEXT
)
