package datastruct

import (
	"fmt"

	"github.com/lennardclaproth/golang-playground/algorithms/crypto"
)

type HashTable[T comparable] struct {
	buckets []*SinglyLinkedList[T]
}

func NewHashTable[T comparable](size int) (*HashTable[T], error) {
	if size <= 0 {
		return nil, fmt.Errorf("cannot create a hashtable of a size smaller or equal to 0")
	}

	return &HashTable[T]{
		buckets: make([]*SinglyLinkedList[T], 0, size),
	}, nil
}

// Insert adds an item to the hashtable. To prevent collisions each bucket contains
// a singly linked list. When a collision is detected a new node will be added
// to the tail of the singly linked list.
func (h *HashTable[T]) Insert(i T) error {
	// Get the index
	index := h.getIndex(i)
	// Get the head of the bucket at the index
	head := h.buckets[index]
	// If the head is nil it means we do not have a collision, instantiate a new singly linked list
	if head == nil {
		h.buckets[index] = NewSinglyLinkedList(i)
		return nil
	}
	// Set the current to the head of the singly linked list in the bucket
	curr := head
	for curr != nil {
		if curr.data == i {
			return fmt.Errorf("duplicate entry: %v", i)
		}
		// When next is nil it means we reached the end, we can break the loop
		if curr.next == nil {
			break
		}
		// Set curr to the next value in the singly linked list.
		curr = curr.next
	}
	// Insert a new singlyLinkedList (this is actually a node)
	curr.Insert(NewSinglyLinkedList(i))
	return nil
}

// getIndex gets the index based of the data of the item. It does this
// by building a hash of the bytes of the item and creating a hash of off those bytes.
// With this hash we can determine the index by using the modulo operator and the value
// of the hash i.e. 5 % 2 = 1 so index 1 is used.
func (h *HashTable[T]) getIndex(i T) int {
	// Create bytes of the value passed into the insert function
	bytes := []byte(fmt.Sprintf("%v", i))
	// Create a hash string based on the bytes
	hash := crypto.HashBytes(bytes)
	// Determine the index based on the hash mod len(buckets) which causes the index to wrap around.
	return int(hash % uint64(len(h.buckets)))
}

func (h *HashTable[T]) Remove(i T) error {
	index := h.getIndex(i)
	head := h.buckets[index]

	if head == nil {
		return fmt.Errorf("data not found in hash table")
	}

	if head.data == i {
		h.buckets[index] = head.next
		return nil
	}

	curr := head
	for curr.next != nil {
		if curr.next.data == i {
			curr.next = curr.next.next
			return nil
		}
		curr = curr.next
	}

	return fmt.Errorf("data not found in hash table")
}
