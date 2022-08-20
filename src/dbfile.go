package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"
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
		write_in_synch(f2, db.dblog)
		// refresh log 
		db.dblog = nil
		db.dblog = make([]string, 1)
		db.dblog[0] = ""
		flush_lock(&wg)
		flush_value = 0
		flush_unlock(&wg)
	}
}

func write_in_synch(f2 *os.File, dblog []string) {
	sort.Slice(dblog, func(i,j int)bool {
		// first 8 chars are the timestamp: hour:time:second
		// 20:40:58 ... 
		hour1 := strings.Index(dblog[i], ":")
		minute1 := strings.Index(dblog[i][hour1 + 1:], ":")
		second1 := strings.Index(dblog[i][minute1 + 1:], ":")

		hour2 := strings.Index(dblog[i], ":")
		minute2 := strings.Index(dblog[i][hour2 + 1:], ":")
		second2 := strings.Index(dblog[i][minute2 + 1:], ":")
		
		if hour1 < hour1 {
			return true
		}else if hour1 == hour2 {
			if minute1 < minute2 {
				return true
			}else if minute1 == minute2 {
				if second1 < second2 {
					return true 
				}
			}
		}
		return false
	})
	for _,v := range dblog{
		_, err2 := io.WriteString(f2, v + "\n")
		if err2 != nil{
			log.Fatal(err2)
		}
	}
	
}


