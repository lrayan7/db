package main

import (
	// "fmt"
	// "fmt"
	"net/http"
	// "net/url"
	// "os"
	"sync"
)

// to terminate db process
var terminate chan bool 
var db_msg string 
var wg sync.WaitGroup
var LOAD_FACTOR int = 5
// into Storage
var FLUSH_TIME int = 3


func main() {
    init_db()
    http.HandleFunc("/homepage", homepageHandler) // Update this line of code
    http.HandleFunc("/dbgo", dbReqHandler) // Update this line of code
    http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("static"))))
    http.ListenAndServe(":8080", nil)
}
func homepageHandler(w http.ResponseWriter, r *http.Request){ 
  http.ServeFile(w, r, "static/index.html")
}
func dbReqHandler(w http.ResponseWriter, r *http.Request){
    
    r.ParseForm()
    item := r.Form
    if item["Action"][0] == "terminate" {
      terminate <- true
    }
    if checkFormSyntax(item["Action"][0]) == false{
      return 
    }
    item_channel <- item
    x := <- listen_on_item_channel
    listen_on_item_channel <- x + 1

}
// will be converted to a supervisor process later on
// should also think if spawning process for each request takes more time
// than passing the request via channel to the supervisor 


func checkFormSyntax(s string) bool{
	if s == "INIT" || s == "ADD" || s == "DELETE" || s == "READ" {
		return true
	}
	return false
} 