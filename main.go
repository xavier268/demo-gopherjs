package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

func main() {

	fmt.Printf("js.Global is %s \n", js.Global.String()) // This gets printed on the browser console, or on the node output.

	switch js.Global.String() {

	case "[object global]": // node
		fmt.Println("Running node")
		fmt.Println("Available Global keys : ", js.Keys(js.Global))
		fmt.Println(`js.Global.Get("global")`, js.Global.Get("global"))
		fmt.Println(`js.Global.Get("window")`, js.Global.Get("window"))

		// Demonstrate a cleaner way to test if running node or browser ...
		fmt.Println("Testing for window key in js.Global : ", js.Global.Get("window") != js.Undefined) // false
		fmt.Println("Testing for global key in js.Global : ", js.Global.Get("global") != js.Undefined) // true

	case "[object Window]": // browser
		fmt.Println("Running on a browser")
		fmt.Println("Available Global keys : ", js.Keys(js.Global))
		fmt.Println(`js.Global.Get("global")`, js.Global.Get("global"))
		fmt.Println(`js.Global.Get("window")`, js.Global.Get("window"))

		// Demonstrate a cleaner way to test if running node or browser ...
		fmt.Println("Testing for window key in js.Global : ", js.Global.Get("window") != js.Undefined) // true
		fmt.Println("Testing for global key in js.Global : ", js.Global.Get("global") != js.Undefined) // false

		doc := js.Global.Get("document")

		doc.Call("writeln", "<div>Test paragraph 1</div>") // This will write directly inside the document
		doc.Call("writeln", "<div>Test paragraph 2</div>")

		fmt.Println("Nb of div : ", doc.Call("getElementsByTagName", "div").Length()) // Note the use of the Length() method, not the len(object) !

	default: // unknown runtime ?!
		fmt.Println("Unknown runtime environement, neither node nor browser ?")
		fmt.Println("Testing for window key in js.Global : ", js.Global.Get("window") != js.Undefined)
		fmt.Println("Testing for global key in js.Global : ", js.Global.Get("global") != js.Undefined)
	}

}
