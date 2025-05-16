package controller

import ("github.com/lennardclaproth/chronosdb/internal/engine/node")

type Controller interface {
	NewNode(name string) (*node.Node, error)
	LoadNode(name string) (*node.Node, error)
	InsertRecord(nodeName string, record []byte) error
	GetWorkerStatuses() []*node.Status
	Shutdown()
	Recover()
}