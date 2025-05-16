package node

type Status struct {
	Running   bool
	LastError error
	LastTask  string
	Processed uint32
}
