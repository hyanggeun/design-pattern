package main

import (
	"example.com/pkg/singleton/instanceWithInit"
	"example.com/pkg/singleton/instanceWithSyncOnce"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	//100 goroutine
	for i:=0 ;i <100 ;i ++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup){
			defer wg.Done()
			instance := instanceWithSyncOnce.GetInstance()
			instance2 := instanceWithInit.GetInstance()
			instance.Add()
			instance2.Add()
		}(&wg)
	}
	wg.Wait()
	instanceWithSyncOnce.GetInstance().PrintNum()
	instanceWithInit.GetInstance().PrintNum()
}
