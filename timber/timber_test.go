package timber_test

import (
	"errors"
	"github.com/panpito/timber/timber"
	"github.com/panpito/timber/timber/example"
	"github.com/panpito/timber/timber/interface"
	"log"
	"testing"
)

func init() {
	log.SetFlags(0)
}

func TestSucess_complex(t *testing.T) {
	//	given
	tree := timber_example.Branch{
		Label: "branch",
		Items: []timber_interface.NodeTree{
			timber_example.Branch{
				Label: "branch2",
				Items: []timber_interface.NodeTree{
					timber_example.Leaf{Label: "leaf2"},
					timber_example.Leaf{Label: "leaf3"},
				},
			},
			timber_example.Leaf{Label: "leaf"},
		},
	}

	printer := timber.NewDefaultPrinter()

	//	when
	result, err := printer.Print(tree)

	//	then
	Equal(t, nil, err)
	castedResult := result.([]string)
	Equal(t, 5, len(castedResult))
	Equal(t, "└ branch", castedResult[0])
	Equal(t, "  ├ branch2", castedResult[1])
	Equal(t, "  │ ├ leaf2", castedResult[2])
	Equal(t, "  │ └ leaf3", castedResult[3])
	Equal(t, "  └ leaf", castedResult[4])
}

// TODO: potentiel bug to fix (should have been TDD)
func TestSucess_emptyTree(t *testing.T) {
	//	given
	tree := timber_example.Branch{}

	printer := timber.NewDefaultPrinter()

	//	when
	result, err := printer.Print(tree)

	//	then
	Equal(t, nil, err)
	castedResult := result.([]string)
	Equal(t, 1, len(castedResult))
}

func TestFailure_printFnFailure(t *testing.T) {
	//	given
	tree := timber_example.Branch{}

	faultyPrintFn := func(display string) (interface{}, error) {
		log.Print(display)
		return display, errors.New("bam")
	}
	printer := timber.NewCustomPrinter(faultyPrintFn, make([]string, 0), timber.DefaultResultFn)

	//	when
	result, err := printer.Print(tree)

	//	then
	Equal(t, "bam", err.Error())
	Equal(t, nil, result)
}

func TestFailure_resultFnFailure(t *testing.T) {
	//	given
	tree := timber_example.Branch{}

	faultyResultFn := func(previousResult, currentItem interface{}) (interface{}, error) {
		return nil, errors.New("boom")
	}
	printer := timber.NewCustomPrinter(timber.DefaultPrinterFn, make([]string, 0), faultyResultFn)

	//	when
	result, err := printer.Print(tree)

	//	then
	Equal(t, "boom", err.Error())
	Equal(t, nil, result)
}

func Equal(t *testing.T, wanted interface{}, received interface{}) {
	if wanted != received {
		t.Errorf("\nwanted: \t%v\nreceived: \t%v", wanted, received)
	}
}
