package main

import (
	"bytes"
	"fmt"
	"time"
)



type msg struct {
	Action string `json:"Action"`
	Table string `json:"Table"`
	Value string `json:"Value"`
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
func (q Queue) q_insert(req Request){
	if(q.size == q.capacity){
		return
	}
  	q.slots[q.size] = req
	q.size++
}
// -- interface to server.go 
func spawn_thread(thread_response chan string,  s []msg){
	if s[0].Action == "INIT"{
		fmt.Println("initiating new table in DB !")
		chosen := s[0].Table
		if _,found := db.tables[chosen]; !found{
			db.tables[chosen] = *make_table(chosen)
			// thread_response <- stringify(t./entries[s[0].Value])
		} 
		// if db.find_table(chosen) != nil {
		// 	fmt.Println("Success !")
		// }
	}
	if s[0].Action == "ADD"{
		// fmt.Println("table found= ", s[0].Value)
		chosen := s[0].Table
		if db.find_table(chosen) == nil {
			fmt.Println("Table was not found, Creating new one")
			db.tables[chosen] = *make_table(chosen)

		}else{
			
			db.find_table(chosen).insert_to_table(s[0].Value) 
			
		}

	}
	if s[0].Action == "READ"{
		chosen := s[0].Table
		if t := db.find_table(chosen); t == nil {
			fmt.Println("Table was not found, Creating new one")
			db.tables[chosen] = *make_table(chosen)
		}else{ 
			if _, found := t.find_entry_by_name(s[0].Value); found{
				thread_response<- stringify(t.entries[s[0].Value])
				return
			} 
		}
	}
	if s[0].Action == "UPDATE"{
		chosen := s[0].Table
		if t := db.find_table(chosen); t == nil {
			fmt.Println("Table was not found, Creating new one")
			db.tables[chosen] = *make_table(chosen)
		}else{ t.update_to_table(s[0].Value) }
	}
	thread_response <- "true"
}

func stringify(e entry) string{
	
	i := 0
	var buffer bytes.Buffer
	buffer.WriteString(`"`)
	for _,col := range e.cols {
		if i > 0 {
			buffer.WriteString(`,`)			
		}
		buffer.WriteString(string(i))
		buffer.WriteString(`":"`)	
		fmt.Println(col.fieldname, " and ")	
		buffer.WriteString(col.fieldname)
		buffer.WriteString(`"`)
		i++
	}
	
	fmt.Println("[{" + buffer.String() + "}]")
	return "[{" + buffer.String() + "}]" 
}

