package main

import (
	"fmt"
	"net/url"
	"strings"
)


func main_supervisor(){
	
	for{
		
		select{
				// new item added
		case newitem := <- listen_on_item_channel :
			
			if newitem == "new" { 
				
				// create Channel struct on mem
				var newchannel *Channel = &Channel{}
				newchannel.channel = make(chan url.Values, 1)
				// send item to new worker's channel
				item_to_be_sent :=<- item_channel
				
				select {
					case newchannel.channel <- item_to_be_sent :
						// push new channel to channel_queue
						ret := (&worker_channel).push(newchannel)
						if ret == "SUCCESS" { 
							go func() {
								worker(newchannel)	
							}()
						}
				}
			}
		case action := <- terminate :
			if action { return }
		}
	}
}


func worker(thread_channel *Channel){

	running_workers++
	s :=<- thread_channel.channel
	str := map_to_string(s["Value"]) // parse values
	chosenTable := s["Table"][0]
	action := s["Action"][0]
	
	for { 	
		switch action { 
		case "INIT" :
			fmt.Println("initiating new table in DB !")
			
			if _,found := db.tables[chosenTable]; !found{
				db.tables[chosenTable] = make_table(chosenTable)
			}
			goto FINISHED_PROCESSING
		case "ADD" :
			if db.find_table(chosenTable) == nil {
				fmt.Println("Table was not found !")
			}else{ 
				db.find_table(chosenTable).insert_to_table(str)
			}
			goto FINISHED_PROCESSING
		case "READ" :
			if t := db.find_table(chosenTable); t == nil {
				fmt.Println("Table was not found !")
			}else{ 
				if _, found := t.find_entry_by_name(strings.Split(str, ",")[0]); found{
					return
				} 
			}
			goto FINISHED_PROCESSING
		}
	}
FINISHED_PROCESSING:	
	
	write_to_log(action, chosenTable, str);
	running_workers--
	(&worker_channel).delete(thread_channel)
}

const MAX_FLUSH_VALUE = 3
var flush_value int = 0
func log_write() {
	for{
		if flush_value >= MAX_FLUSH_VALUE {
			flush_log()
		}
	}
}

