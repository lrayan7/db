package main

type column struct {
	fieldname string
}
type entry struct {
	line_number int
	cols map[string]*column
}

var Count int = 0 

func (e *entry) find_col_by_name(s string) (*column, bool){
	empty := column{}
	if val, ok := e.cols[s]; ok {
		return val, true
	}
	return &empty, false
}
func (e *entry) insert_to_entry(s string) {
	if _, found := e.find_col_by_name(s); found{
		e.cols[s] = &column{s}
	}
}
func independent_make_entry(s string) *entry{
	var e entry
	e.cols = make(map[string]*column)
	e.cols[s] = &column{s} 
	return &e
}