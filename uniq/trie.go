package uniq

import (
	"iter"
)

type Node struct {
	Value    byte
	Position int
	Children []*Node
}

func (n Node) Has(value []byte) bool {
	var current *Node = &n
	for _, b := range value {
		var found bool
		for _, n := range current.Children {
			if n.Value == b {
				current = n
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func (n *Node) Insert(value []byte, position int) {
	_ = n.insert(value, position)
}

func (n *Node) insert(value []byte, position int) *Node {
	// if it exists, don't insert
	for _, c := range n.Children {
		if c.Value == value[0] {
			return c.insert(value[1:], position)
		}
	}

	// create a new child if there are no children
	child := &Node{
		Value:    value[0],
		Position: -1,
	}
	n.Children = append(n.Children, child)
	if len(value) == 1 {
		child.Position = position
		return child
	}
	return child.insert(value[1:], position)
}

type Value struct {
	Value    []byte
	Position int
}

func (n Node) walk(prefix []byte, yield func(Value) bool) {
	newPrefix := append(prefix, n.Value)
	if n.Position >= 0 {
		if !yield(Value{Value: newPrefix, Position: n.Position}) {
			return
		}
	}
	for _, n := range n.Children {
		n.walk(newPrefix, yield)
	}
}

func (n Node) Nodes() iter.Seq[Value] {
	return func(yield func(Value) bool) {
		n.walk(nil, yield)
	}
}
