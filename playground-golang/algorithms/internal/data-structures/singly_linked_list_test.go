package datastruct

import (
	"testing"
)

// TestInsert checks if nodes are inserted correctly.
func TestInsert(t *testing.T) {
	head := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	// Insert nodes: head -> node2 -> node3
	head.Insert(node2)
	node2.Insert(node3)

	// Verify the structure
	if head.next != node2 {
		t.Errorf("Expected head.next to be node2, got %+v", head.next)
	}

	if head.next.next != node3 {
		t.Errorf("Expected head.next.next to be node3, got %+v", head.next.next)
	}

	if head.next.next.next != nil {
		t.Errorf("Expected end of list to be nil, got %+v", head.next.next.next)
	}

	if head.data != 1 || head.next.data != 2 || head.next.next.data != 3 {
		t.Errorf("Unexpected data sequence: got %v -> %v -> %v", head.data, head.next.data, head.next.next.data)
	}
}

// TestInsertIntoEmptyList checks inserting a node into a single-node list.
func TestInsertIntoEmptyList(t *testing.T) {
	head := NewNode(42)
	newNode := NewNode(99)

	head.Insert(newNode)

	if head.next != newNode {
		t.Errorf("Expected head.next to be newNode, got %+v", head.next)
	}

	if head.next.data != 99 {
		t.Errorf("Expected new node to have data 99, got %v", head.next.data)
	}
}

// TestMultipleInserts ensures that multiple insertions work correctly.
func TestMultipleInserts(t *testing.T) {
	head := NewNode(1)

	// Insert multiple nodes
	for i := 2; i <= 5; i++ {
		newNode := NewNode(i)
		head.Insert(newNode)
	}

	// Check the data of the first node's next
	if head.next.data != 5 {
		t.Errorf("Expected first inserted node's data to be 5, got %v", head.next.data)
	}

	// Check the data of the second node
	if head.next.next.data != 4 {
		t.Errorf("Expected second inserted node's data to be 4, got %v", head.next.next.data)
	}
}

// TestInsertWithDifferentTypes ensures generic types work as expected.
func TestInsertWithDifferentTypes(t *testing.T) {
	head := NewNode("first")
	newNode := NewNode("second")

	head.Insert(newNode)

	if head.next.data != "second" {
		t.Errorf("Expected head.next.data to be 'second', got %v", head.next.data)
	}
}
