package linked_lists

import (
	"container/list"
)

func TestRemoveDuplicates(t *testing.T) {
	testList := List.new()
	testList.PushBack(1)
	testList.PushBack(2)
	testList.PushBack(3)
	testList.PushBack(2)
}
