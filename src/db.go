package main

import (
	"fmt"
	"net/url"
	"os"
)


type DB struct {
	tables map[string]*table
	dblog []string
}
var first bool = true
var db *DB
var item_channel chan url.Values = make(chan url.Values)
// on system start
// init db on memory 
func init_db() {
	
	// start log
	if first == true {
		_, _ = os.Create("storage.json")
		// defer f.Close()
		// _, err2 := io.WriteString(f, "{ \"LOG\": \"\" }")
		// if err2 != nil{
		// 	log.Fatal(err2)
		// }
		first = false
	}
	db = &DB{}
	db.tables = make(map[string]*table)
	db.dblog = make([]string, 1)
	terminate = make(chan bool)
	db.dblog[0] = "[{ \"State\": \"init log\" }]"
	// supervisor start
    go func(){
		main_supervisor()
	}()
}

func main_supervisor(){
	// var item_2b_sent url.Values
	// var response_msg string 
	// var thread_response chan string
	for{
		if action := <- terminate; action {
			break
		}
		for i :=0; i<3; i++ {
			item_2b_sent := <- item_channel
			thread_response := make(chan string, 3) 
			// make(chan Issue, len(allIssues))
			// after 3 writes log will be written to storage 
			// forks worker thread
			go func(){
				spawn_thread(thread_response, item_2b_sent)
			}()
			_ = <- thread_response
		}
		fmt.Println("now !")
		write_from_log()
	}
}

func (dbs *DB) insert_to_db(t *table) {
	dbs.tables[t.name] = t
}
func (dbs *DB) find_table(table_name string) *table {
	found_table, found := db.tables[table_name]
	if found {
		return found_table
	}
	return nil
}
func (q *Queue) q_insert(req Request) {
	if q.size == q.capacity {
		return
	}
	q.slots[q.size] = req
	q.size++
}
