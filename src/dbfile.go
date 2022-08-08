package main

import (
	// "encoding/json"
	// "io/ioutil"
	"fmt"
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
func fileHandler_(cmdd string, table_name string, s string){
	if cmdd == "write" {
		fmt.Println("HERE? ? ? ")
		f2, _:= os.OpenFile("log.json", os.O_RDWR, 0644) 
		_, err2 := f2.WriteString(s)
		if err2 != nil{
			log.Fatal(err2)
		}
	}
}

