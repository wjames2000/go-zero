package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/wjames2000/go-zero/core/stat"
)

func main() {
	fmt.Println(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU()+10; i++ {
		go func() {
			for {
				select {
				default:
					time.Sleep(time.Microsecond)
				}
			}
		}()
	}

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			percent := stat.CpuUsage()
			fmt.Println("cpu:", percent)
		}
	}
}
