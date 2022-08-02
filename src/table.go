package main

import (
	"fmt"
	"strings"
)

type DB struct {
	tables map[string]table
}

func (dbs DB) insert_to_db(t *table) {
	dbs.tables[t.name] = *t
}

type table struct {
	name     string
	capacity int
	size     int
	entries  map[string]entry
}

func make_table(name string) *table {
	t := new(table)
	t.name = name
	t.capacity = 200
	t.size = 0
	t.entries = make(map[string]entry)
	// t.entries["Title"] = make_entry("Value")
	return t
}
func (t table) insert_to_table(s string) {
	if(t.capacity > t.size){ 
		if _, found := t.find_entry_by_name(s); !found{ 
			t.make_entry(s)
			t.size++
		}
		return 
	}
	return 
}

func (t table) find_entry_by_name(s string) (entry, bool){
	empty := entry{}
	if val, ok := t.entries[s]; ok {
		return val, true
	}
	return empty, false 
}

func (t table) update_to_table(s string) {
	names := strings.Split(s, ">")
	oldname := names[0]
	newname := names[1]
	if 	_, found := t.find_entry_by_name(oldname); found{
		t.entries[oldname]= independent_make_entry(newname)
		return 
	}
	return
}

func independent_make_entry(s string) entry{
	var e entry
	e.cols = make(map[string]column)
	e.cols[s] = column{s} 
	return e
}

type column struct {
	fieldname string
}
type entry struct {
	cols map[string]column
}

func (t table) make_entry(s string){
	var newentry entry
	newentry.cols = make(map[string]column)
	t.entries[s] = newentry
	
	t.entries[s].cols[s] = column{s}
	fmt.Println("check ", s)
	return 
}
func (e entry) find_col_by_name(s string) (column, bool){
	empty := column{}
	if val, ok := e.cols[s]; ok {
		return val, true
	}
	return empty, false
}
func (e entry) insert_to_entry(s string) {
	if _, found := e.find_col_by_name(s); found{
		e.cols[s] = column{s}	
	}
}

func init_table(name string) *table {
	new_table := make_table(name)
	return new_table
}

var db DB

func init_db() {
	db = DB{}
	db.tables = make(map[string]table)
}
func (dbs DB) find_table(table_name string) *table {
	found_table, found := db.tables[table_name]
	if found {
		return &found_table
	}
	return nil
}