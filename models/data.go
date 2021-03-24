package models

var DB []Item

type Item struct {
	a	int	`json:"a"`
	b	int	`json:"b"`
	c	int	`json:"c"`
	Nroots	int	
}