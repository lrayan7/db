package main

import (
	"net/http"
)

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
    
    // item is valid at this point
    go func(){ 
      listen_on_item_channel <- "new" 
      item_channel <- item
    }()
}