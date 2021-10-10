package list

type MyLinkedList struct {
	len  int
	node *Node
}

type Node struct {
	val int
	next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		len: 0,
		node: &Node{val: -1},
	}
}

func (this *MyLinkedList) Len() int {
	return this.len
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Len() {
		return -1
	}
	target := this.node
	for i := 0; i < index; i ++ {
		target = target.next
	}
	return target.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	if this.Len() == 0 {
		this.len = 1
		this.node.val = val
		return
	}

	this.node = &Node {
		val: val,
		next: this.node,
	}
	this.len = this.Len() + 1
}

func (this *MyLinkedList) AddAtTail(val int)  {
	if this.Len() == 0 {
		this.node.val = val
		this.len = 1
		return
	}
	tailNode := this.node
	for tailNode.next != nil {
		tailNode = tailNode.next
	}
	tailNode.next = &Node {
		val: val,
		next: nil,
	}
	this.len = this.Len() + 1
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.Len() {
		return
	}

	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.Len() {
		this.AddAtTail(val)
		return
	}

	target := this.node
	for i := 0; i < index - 1; i ++ {
		target = target.next
	}
	newNode := &Node {
		val: val,
		next: target.next,
	}
	target.next = newNode
	this.len = this.Len() + 1
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.Len() {
		return
	}

	if index == 0 {
		this.node = this.node.next
		if this != nil {
			this.len -= 1
		}
		return
	}

	target := this.node
	for i := 0; i < index - 1; i ++ {
		target = target.next
	}
	target.next = target.next.next
	this.len -= 1
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
