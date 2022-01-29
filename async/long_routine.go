package async

import (
	"fmt"
	"math/rand"
	"time"
)

type JobRet struct {
	id     string
	result string
}

func LongRoutine(id string, in chan JobRet) {
	fmt.Println("Go routine", id, "started")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	result := ""
	for i := 0; i < len(id); i++ {
		result = result + string((int(id[i])*10)%60+40)
	}
	jobRet := JobRet{id, result}
	fmt.Println("Go routine", id, "end with jobRet =", jobRet)
	in <- jobRet
}
