package main

import (
	"fmt"
	"net/http"
	"net/url"
	// "os"
	"sync"
)

var db_msg string 
var wg sync.WaitGroup
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
    checkFormSyntax(item["Action"][0])
    
    go func(){
      spawnner_wrapper(item)
    }()
}
// will be converted to a supervisor process later on
// should also think if spawning process for each request takes more time
// than passing the request via channel to the supervisor 
func spawnner_wrapper(message url.Values){
    thread_response := make(chan string)    
    wg.Add(1)
    go func(){
      spawn_thread(thread_response, message)
      wg.Done() // not sure still 
    }()
    // wait for child to finish
    wg.Wait()
    response_msg := <-thread_response
    fmt.Println(response_msg, " !")    
}

func checkFormSyntax(s string) bool{
	if s == "INIT" || s == "ADD" || s == "DELETE" || s == "READ" {
		return true
	}
	return false
} 