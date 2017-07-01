# tree

This package allows you to turn any directory system into a tree, made of maps and interfaces, into go.
Any file type can be supported if specified.

## Installation

Assuming you have setup go correctly, you can run
```
go get github.com/conradludgate/tree
```

Currently built in supported file types are:

*	zip
*	json
*	yaml
*	toml
*	txt

## Example Usage

```
entities/
	monster/
		name.txt
		lore.txt

game.json
```

entities/monster/name.txt
```
Evil Monster
```

entities/monster/lore.txt
```
This evil mosnter once roamed the empty streets in search for lost souls
Now it just exists within a go package example
How sad is that!
```

game.json
```json
{
	"onload": "Hello World!",
	"onclose": "Goodbye :("
}
```

Given the file structure above, which can be found in the [example](/example) directory, you can use the following code

[tree_example.go](/tree_example.go)
```go
package main

import (
	"fmt"

	"github.com/conradludgate/tree"
)

func main() {
	tree.HandleJSON(nil)
	tree.HandleTXT(nil)

	data, err := tree.GenerateTree("game", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", data)
}
```

Will produce the following data structure. **I have added whitespace for easy reading**

```
map[
	entities:map[
		monster:map[
			name:Evil Monster 
			lore:This evil mosnter once roamed the empty streets in search for lost souls
				Now it just exists within a go package example
				How sad is that!
			]
		] 

	game:map[
		onload:Hello World! 
		onclose:Goodbye :(
	]
]
```

Alternatively, Marshalled into JSON format

```json
{
        "entities": {
                "monster": {
                        "lore": "This evil mosnter once roamed the empty streets in search for lost souls\nNow it just exists within a go package example\nHow sad is that!",
                        "name": "Evil Monster"
                }
        },
        "game": {
                "onclose": "Goodbye :(",
                "onload": "Hello World!"
        }
}
```

## Breaking down the example

To use tree, we must perform 2 small tasks.

1.	Specify File Handles
2.	Parse our file

### Specifying Handles

Handlers in this package work similarly to how the [net/http Handlers](https://golang.org/pkg/net/http/#Handler) work.

To specify a handle, we can simply use the [Handle](http://godoc.org/github.com/conradludgate/tree#Handle) or [HandleFunc](http://godoc.org/github.com/conradludgate/tree#HandleFunc) methods.
These take in a file extention, as a string, and a Handler.

All the built in file handles come with a simple method to enable them on the extention handler, such as `tree.HandleJSON(nil)` instead of `tree.HandleFunc(tree.JSON, tree.HandlerJSON)`

### Parsing Files

Files can be parse in 2 ways. Both work the same
The first and simplest way is to specify our file path as a string. This will use os.Open to find the *os.File

The second method uses an *os.File instead of a file path and skips the middle step above.

The [GenerateTree](http://godoc.org/github.com/conradludgate/tree#GenerateTree) function returns an interface{} and an error.

### How to use the results

The returned interface{} doesn't provide much use, so we've provided a few small tools to get through the tree.

The functions [Get](http://godoc.org/github.com/conradludgate/tree#Get) and [Insert](http://godoc.org/github.com/conradludgate/tree#Insert), along with their WithSlice equivelents use the inputs to recurse through the given tree to find or insert more interface{} values, which are still not that useful but it helps. You can use [type assertion](https://golang.org/ref/spec#Type_assertions) or [reflect](https://golang.org/pkg/reflect/) to help with using the vague interface{} types.

## Some extra notes

File extentions are handled as case insensitive but file paths should be handled as case sensitive.
Files will lose their extention when inserted into the tree.