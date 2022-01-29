package main

import (
	"fmt"

	"github.com/renatobrittoaraujo/fullcycle-go-video-encoder/async"
)

func main() {
	a := make(chan async.JobRet)
	b := make(chan async.JobRet)
	go async.LongRoutine("12708317289032178903", a)
	go async.LongRoutine("123790-1738921381209", b)
	var1 := <-a
	var2 := <-b
	fmt.Println("Got", var1, var2)
}
