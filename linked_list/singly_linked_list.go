package linked_list

import "fmt"

type SinglyLinkedList struct {
	head *snode
}

// NewSinglyLinkedList creates a new SinglyLinkedList with head having data and next as nil
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Add adds a node to the linked list at last
func (sll *SinglyLinkedList) Add(data any) bool {
	newNode := &snode{data, nil}
	trav := sll.head
	if sll.head == nil {
		sll.head = newNode
		return true
	}
	for trav.next != nil {
		trav = trav.next
	}
	trav.next = newNode
	return true
}

// Display prints contents of the linked list
func (sll *SinglyLinkedList) Display() {
	if sll.head == nil {
		fmt.Println("singly linked list is empty, please add at least one node")
		return
	}
	trav := sll.head
	for trav != nil {
		fmt.Println(trav.data)
		trav = trav.next
	}
}
