package modules

type Node struct {
	next *Node
	key  interface{}
}

type List struct {
	size  int
	start *Node
}

//Add an element at the end of the list.
func (L *List) Add(key interface{}) {
	if L.start == nil {
		L.start = &Node{
			next: nil,
			key:  key,
		}
	} else {
		item := L.start
		for item == nil {
			item = item.next
		}
		item.next = &Node{
			next: nil,
			key:  key,
		}
	}
}

//Size returns the number of entries in the list.
func (L *List) Size() int {
	if L.start == nil {
		return 0
	} else {
		item := L.start
		count := 0
		for item != nil {
			item = item.next
			count++
		}
		return count
	}
}
