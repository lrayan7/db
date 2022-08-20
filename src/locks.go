package main

import (
	"sync"
	"time"
)
var log_lock sync.Mutex

var mutex sync.Mutex 
var logfile_bytime map[time.Time]string = make(map[time.Time]string)
var logfile_byorder []time.Time 

// func take_lock(s string, wg *sync.WaitGroup){
// 	fmt.Println("thread #" ,syscall.Getpid() ," is taking lock")
// 	queue_lock.Lock()
// 	t := time.Now()
// 	logfile_bytime[t] = "COMMITING: " + s
// 	logfile_byorder = append(logfile_byorder, t)
// }

func flush_lock(wg *sync.WaitGroup) {
	log_lock.Lock()
}
func flush_unlock(wg *sync.WaitGroup) {
	log_lock.Unlock()
}