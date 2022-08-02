package main

import (
	// "fmt"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
	"sync"
	// "time"
	// "encoding/json"
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

    request_action, _ := ioutil.ReadAll(r.Body)
    actionJsonFormat := string(request_action)
    // fmt.Println(actionJsonFormat)
    var buffer bytes.Buffer
    buffer.WriteString(`"`)
    for _, char := range actionJsonFormat {
        if char != '&' && char != '=' {
          buffer.WriteString(string(char))
        }
        if char == '&' {
          buffer.WriteString(`","`)
        } 
        if char == '=' {
          buffer.WriteString(`":"`)
        }
    }
    buffer.WriteString(`"`)
    actionJsonFormat = "[{" + buffer.String() + "}]"
    var newMsg []msg
    err := json.Unmarshal([]byte(actionJsonFormat), &newMsg)
    // fmt.Println("now", newMsg)
    // fmt.Println("")
    // fmt.Println(actionJsonFormat)
    if err != nil {
      // panic(err)
      fmt.Println("Error !")
    }
    // cast this string which is json to this struct
    // for each http request to /dbgo, will spawn new wrapper 
    // each wrapper will spawn a child thread that the wrapper will
    // wait for, and by the child thread finishing its job, wrapper will exit 
    go func(){
      spawnner_wrapper(newMsg)
    }()
}   
func spawnner_wrapper(message []msg){
    thread_response := make(chan string)    
    wg.Add(1)
    
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
    wg.Wait() // might be unnecessary 
    response_msg := <-thread_response
    fmt.Println("answer is= ", response_msg, " !")    
}