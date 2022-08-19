package main

import (
	"io"
	"log"
	"os"
)

// serialize map of strings
/*
	table_name	<index>

*/
type index struct{
	table_name string
}
func storageHandler(cmdd string){
	if cmdd == "write" {
		f2, _:= os.OpenFile("storage.json", os.O_APPEND | os.O_RDWR, 0644) 
		// put synch of log files here 
		// nned to save timestamps of queries before that  
		for _,v := range db.dblog{
			_, err2 := io.WriteString(f2, v + "\n")
			if err2 != nil{
				log.Fatal(err2)
			}
		}
		// refresh log 
		db.dblog = nil
		db.dblog = make([]string, 1)
		db.dblog[0] = ""
	}
}

