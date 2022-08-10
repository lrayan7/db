package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)
var queue_lock sync.Mutex
var log_lock sync.Mutex

var mutex sync.Mutex 
var logfile_bytime map[time.Time]string = make(map[time.Time]string)
var logfile_byorder []time.Time 

func take_lock(s string, wg *sync.WaitGroup){
	fmt.Println("thread #" ,syscall.Getpid() ," is taking lock")
	queue_lock.Lock()
	t := time.Now()
	logfile_bytime[t] = "COMMITING: " + s
	logfile_byorder = append(logfile_byorder, t)
}

func give_lock(wg *sync.WaitGroup){
	queue_lock.Unlock()
}

func take_channel_lock(){
	queue_lock.Lock()
}

func give_channel_lock(){
	queue_lock.Lock()
}