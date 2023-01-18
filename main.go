package main

import (
	"github.com/panpito/timber/timber"
	"github.com/panpito/timber/timber/example"
	"github.com/panpito/timber/timber/interface"
)

func main() {
	tree := timber_example.Branch{
		Label: "branch",
		Items: []timber_interface.NodeTree{
			timber_example.Branch{
				Label: "branch",
				Items: []timber_interface.NodeTree{
					timber_example.Leaf{Label: "leaf"},
					timber_example.Leaf{Label: "leaf"},
					timber_example.Leaf{Label: "leaf"},
				},
			},
			timber_example.Leaf{Label: "leaf"},
		},
	}

	printer := timber.NewDefaultPrinter()

	printer.Print(tree)
}
