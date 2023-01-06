package main

import (
	"github.com/panpito/timber/timber"
	timber_interface "github.com/panpito/timber/timber/interface"
)

func main() {
	l0 := Leaf{label: "l0"}
	l1 := Leaf{label: "l1"}

	b1 := Branch{
		label: "branch1",
		components: []timber_interface.NodeTree{
			Leaf{label: "l3"},
		},
	}

	b3 := Branch{
		label: "branch3",
		components: []timber_interface.NodeTree{
			Leaf{label: "l7"},
		},
	}

	b2 := Branch{
		label: "branch2",
		components: []timber_interface.NodeTree{
			Leaf{label: "l4"},
			Leaf{label: "l5"},
			Leaf{label: "l6"},
			b3,
		},
	}

	b0 := Branch{
		label: "branch0",
		components: []timber_interface.NodeTree{
			l0, l1, b1, b2, Leaf{label: "l8"},
		},
	}

	timber.Print(b0)
}

// For example
type Leaf struct {
	label string
}

func (l Leaf) Components() []timber_interface.NodeTree {
	return nil
}

func (l Leaf) Display() string {
	return l.label
}

type Branch struct {
	label      string
	components []timber_interface.NodeTree
}

func (b Branch) Display() string {
	return b.label
}

func (b Branch) Components() []timber_interface.NodeTree {
	return b.components
}
