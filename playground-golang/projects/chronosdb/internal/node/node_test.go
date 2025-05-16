package node

import (
	"context"
	"sync"
	"testing"

	"github.com/lennardclaproth/chronosdb/internal/engine/page"
)

func TestNewNode(t *testing.T) {
	// Arrange

	// Act
	n := NewNode("test-node")

	// Assert
	if n.Name != "test-node" {
		t.Errorf("expected node name 'test-node', got '%s'", n.Name)
	}

	if n.currPage == nil {
		t.Error("expected currPage to be initialized")
	}

	if n.insertWorker == nil {
		t.Error("expected insertWorker to be initialized")
	}

	if n.flushWorker == nil {
		t.Error("expected flushWorker to be initialized")
	}
}

func TestNodeInsertLotsOfRecords_WithWaitGroup(t *testing.T) {
	n := NewNode("stress-node")

	var wg sync.WaitGroup
	totalRecords := 10000
	wg.Add(totalRecords)

	// Override HandleFunc to include WaitGroup tracking
	n.insertWorker.handleFunc = func(record []byte) error {
		defer wg.Done()
		err := n.currPage.InsertRecord(record)
		if err != nil {
			n.flushWorker.Enqueue(n.currPage)
			n.currPage = page.Create(*n.currPage)
			_ = n.currPage.InsertRecord(record)
		}
		return nil
	}

	n.insertWorker.Start(context.Background()) // Restart worker with new HandleFunc

	// Produce records
	for i := 0; i < totalRecords; i++ {
		n.InsertRecord([]byte("test record"))
	}

	// Wait until all records are processed
	wg.Wait()

	if n.insertWorker.Status.Processed < uint32(totalRecords) {
		t.Errorf("expected %d records, got %d", totalRecords, n.insertWorker.Status.Processed)
	}
}
