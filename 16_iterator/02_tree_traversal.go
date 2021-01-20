// Iterator - Tree Traversal

// We've already touched the notion that
// the concept of an iterator in Go isn't as
// common as it is in other programming languages.

// But we'll take a look on an example where it's in
// fact relevant and where we can't get away without using
// an iterator.

// What we're going to use is a very simple binary tree.
// So a binary tree is basicaly a data structure which has
// some sort of root, like let's say we have the number one at,
// the root, and then it has two branches.

//		1
//	  /  \
//	 2	  3

// <- So every element can either have one or two or zero branches,
//	  and here the number 1 has branches, and we have the nodes 2 and 3
//	  on the end.
//	  Neither 2 nor 3 have children of their own, but they could ¯\_(ツ)_/¯

// With that, what we want to be able to do is to set up
// a data structure where we actually construct these nodes
// and specify the connections between the nodes.

package main

import "fmt"

// Let's start, with a type.

type Node struct {
	Value               int
	left, right, parent *Node
}

// And now we need to have a bunch of factories,
// for actually initializing all of this and they're
// not going to be that simple.

// First, let's make a factory where we just have the value.
// No information about the parent nor the children.

func NewNode(value int, left, right *Node) *Node {
	n := &Node{Value: value, left: left, right: right}
	left.parent = n
	right.parent = n
	return n
}

func NewTerminalNode(value int) *Node {
	return &Node{Value: value}
}

// <- This one is usable for terminal nodes, because
//	  they don't have any children

// Now, we want to be able to travers our little tree,
// and there are different ways, different algorithms of acheiveing it.
// The three most common ones are:
// -> in-order: 213
// -> preorder: 123
// -> postorder: 231

// We'll use just the in-order, but different approaches
// will result in us having different Iterators.

// Well then, we're going to try is build an inward iterator.
// And it obviously has to have reference to the root node, and
// it also has to have other things besides.

type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

// ↑↑↑ This bool value indicates whether or not
//	   we return the starting value, with that we
//	   avoid indexes starting with -1, like in last example

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{root, root, false}
	// we need to find the left-most element
	for i.Current.left != nil {
		i.Current = i.Current.left
	}

	return i
}

// But now, we're going to have a bunch of utility functions
// for just working with a situation like let's say we want to
// reset the iterator.

func (i *InOrderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
}

// Like in previous example, we need to move to the next element.

func (i *InOrderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
		// ↑↑↑ meaning whoever is iterating this object
		//	   is welcome to take the current value, because
		//	   the current value has the starting value we've ensured
	}

	// we now have to traverse the entire tree, from left to right correctly
	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	}
	p := i.Current.parent
	for p != nil && i.Current == p.right {
		i.Current = p
		p = p.parent
	}
	i.Current = p

	return i.Current != nil
}

// And this works, but let's suppose that we want to have
// a really nicely packaged implementation of both in-order travers
// as well as other forms of traversal.
// We can go ahead and create a new struct.

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

// Recap:
// ->	This has hopefully ilustrated why we would want
//		to construct different iterator objects
// ->	An iterator in this particular case is nothing more
//		than some structure which has obviously a pointer to the
//		elements of whatever it is traversing
// -> 	It also has, in this case, a current pointer which is what we
//		use to access the elements on which the iterator is currently stopped

func main() {
	root := NewNode(
		1,
		NewTerminalNode(2),
		NewTerminalNode(3),
	)
	it := NewInOrderIterator(root)

	for it.MoveNext() {
		fmt.Printf("%d,", it.Current.Value)
	}
	fmt.Println("\b")

	t := NewBinaryTree(root)
	for i := t.InOrder(); i.MoveNext(); {
		fmt.Printf("%d,", i.Current.Value)
	}
	fmt.Println("\b")
}
