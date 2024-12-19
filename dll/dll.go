package dll

import (
	"fmt"
)

func DllInit() {
	n1 := Node{
		value: 8,
		next:  nil,
		prev:  nil,
	}
	n2 := Node{
		value: 10,
		next:  nil,
		prev:  &n1,
	}
	n3 := Node{
		value: 12,
		next:  nil,
		prev:  &n2,
	}
	n1.next = &n2
	n2.next = &n3

	fmt.Print(n1.value)
	fmt.Print("->")
	fmt.Print(n1.next.value)
	fmt.Print("->")
	fmt.Print(n1.next.next.value)
	fmt.Print("<-")
	nx := n1.next.next
	fmt.Print(nx.prev.value)
	fmt.Print("<-")
	fmt.Print(nx.prev.prev.value)
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}
