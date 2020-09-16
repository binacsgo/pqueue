package pqueue

import (
	"container/list"
	"sync"
)

type Keytype interface{}

type Valuetype interface {
	KeyEqual(interface{}) bool
}

type PQueue struct {
	size   int64
	kvlist *list.List
	mp     map[Keytype]*list.Element

	mtx sync.RWMutex
}

func NewPQueue() *PQueue {
	return &PQueue{
		size:   0,
		kvlist: list.New(),
		mp:     make(map[Keytype]*list.Element),
	}
}

func (pq *PQueue) Size() int64 {
	pq.mtx.RLock()
	defer pq.mtx.RUnlock()
	return pq.size
}

func (pq *PQueue) Get(key Keytype) *list.Element {
	pq.mtx.RLock()
	defer pq.mtx.RUnlock()
	value, ok := pq.mp[key]
	if ok {
		return value
	}
	return nil
}

func (pq *PQueue) SetMP(key Keytype, value Valuetype) bool {
	ele := pq.Get(key)
	if ele != nil {
		ele.Value = value
		pq.kvlist.MoveToBack(ele)
		pq.mp[key] = pq.kvlist.Back()
		return true
	}
	pq.kvlist.PushBack(value)
	pq.mp[key] = pq.kvlist.Back()
	pq.size++
	return false
}

func (pq *PQueue) Set(key Keytype, value Valuetype) bool {
	pq.mtx.Lock()
	defer pq.mtx.Unlock()
	for i := pq.kvlist.Front(); i != nil; i = i.Next() {
		v := i.Value.(Valuetype)
		// TODO more definition about KeyEqual
		if value.KeyEqual(v) {
			i.Value = value
			pq.kvlist.MoveToBack(i)
			pq.mp[key] = pq.kvlist.Back()
			return true
		}
	}
	pq.kvlist.PushBack(value)
	pq.mp[key] = pq.kvlist.Back()
	pq.size++
	return false
}

func (pq *PQueue) GetMin() Valuetype {
	pq.mtx.RLock()
	defer pq.mtx.RUnlock()
	front := pq.kvlist.Front()
	if front != nil {
		return front.Value.(Valuetype)
	}
	return nil
}

func (pq *PQueue) DelMin() Valuetype {
	pq.mtx.Lock()
	defer pq.mtx.Unlock()
	front := pq.kvlist.Front()
	if front != nil {
		pq.size--
		return pq.kvlist.Remove(front).(Valuetype)
	}
	return nil
}
