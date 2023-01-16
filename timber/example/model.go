package timber_example

import "github.com/panpito/timber/timber/interface"

type Leaf struct {
	Label string
}

func (l Leaf) Components() []timber_interface.NodeTree {
	return nil
}

func (l Leaf) Display() string {
	return l.Label
}

type Branch struct {
	Label string
	Items []timber_interface.NodeTree
}

func (b Branch) Display() string {
	return b.Label
}

func (b Branch) Components() []timber_interface.NodeTree {
	return b.Items
}
