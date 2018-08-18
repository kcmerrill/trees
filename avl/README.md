# avl tree

This is a [AVL tree](https://en.wikipedia.org/wiki/AVL_tree). 

The key is an `interface{}` and uses a `LessThan` function in which you can supply two keys and figure out for yourself which is less than the other. Two types exist currently, with a few more planned. They are `avl.Strings` and `avl.Ints`. 

## Usage

```golang
package main

import (
	"github.com/kcmerrill/tree/avl"
)

func main() {
	tree := avl.New(avl.Ints)
	for x := 0; x < 100; x++ {
		tree.Insert(x)
	}
}

```