package timber

import (
	"github.com/panpito/timber/timber/interface"
	log "github.com/sirupsen/logrus"
)

func Print(tree timber_interface.NodeTree) {
	doPrint(tree, "")
}

func doPrint(tree timber_interface.NodeTree, prefix string) {
	log.Print(prefix, tree.Display())

	for _, child := range tree.Components() {
		doPrint(child, prefix+"  ")
	}
}

// Example of ASCII tree

// my-app/
//├─ node_modules/
//├─ public/
//│  ├─ favicon.ico
//│  ├─ index.html
//│  ├─ robots.txt
//├─ src/
//│  ├─ index.css
//│  ├─ index.js
//├─ .gitignore
//├─ package.json
//├─ README.md
