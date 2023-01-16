package timber

import (
	"fmt"
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

type PrinterFn func(display string) (interface{}, error)
type ResultFn func(previousResult, currentItem interface{}) (interface{}, error)

var DefaultPrinterFn = func(display string) (interface{}, error) {
	log.Print(display)
	return display, nil
}

var DefaultResultFn = func(previousResult, currentItem interface{}) (interface{}, error) {
	return append(previousResult.([]string), currentItem.(string)), nil
}

type customPrinter struct {
	printerFn PrinterFn
	result    interface{}
	resultFn  ResultFn
}

func NewCustomPrinter(printerFn PrinterFn, result interface{}, resultFn ResultFn) *customPrinter {
	return &customPrinter{printerFn: printerFn, result: result, resultFn: resultFn}
}

func NewDefaultPrinter() *customPrinter {
	return &customPrinter{printerFn: DefaultPrinterFn, result: make([]string, 0), resultFn: DefaultResultFn}
}

func (printer *customPrinter) Print(tree timber_interface.NodeTree) (interface{}, error) {
	if err := printer.doPrint(tree, node{"", true}); err != nil {
		return nil, err
	}

	return printer.result, nil
}

func (printer *customPrinter) doPrint(tree timber_interface.NodeTree, node node) error {
	display := fmt.Sprint(node.display(), tree.Display())
	displayResult, errPrint := printer.printerFn(display)
	if errPrint != nil {
		return errPrint
	}

	updatedResult, errResult := printer.resultFn(printer.result, displayResult)
	if errResult != nil {
		return errResult
	}
	printer.result = updatedResult

	length := len(tree.Components())
	for idx, child := range tree.Components() {
		isLast := idx == length-1
		printer.doPrint(child, node.createChildNode(isLast))
	}

	return nil
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
