package main 
import "strings"

func map_to_string(m []string) string{
	s := ""
	for _,v := range m{
		s += v + ","
	}
	return strings.TrimSuffix(s,",")
}
func write_to_log(cmdd string, table string, s string){
	db.dblog = append(db.dblog, "\n" + stringify(cmdd, table, s))
	flush_lock(&wg)
	flush_value++
	flush_unlock(&wg)
}
func flush_log(){
	storageHandler("write")
}
func stringify(cmdd string, table string, s string) string{
	return "{\"Action\": " + `"` + cmdd + `"` + ", \"Table\": " + `"` + table + `"` + ", \"Value\": [" + s + "]}"
}
func checkFormSyntax(s string) bool{
	if s == "INIT" || s == "ADD" || s == "DELETE" || s == "READ" {
		return true
	}
	return false
} 