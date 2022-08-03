package main

import (
	"bytes"
	"fmt"
	"strings"

	// "reflect"
	"strconv"
	// "strings"
)

func map_to_string(m map[string]interface{}) string{
	s := ""
	for _,v := range m{
		s += fmt.Sprint(v) + ","
	}
	return s
}
// -- interface to server.go
func spawn_thread(thread_response chan string,  s msg){
	if s.Action == "INIT"{
		fmt.Println("initiating new table in DB !")
		chosen := s.Table
		if _,found := db.tables[chosen]; !found{
			db.tables[chosen] = make_table(chosen)
		} 
		return
	}
	str := map_to_string(s.Value)
	// writer
	if s.Action == "ADD"{
		// take lock 
		take_lock(s.Action + s.Table + str, &wg)
		chosen := s.Table
		if db.find_table(chosen) == nil {
			fmt.Println("Table was not found !")
		}else{ 
			db.find_table(chosen).insert_to_table(str)
		}
		give_lock(&wg)
		return
	}
	// multiple readers 
	if s.Action == "READ"{
		chosen := s.Table
		if t := db.find_table(chosen); t == nil {
			fmt.Println("Table was not found !")
		}else{ 
			if e, found := t.find_entry_by_name(strings.Split(str, ",")[0]); found{
				thread_response<- stringify(t, e)
				return
			} 
		}
		thread_response <- "blank"
		return
	}
	// deal with it later - strings.Split(str, ",")[0]
	if s.Action == "UPDATE"{
		take_lock(s.Action + s.Table + str, &wg)
		chosen := s.Table
		if t := db.find_table(chosen); t == nil {
			fmt.Println("Table was not found !")	
		}else{ t.update_to_table(str) }
		give_lock(&wg)
		return
	}
	thread_response <- "blank"
}
func stringify(t *table, e *entry) string{
	i := 0
	var buffer bytes.Buffer
	buffer.WriteString(`"`)
	for _,col := range e.cols {
		if i > 0 {
			buffer.WriteString(`,`)			
		}else{ i = e.line_number }
		buffer.WriteString(strconv.Itoa(i))
		buffer.WriteString(`":"`)	
		fmt.Println(col.fieldname, " and ")	
		buffer.WriteString(col.fieldname)
		buffer.WriteString(`"`)
		i++
	}
	fmt.Println("[{" + buffer.String() + "}]")
	return "[{" + buffer.String() + "}]" 
}

