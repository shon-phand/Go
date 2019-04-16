package main

import(
	"fmt"
	"time"
	"math/rand"
	"sync"
)

type Sample struct{
		ID int
		NA int
		ZN int
		CL int
}

type Result struct{
	sample Sample
	result int
}

var jobs=make(chan Sample)
var results=make(chan Result)


func allocate(no_of_sample int){
	for i:=0;i<no_of_sample;i++{
		na:=rand.Intn(3)
		zn:=rand.Intn(3)
		cl:=rand.Intn(3)
		job:=Sample{i,na,zn,cl}
		jobs<-job
	}
	close(jobs)
}

func result(done chan []Result){
	testResults:=[]Result{}
	for result:=range results{
		//fmt.Println("in a result",result)
		
		testResults=append(testResults,result)
		//fmt.Print(testResults)
	}
	done <- testResults
}

func createWorkerPool(no_of_worker int){
	var wg sync.WaitGroup
	for i:=0;i<no_of_worker;i++{
		wg.Add(1)
		go worker(&wg)
		
	}
	wg.Wait()
	close(results)

} 

func worker(wg *sync.WaitGroup){
	for job:=range jobs{
		output:= Result{job, calculation(job)}
		results<-output
	}
	wg.Done()
}

func calculation(job Sample)int{
	na:=job.NA
	zn:=job.ZN
	cl:=job.CL

	op:= 2*na+3*zn+2*cl
	return op
}


func main(){

	testResults:=[]Result{}
	startTime:=time.Now()
	no_of_sample:=10
	go allocate(no_of_sample)
	done:=make(chan []Result)
	go result(done)
	no_of_worker:=10
	createWorkerPool(no_of_worker)
	testResults= <-done
	fmt.Println(testResults)
	endTime:=time.Now()
	diff:=endTime.Sub(startTime)
	fmt.Println("total time taken : ",diff)	

}