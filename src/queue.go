package main

import "time"



type Request struct{
	request string
	timestamp time.Time
}
type Queue struct{
	size int 
	capacity int 
	slots []Request // FIFO
}