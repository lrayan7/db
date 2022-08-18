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

type Channel struct{
	channel *chan string
	timestamp time.Time
}
func (c Channel) init() {
	*c.channel = make(chan string)
}
type ChannelQueue struct{
	size int 
	capacity int 
	slots *[]Channel // FIFO
}


func (q *ChannelQueue) pop() Channel{
	if q.size == 0 {
		return Channel{}
	}
	slots := q.slots
	ret_slot := (*slots)[0]
	(*slots) = (*slots)[1:]
	q.size--
	return ret_slot
}

func (q *ChannelQueue) delete(c *Channel) {
	if q.size == 0 {
		return 
	}
	found := false
	slots := (*q.slots)
	for k,_ := range slots {
		if &slots[k] == c {
			found = true
			tmp1 := slots[:k]
			tmp2 := slots[k+1:]
			(*q.slots) = append(tmp1, tmp2...) 
		}
	}
	if found == false {
		return 
	}
	q.size--
	return 
}

func (q *ChannelQueue) push(d Channel) string{
	
	if q.size < q.capacity { 
		slots := q.slots
		(*slots) = append((*slots), d)
		q.size++
		return "SUCCESS"
	}

	return "FAILURE"
}

func (q *ChannelQueue) top() Channel{
	if q.size == 0 {
		return Channel{}
	}
	slots := q.slots
	ret_slot := (*slots)[0]
	return ret_slot
}