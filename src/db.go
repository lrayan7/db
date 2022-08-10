package main

import (
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
var listen_on_item_channel chan int = make(chan int)
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
	db.dblog[0] = "[{ \"State\": \"init log\" }]"
	// supervisor start
	req_worker := make(chan string)
	listen_on_item_channel <- 0
	terminate = make(chan bool)
	terminate <- false
	// no requests on item channel
	listen_on_item_channel <- 0
    go func(){
		listener(req_worker)
		main_supervisor(req_worker)
	}()
}

func main_supervisor(req_worker chan string){
	// var item_2b_sent url.Values
	// var response_msg string 
	var thread_response chan string = make(chan string)
	for{
		select{ 
		case action := <- terminate :
			if action { return }
		case listen := <- req_worker :
			if listen == "yes" {
				go func(){
					worker(thread_response)
				}()		
			}
		}
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

func listener(req_worker chan string){
	for{
		select { 
		case wait_queue_size := <-listen_on_item_channel :
			if wait_queue_size >= LOAD_FACTOR {
				req_worker <- "yes"
			}
		}
	}

}