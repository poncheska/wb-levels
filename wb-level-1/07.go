package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись в map.

func main() {
	am := NewAtomicIntIntMap()
	wg := &sync.WaitGroup{}
	n := 100

	wg.Add(100)
	for i := 0; i < n; i++ {
		go func(v int) {
			am.Write(v, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(am.GetMap())
}

type AtomicIntIntMap struct {
	m  map[int]int
	mx sync.RWMutex
}

func NewAtomicIntIntMap() *AtomicIntIntMap {
	return &AtomicIntIntMap{
		m:  make(map[int]int),
		mx: sync.RWMutex{},
	}
}

func (am *AtomicIntIntMap) Write(key, value int) {
	am.mx.Lock()
	defer am.mx.Unlock()
	am.m[key] = value
}

func (am *AtomicIntIntMap) Delete(key int) {
	am.mx.Lock()
	defer am.mx.Unlock()
	delete(am.m, key)
}

func (am *AtomicIntIntMap) Get(key int) (int, bool) {
	am.mx.RLock()
	defer am.mx.RUnlock()
	v, ok := am.m[key]
	return v, ok
}

func (am *AtomicIntIntMap) GetMap() map[int]int {
	am.mx.RLock()
	defer am.mx.RUnlock()
	return am.m
}
