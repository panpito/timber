package timber_interface

type Printer interface {
	Print(tree NodeTree) (interface{}, error)
}
