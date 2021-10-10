package list

import "testing"

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	//t.Logf("Len: %v, value: %v", obj.Len(), obj.Get(0))
	//obj.AddAtTail(3)
	//t.Logf("Len: %v, value1: %v, value2: %v", obj.Len(), obj.Get(0), obj.Get(1))
	//obj.AddAtIndex(1,2)
	//t.Logf("Len: %v, value1: %v, value2: %v", obj.Len(), obj.Get(0), obj.Get(1))
	//t.Logf("Len: %v, value1: %v, value2: %v, value3: %v", obj.Len(), obj.Get(0), obj.Get(1), obj.Get(2))
	//obj.DeleteAtIndex(1)
	//t.Logf("Len: %v, value: %v", obj.Len(), obj.Get(1))
	obj.AddAtHead(2)
	t.Logf("Len: %v, v1: %v, v2: %v, v3: %v", obj.Len(), obj.Get(0),obj.Get(1), obj.Get(2))
	obj.AddAtHead(3)
	t.Logf("Len: %v, v1: %v, v2: %v, v3: %v", obj.Len(), obj.Get(0),obj.Get(1), obj.Get(2))
	//obj.DeleteAtIndex(0)
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	obj := Constructor()
	obj.AddAtTail(1)
	obj.AddAtTail(2)
	obj.AddAtTail(3)
	t.Logf("Len: %v, v1: %v, v2: %v, v3: %v", obj.Len(), obj.Get(0),obj.Get(1), obj.Get(2))
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	obj := Constructor()
	obj.AddAtIndex(0, 1)
	obj.AddAtIndex(0, 2)
	obj.AddAtIndex(0, 3)
	t.Logf("Len: %v, v1: %v, v2: %v, v3: %v", obj.Len(), obj.Get(0),obj.Get(1), obj.Get(2))

	obj.AddAtIndex(3, 4)
	obj.AddAtIndex(3, 5)
	t.Logf("Len: %v, v4: %v, v5: %v", obj.Len(), obj.Get(3),obj.Get(4))
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	obj := Constructor()
	obj.AddAtIndex(0, 1)
	obj.AddAtIndex(0, 2)
	obj.AddAtIndex(0, 3)
	obj.AddAtIndex(3, 4)
	obj.AddAtIndex(3, 5)
	// 3->2->1->5->4
	obj.DeleteAtIndex(1)
	obj.DeleteAtIndex(3)
	t.Logf("Len: %v, v1: %v, v2: %v, v3: %v", obj.Len(), obj.Get(0),obj.Get(1), obj.Get(2))
}
