package array

import (
	"log"
	"testing"
)

var arr = NewArray()

func TestAdd(t *testing.T) {
	arr.Add(1, 2, 3, 4, 5)
	if arr.Count() != 5 {
		t.Fatal("arr count should equal 5")
	}
	arr.ForEach(func(e interface{}) { e = e.(int) * 2; log.Println(e) })
}

func TestRemove(t *testing.T) {
	arr.Add(1, 2, 3, 4, 5)
	arr.Remove(2)
	arr.Remove(1)
	if arr.Count() != 3 {
		t.Fatal("arr count should equal 3")
	}
	arr.ForEach(func(e interface{}) { log.Println(e) })
}

func TestClear(t *testing.T) {
	arr.Add(1, 2, 3, 4, 5)
	arr.Clear()
	if arr.Count() != 0 {
		t.Fatal("arr count should equal 0")
	}
	arr.ForEach(func(e interface{}) { log.Println(e) })
}

func TestSort(t *testing.T) {
	arr.Add(5, 3, 1, 2, 4)
	arr2 := NewArray()
	arr2.Add(1, 2, 3, 4, 5)
	arr.Sort(func(i, j int) bool {
		return arr.Data()[i].(int) < arr.Data()[j].(int)
	})
	d1 := arr.Data()
	d2 := arr2.Data()
	for i := 0; i < 5; i++ {
		if d1[i] != d2[i] {
			t.FailNow()
		}
	}
}

func TestFilter(t *testing.T) {
	arr.Add(1, 2, 3, 4, 5)
	a2 := arr.Filter(func(e interface{}) bool {
		return e.(int) > 3
	}).Filter(func(e interface{}) bool {
		return e.(int) > 4
	})
	if a2.data[0] != 5 {
		t.FailNow()
	}
}
