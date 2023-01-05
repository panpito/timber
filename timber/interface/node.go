package timber_interface

type Node interface {
	Display() string
}

type NodeTree interface {
	Node // interface composition in Go!
	Components() []NodeTree
}
