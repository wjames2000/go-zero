package main

import (
	"fmt"
	"time"

	"github.com/wjames2000/go-zero/core/executors"
)

func main() {
	executor := executors.NewBulkExecutor(func(items []interface{}) {
		fmt.Println(len(items))
	}, executors.WithBulkTasks(10))
	for {
		executor.Add(1)
		time.Sleep(time.Millisecond * 90)
	}
}
