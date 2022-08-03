package main

import (
	"fmt"
	"sync"
	"time"
)


var mutex sync.Mutex 
var logfile_bytime map[time.Time]string = make(map[time.Time]string)
var logfile_byorder []time.Time 

func take_lock(s string, wg *sync.WaitGroup){
	
	mutex.Lock()
	t := time.Now()
	logfile_bytime[t] = "COMMITING: " + s
	logfile_byorder = append(logfile_byorder, t)
	// mutex.Unlock()
}

func give_lock(wg *sync.WaitGroup){
		
	mutex.Unlock()
	fmt.Println("here ")

}

