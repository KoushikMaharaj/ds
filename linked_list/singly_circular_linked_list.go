package linked_list

import "fmt"

type SinglyCircularLinkedList struct {
	head *snode
}

// NewSinglyCircularLinkedList creates a new SinglyLinkedList with head having data and next as nil
func NewSinglyCircularLinkedList() *SinglyCircularLinkedList {
	return &SinglyCircularLinkedList{}
}

func (csll *SinglyCircularLinkedList) Add(data any) bool {
	newNode := &snode{data: data, next: nil}
	if csll.head == nil {
		csll.head = newNode
		newNode.next = csll.head
		return true
	}
	trav := csll.head
	for trav.next != csll.head {
		trav = trav.next
	}
	trav.next = newNode
	newNode.next = csll.head
	return true
}

func (csll *SinglyCircularLinkedList) Display() {
	if csll.head == nil {
		fmt.Println("Doubly linked list is empty, please add at least one node")
		return
	}
	trav := csll.head
	for {
		fmt.Println(trav.data)
		trav = trav.next
		if trav == csll.head{
			break
		}
	}
}
