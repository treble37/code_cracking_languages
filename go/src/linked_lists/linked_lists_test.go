package linked_lists

import (
	"container/list"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var (
		testList    = list.New()
		exampleList = list.New()
		uniqueList  = list.New()
	)
	testList.PushBack(1)
	testList.PushBack(2)
	testList.PushBack(3)
	testList.PushBack(2)
	uniqueList.PushBack(1)
	uniqueList.PushBack(2)
	uniqueList.PushBack(3)
	exampleList = removeDuplicates(testList)
	if equalValuesInLists(exampleList, uniqueList) == false {
		t.Fatalf("removeDuplicates failed")
	}
}
