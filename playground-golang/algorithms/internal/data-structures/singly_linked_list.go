package datastruct

// A linkedList is a list whith a reference to the
// next item and its previous item. reading from the
// linkedlist has a time complexity of O(n) inserting in
// the linkedlist has a time complexity of O(1) because inserting
// does not require looping over the elements in the list.
type SinglyLinkedList[T any] struct {
	data T
	next *SinglyLinkedList[T]
}

// When inserting in a singly linked list we need to set the
// pointer of the next of the new element to the next of the
// curr element. After which we set the next of the curr element
// to the new element.
func (curr *SinglyLinkedList[T]) Insert(new *SinglyLinkedList[T]) {
	new.next = curr.next
	curr.next = new
}

// A constructor for creating a new node
func NewSinglyLinkedList[T any](data T) *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		data: data,
	}
}
