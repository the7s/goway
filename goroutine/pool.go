package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

var wg sync.WaitGroup

func PoolMain() {

	jobChan := make(chan *Job)
	resultChan := make(chan *Result)
	createPool(64, jobChan, resultChan)

	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id: %v , randnum: %v , result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}

	}(resultChan)

	var id int
	for id < 10000 {
		wg.Add(1)
		rNum := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: rNum,
		}
		jobChan <- job
		id++
	}
	wg.Wait()
}

// createPool 创建线程池
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据线程个数,运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {

				rNum := job.RandNum
				var sum int
				for rNum != 0 {
					tmp := rNum % 10
					sum += tmp
					rNum /= 10
				}

				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
				wg.Done()
			}
		}(jobChan, resultChan)
	}
}
