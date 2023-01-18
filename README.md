# Timber

## About

Printing an ASCII tree with Go (education purpose).

Thanks to https://github.com/Tufin/asciitree for the idea.

Example of ASCII tree:

```shell
├ root1
│ ├ sibling1
│ └ sibling2
│   ├ sibling1
│   └ sibling2
└ root2
  ├ sibling1
  └ sibling2
```

## Usage

```go
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
```
will output
```shell
2023/01/18 17:35:19 └ branch
2023/01/18 17:35:19   ├ branch
2023/01/18 17:35:19   │ ├ leaf
2023/01/18 17:35:19   │ ├ leaf
2023/01/18 17:35:19   │ └ leaf
2023/01/18 17:35:19   └ leaf 
```

If a custom printer function or result aggregator function is needed, use:
```go
func NewCustomPrinter(printerFn PrinterFn, result interface{}, resultFn ResultFn) *customPrinter
```

## Installation

### Prerequisite

```shell
$ go version
go version go1.19.3 windows/amd64
```

### Command

`go get https://github.com/panpito/timber`

## Contributing

Please free to fork, create a branch and open a pull request.

## License

This is under MIT license.

## Contact

Please contact:
[Twitter](https://twitter.com/Panpit0)
