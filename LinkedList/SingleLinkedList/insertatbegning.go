package main

import "fmt"

//Create a ListNode and LinkedList struct
// ListNode...
type ListNode struct {
	data interface{}
	next *ListNode
}

// LinkedList...
type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) insertAtStarting(data interface{}) {
	//Create a new node that can be added to the begning of the linkedlist
	newNode := &ListNode{
		data: data,
	}
	if ll.head == nil {
		//LinkedList is empty so the new node will be the first node
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.size++
	return
}

func (ll *LinkedList) display() error {
	if ll.head == nil {
		return fmt.Errorf("List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

func main() {

	ll := LinkedList{}
	ll.insertAtStarting("A")
	ll.insertAtStarting("B")
	ll.insertAtStarting("C")
	err := ll.display()
	if err != nil {
		fmt.Println(err.Error)
	}

}
