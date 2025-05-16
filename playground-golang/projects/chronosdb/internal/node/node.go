package node

import (
	"context"
	"log"

	"github.com/lennardclaproth/chronosdb/internal/engine/page"
	"github.com/lennardclaproth/chronosdb/internal/engine/writer"
)

type Node struct {
	Name string
	// Index Index
	// Header Header
	currPage     *page.Page
	head         *page.Page
	insertWorker *Worker[[]byte]
	flushWorker  *Worker[*page.Page]
	fileWriter   writer.Writer
}

func NewNode(name string) *Node {
	// Initialize node
	h := page.CreateHead()
	node := &Node{
		Name:     name,
		head:     h,
		currPage: h,
	}

	fw, err := writer.NewFileWriter(node.Name)

	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
	
	node.fileWriter = fw

	// Setup workers and worker functions
	node.insertWorker = NewWorker(1000, func(record []byte) error {
		err := node.currPage.InsertRecord(record)
		if err != nil {
			node.flushWorker.Enqueue(node.currPage)
			node.currPage = page.Create(*node.currPage)
			_ = node.currPage.InsertRecord(record)
		}
		return nil
	})

	node.flushWorker = NewWorker(1000, func(page *page.Page) error {
		// should convert the page to bytes and push it to the writer
		node.fileWriter.WritePage(page)
		return nil
	})

	return node
}

func LoadNode(nodeName string) (*Node, error) {
	// Loads a node

	return &Node{}, nil
}

func (n *Node) Start(ctx context.Context) {
	n.insertWorker.Start(ctx)
}

func (n *Node) InsertRecord(record []byte) {
	n.insertWorker.Enqueue(record)
}
