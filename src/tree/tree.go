package tree

// import (
// 	""
// )

type Key interface {
	Less(x interface{}) bool
}

type node struct {
	key      Key
	value    interface{}
	children []*node
	k        int
}

func new_node(key Key, value interface{}, k int) *node {
	return &node{children: make([]*node, k), value: value, k: k}
}

func (n *node) traverse(action func(i int, el *node)) {

	// is leaf
	if n == nil {
		return
	}
	action(0, n)
	for _, el := range n.children {
		// recurse first, or action first?
		el.traverse(action)
	}
	return
}

func (n *node) traverse_depth_first_bottom_up(action func(i int, el *node)) {

	// is leaf
	if n == nil || len(n.children) == 0 {
		return
	}
	for i, el := range n.children {
		// recurse first, or action first?
		el.traverse(action)
		action(i, el)
	}
	return
}

func (n *node) traverse_breadth_first(action func(i int, el *node)) {

	// is leaf
	if n == nil || len(n.children) == 0 {
		return
	}

	queue := make([]*node, 0)

	for _, el := range n.children {
		queue = append(queue, el)
	}
	for i := 0; i < len(queue); i++ {
		el := queue[i]
		// visit and retrieve children
		action(i, el)
		children := el.children
		// queue up children
		for _, child := range children {
			if child != nil {
				queue = append(queue, child)
			}
		}

	}
}

func (n *node) traverse_breadth_first_right_left(action func(i int, el *node)) {

	// is leaf
	if n == nil || len(n.children) == 0 {
		return
	}
	for i := len(n.children) - 1; i >= 0; i-- {
		el := n.children[i]
		action(i, el)
	}
	for i := len(n.children) - 1; i >= 0; i-- {
		el := n.children[i]
		el.traverse(action)
	}

	return
}

type Tree struct {
	K    int
	root *node
}

func (t *Tree) Add(key string, value Value) {
	if t.root == nil {
		t.root = new_node(value, t.K)
		return
	}

	found := false
	var targetParent *node
	var childIndex int

	t.root.traverse(func(i int, el *node) {
		if found || el == nil {
			return
		}

		for i := 0; i < len(el.children); i++ {
			if el.children[i] == nil {
				found = true
				targetParent = el
				childIndex = i
				return
			}
		}
	})

	targetParent.children[childIndex] = new_node(value, t.K)
}

func (t *Tree) Search(key string) Value {
	return nil
}
