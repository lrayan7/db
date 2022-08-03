package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
type teststruct struct{
  Action string `json:"Action"`
  Table string `json:"Table"`
  X map[string]interface{} `json:"Value"`
}  
func dbReqHandler(w http.ResponseWriter, r *http.Request){
    request_action, _ := ioutil.ReadAll(r.Body)
//--------------------------------------------------------------------------------------------
//              JSON handling - can be cleaner for sure 
//--------------------------------------------------------------------------------------------
    actionJsonFormat := string(request_action)
    actionJsonFormat = strings.ReplaceAll(actionJsonFormat, "=", `":"`)
    actionJsonFormat = strings.ReplaceAll(actionJsonFormat, "&", `","`)
    actionJsonFormat = strings.ReplaceAll(actionJsonFormat, "+", `:`)
    for i,_ := range strings.Split(actionJsonFormat, ","){
        if i >= 2{
          str2 := "x" + strconv.Itoa(i-1)
          actionJsonFormat = strings.Replace(actionJsonFormat, "-", str2,1)
        }
    }
    actionJsonFormat = `{"` + actionJsonFormat + `"}`
//--------------------------------------------------------------------------------------------
/*
    here unmarshal incoming json to struct where we have 2 const fields and more 
    dynamical fields depending on client input.
*/
    newMsg := msg{}
    if err := json.Unmarshal([]byte(actionJsonFormat), &newMsg); err != nil{
      fmt.Println("Error !")
    }
    if err := json.Unmarshal([]byte(actionJsonFormat), &newMsg.Value); err != nil{
      fmt.Println("Error !")
    }
    delete(newMsg.Value, "Action")
    delete(newMsg.Value, "Table")
    // for _,val := range newMsg.Value{
    //   fmt.Println("-> ", val)
    // }
    
    go func(){
      wg.Add(1)
      spawnner_wrapper(newMsg)
    }()
}
func spawnner_wrapper(message msg){
    thread_response := make(chan string)    
    go func(){
      spawn_thread(thread_response, message)
      wg.Done()
    }()
    responded := false
    for {
      select {
      case <-thread_response:
        responded = true
        break
      }
      if(responded){ break }
    }
     
    response_msg := <-thread_response
    fmt.Println("answer is= ", response_msg, " !")    
}