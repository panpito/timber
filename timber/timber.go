package timber

import (
	"github.com/panpito/timber/timber/interface"
	log "github.com/sirupsen/logrus"
)

// Example of ASCII tree
//https://github.com/Tufin/asciitree

//├ root1
//│ ├ sibling1 (root1 is the parent, and root1 isn't last)
//│ └ sibling2
//│   ├ sibling1
//│   └ sibling2
//└ root2
//  ├ sibling1 (root2 is the parent, and root2 is last)
//  └ sibling2

func Print(tree timber_interface.NodeTree) {
	doPrint(tree, node{"", true})
}

func doPrint(tree timber_interface.NodeTree, node node) {
	log.Print(node.display(), tree.Display())

	length := len(tree.Components())
	for idx, child := range tree.Components() {
		isLast := idx == length-1
		doPrint(child, node.createChildNode(isLast))
	}
}

// node
type node struct {
	prefix string
	isLast bool
}

func (n node) display() string {
	if n.isLast {
		return n.prefix + "\u2514 " // └
	}
	return n.prefix + "\u251c " // ├
}

func (n node) createChildNode(isLast bool) node {
	return node{
		prefix: n.prefix + n.nextPrefix(),
		isLast: isLast,
	}
}

func (n node) nextPrefix() string {
	if n.isLast {
		return "  "
	}
	return "\u2502 " // │
}
