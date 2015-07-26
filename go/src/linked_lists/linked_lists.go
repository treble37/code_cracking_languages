package linked_lists

import (
	"container/list"
	"fmt"
)

func nthToLastElement(i int, emap []int) int {
  m := len(emap)
	return emap[m-i-1]
}

func removeDuplicates(l *list.List) *list.List {
	var newList = list.New()
	var n map[int]bool = make(map[int]bool)
	for e := l.Front(); e != nil; e = e.Next() {
		n[e.Value.(int)] = true
	}
	for k, _ := range n {
		newList.PushBack(k)
	}
	return newList
}

func equalValuesInLists(list1 *list.List, list2 *list.List) bool {
	var flag bool = true
	for e := list1.Front(); e != nil; e = e.Next() {
		flag = flag && elemInList(e.Value.(int), list2)
	}
	return flag
}

func elemInList(n int, l *list.List) bool {
	var m map[int]bool = make(map[int]bool)
	var flag bool = false
	for e := l.Front(); e != nil; e = e.Next() {
		m[e.Value.(int)] = true
	}
	if m[n] == true {
		flag = true
	}
	return flag
}

func printListElements(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
