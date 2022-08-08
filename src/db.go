package main

import "os"

type DB struct {
	tables map[string]*table
}
var first bool = true
var db *DB

// on system start
// init db on memory 
func init_db() {
	
	// start log
	if first == true {
		_, _ = os.Create("log.json")
		first = false
	}
	db = &DB{}
	db.tables = make(map[string]*table)
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
