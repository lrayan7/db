package main

type DB struct {
	tables map[string]*table
}
var db *DB
func init_db() {
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
func (q *Queue) q_insert(req Request){
	if(q.size == q.capacity){
		return
	}
  	q.slots[q.size] = req
	q.size++
}

