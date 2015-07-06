package main

import (
	"fmt"
	"sort"
	"sync"
)

type pointStore struct {
	pointMap map[string]*point
	lock     *sync.RWMutex
}

func newPointStore() *pointStore {
	return &pointStore{
		pointMap: make(map[string]*point),
		lock:     &sync.RWMutex{},
	}
}

func (ps *pointStore) keys() []string {
	ps.lock.Lock()
	keys := make([]string, 0)
	for k, _ := range ps.pointMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	ps.lock.Unlock()
	return keys
}

func (ps *pointStore) add(p *point) error {
	var err error
	ps.lock.Lock()
	if _, ok := ps.pointMap[p.Name]; ok {
		err = ps.pointMap[p.Name].add(p)
	} else {
		ps.pointMap[p.Name] = p
	}
	ps.lock.Unlock()
	return err
}

func (ps *pointStore) get(name string) (*point, error) {
	ps.lock.Lock()
	if p, ok := ps.pointMap[name]; ok {
		ps.lock.Unlock()
		return p, nil
	}
	return &point{}, fmt.Errorf("point does not exist")
}
