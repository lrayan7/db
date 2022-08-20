package main

import (
	"net/url"
	"time"
)

type Channel struct{
	channel chan url.Values
	timestamp time.Time
}

type ChannelQueue struct{
	size int 
	capacity int 
	slots []Channel // FIFO
}

func (q *ChannelQueue) init() {
	q.size = 1
	q.capacity = MAX_WORKER_CAPACITY
	var newchannel Channel = Channel{}
	newchannel.channel = make(chan url.Values, 1)
	q.slots = append(q.slots, newchannel)
}

func (q *ChannelQueue) pop() Channel{
	if q.size == 0 {
		return Channel{}
	}
	slots := q.slots
	ret_slot := slots[0]
	slots = slots[1:]
	q.slots = slots
	q.size--
	return ret_slot
}
// delete by address
func (q *ChannelQueue) delete(c *Channel) {
	if q.size == 0 {
		return 
	}
	found := false
	slots := q.slots
	for k := range slots {
		if &slots[k] == c {
			found = true
			tmp1 := slots[:k]
			tmp2 := slots[k+1:]
			q.slots = append(tmp1, tmp2...) 
		}
	}
	if found == false {
		return 
	}
	q.size--
	return 
}

func (q *ChannelQueue) push(d *Channel) string{
	
	if q.size < q.capacity { 
		slots := q.slots
		slots = append(slots, (*d))
		q.slots = slots
		q.size++
		return "SUCCESS"
	}

	return "FAILURE"
}

func (q ChannelQueue) top() Channel{
	if q.size == 0 {
		return Channel{}
	}
	slots := q.slots
	ret_slot := slots[0]
	q.slots = slots
	return ret_slot
}
