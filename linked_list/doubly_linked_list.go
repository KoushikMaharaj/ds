package linked_list

import "fmt"

type DoublyLinkedList struct {
	head *dnode
}

// NewDoublyLinkedList creates a new DoublyLinkedList with head having data and next as nil
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Add adds a node to the linked list at last
func (dll *DoublyLinkedList) Add(data any) bool {
	newNode := &dnode{data, nil, nil}
	trav := dll.head
	if dll.head == nil {
		dll.head = newNode
		return true
	}
	for trav.next != nil {
		trav = trav.next
	}
	trav.next = newNode
	newNode.prev = trav
	return true
}

// Display prints contents of the linked list
func (dll *DoublyLinkedList) Display() {
	if dll.head == nil {
		fmt.Println("Doubly linked list is empty, please add at least one node")
		return
	}
	trav := dll.head
	for trav != nil {
		fmt.Println(trav.data)
		trav = trav.next
	}
}
