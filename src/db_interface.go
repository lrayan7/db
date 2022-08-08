package main

import (
	// "bytes"
	"fmt"
	"net/url"
	// "os"
	"strings"
	// "strconv"
)
func map_to_string(m []string) string{
	s := ""
	for _,v := range m{
		s += v + ","
	}
	return strings.TrimSuffix(s,",")
}

// -- interface to server.go
func spawn_thread(thread_response chan string,  s url.Values /*msg*/){
	str := map_to_string(s["Value"]) // ..., ..., ..., ..., ...,
	if s["Action"][0] == "INIT"{
		fmt.Println("initiating new table in DB !")
		chosen := s["Table"][0]
		if _,found := db.tables[chosen]; !found{
			db.tables[chosen] = make_table(chosen)
			write_to_log("INIT", chosen, str);
		}
		fmt.Println("HERE interface")
		thread_response <- "blank" 
		
		return
	}
	// fmt.Println("here ", str);
	if s["Action"][0] == "ADD"{
		take_lock(s["Action"][0] + s["Table"][0] + str, &wg)
		chosen := s["Table"][0]
		if db.find_table(chosen) == nil {
			fmt.Println("Table was not found !")
		}else{ 
			db.find_table(chosen).insert_to_table(str)
			write_to_log("ADD", chosen, str);
		}
		give_lock(&wg)
		thread_response <- "blank"
		return
	}
	// multiple readers 
	if s["Action"][0] == "READ"{
		chosen := s["Table"][0]
		if t := db.find_table(chosen); t == nil {
			fmt.Println("Table was not found !")
		}else{ 
			if _, found := t.find_entry_by_name(strings.Split(str, ",")[0]); found{
				thread_response<- "blank"
				return
			} 
		}
		thread_response <- "blank"
		return
	}
	// deal with it later - strings.Split(str, ",")[0]
	if s["Action"][0] == "UPDATE"{
		take_lock(s["Action"][0] + s["Table"][0] + str, &wg)
		chosen := s["Table"][0]
		if t := db.find_table(chosen); t == nil {
			fmt.Println("Table was not found !")	
		}else{ t.update_to_table(str) }
		give_lock(&wg)
		thread_response <- "blank"
		return
	}
}
func write_to_log(cmdd string, table string, s string){
	db.dblog = append(db.dblog, "\n" + stringify(cmdd, table, s))
}
func write_from_log(){
	storageHandler("write")
}
func stringify(cmdd string, table string, s string) string{
	return "{\"Action\": " + `"` + cmdd + `"` + ", \"Table\": " + `"` + table + `"` + ", \"Value\": [" + s + "]}"
}


















// old stringify 
// i := 0
// 	var buffer bytes.Buffer
// 	buffer.WriteString(`"`)
// 	for _,col := range e.cols {
// 		if i > 0 {
// 			buffer.WriteString(`,`)			
// 		}else{ i = e.line_number }
// 		buffer.WriteString(strconv.Itoa(i))
// 		buffer.WriteString(`":"`)	
// 		fmt.Println(col.fieldname, " and ")	
// 		buffer.WriteString(col.fieldname)
// 		buffer.WriteString(`"`)
// 		i++
// 	}
// 	fmt.Println("[{" + buffer.String() + "}]")
// 	return "[{" + buffer.String() + "}]" 