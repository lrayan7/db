package main

import (
	"net/url"
	"os"
	"sync"
)

type DB struct {
	tables map[string]*table
	dblog []string
}

const MAX_WORKER_CAPACITY = 100

var wg sync.WaitGroup
var terminate chan bool 
var finished_init chan int
var running_workers int 
var item_channel chan url.Values 
var listen_on_item_channel chan string 
var worker_channel ChannelQueue

// on system start
// init db on memory 
var db *DB
func init_db() {

	listen_on_item_channel = make(chan string)
	item_channel = make(chan url.Values)
	running_workers = 0
	terminate = make(chan bool)
	// start log
	_, _ = os.Create("storage.json")
	// defer close
	worker_channel.init()
	db = &DB{}
	db.tables = make(map[string]*table)
	db.dblog = make([]string, 1)
	db.dblog[0] = "[{ \"State\": \"init log\" }]"
	// supervisor start
    go func() {
		main_supervisor()
	}()
	// supervisor start
    go func() {
		log_write()
	}()
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


