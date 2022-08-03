package main

import "time"


type msg struct {
	Action string `json:"Action"`
	Table string `json:"Table"`
	Value map[string]interface{} `json:"-"`
}
type Request struct{
	request string
	timestamp time.Time
}
type Queue struct{
	size int 
	capacity int 
	slots []Request // FIFO
}