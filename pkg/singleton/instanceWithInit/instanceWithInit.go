package instanceWithInit

import (
	"fmt"
	"sync"
)

type Instance struct {
	mu sync.Mutex
	num int
}
var instance *Instance

func init() {
	instance = &Instance{num: 0}
}

func GetInstance() *Instance {
	return instance
}
func(i *Instance) Add() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.num++
}

func (i *Instance) PrintNum()  {
	fmt.Println(i.num)
}