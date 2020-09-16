package main

import (
	"log"
	"time"

	"github.com/binacsgo/pqueue"
)

// -------------------- Value def. --------------------

type Value struct {
	Key string
	t   int64
	v1  int64
	v2  int64
	v3  string
}

func (value1 *Value) KeyEqual(value interface{}) bool {
	value2 := value.(*Value)
	return value1.Key == value2.Key
}

// -------------------- something impl. --------------------

type SthImpl struct {
	pq *pqueue.PQueue
}

func (impl *SthImpl) AddValue(key string, v1 int64) {
	log.Printf("before add pqueue-size %+v\n", impl.pq.Size())
	value := &Value{
		Key: key,
		t:   time.Now().Unix(),
		v1:  1,
		v2:  2,
		v3:  "3",
	}
	update := impl.pq.Set(value.Key, value)
	log.Printf("after add pqueue-size %+v, update %+v, key %+v\n", impl.pq.Size(), update, value.Key)
}

func (impl *SthImpl) DelValue() {
	log.Printf("before delete pqueue-size %+v\n", impl.pq.Size())
	// WARN: DelMin() maybe return `nil`, you must do with this situation yourself
	// such as:
	//	min := impl.pq.DelMin()
	//	if min != nil {
	//		do sth.
	//	}
	//	do sth. else
	//
	value := impl.pq.DelMin().(*Value)
	log.Printf("after delete pqueue-size %+v, key %+v\n", impl.pq.Size(), value.Key)
}

func main() {
	impl := SthImpl{pq: pqueue.NewPQueue()}
	impl.AddValue("key1", 1) // now ["key1"]
	impl.AddValue("key2", 1) // now ["key1","key2"]
	impl.AddValue("key3", 1) // now ["key1","key2","key3"]
	impl.AddValue("key1", 1) // now ["key2","key3","key1"]

	// should be "key2", "key3", "key1"
	impl.DelValue()
	impl.DelValue()
	impl.DelValue()
}
