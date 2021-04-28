package instanceWithSyncOnce

import (
	"fmt"
	"sync"
)
var instance *Instance
var once sync.Once

type Instance struct {
	mu sync.Mutex
	num int
}

func GetInstance() *Instance {
	once.Do(func(){
		instance = &Instance{num: 0}
	})
	return instance
}

func (i *Instance) Add() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.num++
}

func (i *Instance) PrintNum() {
	fmt.Println(i.num)
}
