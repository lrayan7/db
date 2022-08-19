package main

import (
	// "fmt"
	"strings"
)

type table struct {
	name     string
	capacity int
	size     int
	entries  map[string]*entry
}
func make_table(name string) *table {
	t := new(table)
	t.name = name
	t.capacity = 200
	t.size = 0
	t.entries = make(map[string]*entry)
	// t.entries["Title"] = make_entry("Value")
	return t
}
func (t *table) insert_to_table(s string) {
	if(t.capacity > t.size){ 
		if _, found := t.find_entry_by_name(s); !found{ 
			t.make_entry(s)
			// fmt.Println("adding.. ", s)
			t.size++
			// fmt.Println("size now.. ", t.size)
		}
		return 
	}
	return 
}
func (t *table) find_entry_by_name(s string) (*entry, bool){
	var empty *entry = &entry{}
	if val, ok := t.entries[s]; ok {
		return val, true
	}
	return empty, false 
}
func (t *table) update_to_table(s string) {
	names := strings.Split(s, ">")
	oldname := names[0]
	newname := names[1]
	if 	_, found := t.find_entry_by_name(oldname); found{
		t.entries[oldname] = independent_make_entry(newname)
		return 
	}
	return
}
func (t *table) make_entry(s string){
	primekey := strings.Split(s, ",")[0]
	var newentry entry
	newentry.cols = make(map[string]*column)
	t.entries[primekey] = &newentry
	t.entries[primekey].line_number = Count
	Count++
	t.entries[primekey].cols[s] = &column{s}
	// fmt.Println("check ", primekey)
	return 
}
func init_table(name string) *table {
	new_table := make_table(name)
	return new_table
}

