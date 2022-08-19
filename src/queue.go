package main

import "time"



type Data struct{
	data string
	timestamp time.Time
}
type DataQueue struct{
	size int 
	capacity int 
	slots *[]Data // FIFO
}



func (q *DataQueue) pop() Data{
	if q.size == 0 {
		return Data{}
	}
	slots := q.slots
	ret_slot := (*slots)[0]
	(*slots) = (*slots)[1:]
	q.size--
	return ret_slot
}

func (q *DataQueue) push(d Data) string{
	
	if q.size < q.capacity { 
		slots := q.slots
		(*slots) = append((*slots), d)
		q.size++
		return "SUCCESS"
	}

	return "FAILURE"
}

func (q *DataQueue) top() Data{
	if q.size == 0 {
		return Data{}
	}
	slots := q.slots
	ret_slot := (*slots)[0]
	return ret_slot
}
