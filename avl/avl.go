package avl

import (
	"fmt"
)

// LessThanFunc takes two interfaces and determines which one is less
type LessThanFunc func(a, b interface{}) bool

// AVL will contain everything needed for our AVL tree
type AVL struct {
	root     *Node
	lessThan LessThanFunc
}

// Node will support a left and a right child
type Node struct {
	key                   interface{}
	height                int
	LeftChild, RightChild *Node
}

// New will return an avl tree
func New(lessThan LessThanFunc) *AVL {
	return &AVL{
		lessThan: lessThan,
	}
}

// Insert given the root node, will insert an item
func (a *AVL) Insert(key interface{}) *Node {
	// plain ole insert, going down the left side of the tree
	a.root = a.insert(a.root, key)
	fmt.Println("root", a.root.key, a.root.height)
	return a.root
}

// isBalanced determines if a root node is balanced or not
func (a *AVL) isBalanced(nodeA, nodeB *Node) bool {
	if a.nodeHeight(nodeA)-a.nodeHeight(nodeB) == 2 {
		return false
	}
	return true
}

// nodeHeight determines the height of a node
func (a *AVL) nodeHeight(node *Node) int {
	if node == nil {
		return -1
	}
	return node.height
}

// heightOfHighestNode will compute the taller node height
func (a *AVL) heightOfHighestNode(node *Node) int {
	lch := a.nodeHeight(node.LeftChild)
	rch := a.nodeHeight(node.RightChild)

	if lch > rch {
		return lch
	}
	return rch
}

// insert ...
func (a *AVL) insert(root *Node, key interface{}) *Node {
	if root == nil {
		// base case, and new insert case as well
		root = &Node{key: key, height: 0, LeftChild: nil, RightChild: nil}
		root.height = a.heightOfHighestNode(root) + 1
		return root
	}

	if a.lessThan(key, root.key) {
		// go down the left side of the tree
		root.LeftChild = a.insert(root.LeftChild, key)
		// is it balanced?
		if !a.isBalanced(root.LeftChild, root.RightChild) {
			// bummer , we need to balance ...
			if a.lessThan(key, root.LeftChild.key) {
				// left, left => right rotate
				root = a.rotateRight(root)
			} else {
				// left, right => left right rotate
				root = a.rotateLeftRight(root)
			}
		}
	}

	if !a.lessThan(key, root.key) {
		// go down the right side of the tree
		root.RightChild = a.insert(root.RightChild, key)
		// is it balanced?
		if !a.isBalanced(root.RightChild, root.LeftChild) {
			// bummer , we need to balance ...
			if !a.lessThan(key, root.RightChild.key) {
				// right, right => left rotate
				root = a.rotateLeft(root)
			} else {
				// right, left rotate
				root = a.rotateRightLeft(root)
			}
		}
	}

	// calculate the root height
	root.height = a.heightOfHighestNode(root) + 1
	return root
}

// rotateRight a given node
func (a *AVL) rotateRight(root *Node) *Node {
	node := root.LeftChild
	root.LeftChild = node.RightChild
	node.RightChild = root
	root.height = a.heightOfHighestNode(root) + 1
	node.height = a.heightOfHighestNode(node) + 1
	return node
}

// rotateLeft a given node
func (a *AVL) rotateLeft(root *Node) *Node {
	node := root.RightChild
	root.RightChild = node.LeftChild
	node.LeftChild = root
	root.height = a.heightOfHighestNode(root) + 1
	node.height = a.heightOfHighestNode(node) + 1
	return node
}

// rotateLeftRight
func (a *AVL) rotateLeftRight(root *Node) *Node {
	root.LeftChild = a.rotateLeft(root.LeftChild)
	return a.rotateRight(root)
}

// rotateRightLeft
func (a *AVL) rotateRightLeft(root *Node) *Node {
	root.RightChild = a.rotateRight(root.RightChild)
	return a.rotateRight(root)
}

// Ints ...
func Ints(a, b interface{}) bool {
	ak, _ := a.(int)
	bk, _ := b.(int)
	return ak < bk
}

// Strings ...
func Strings(a, b interface{}) bool {
	ak, _ := a.(string)
	bk, _ := b.(string)
	return ak < bk
}
