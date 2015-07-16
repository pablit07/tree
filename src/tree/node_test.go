package tree

import (
	"testing"
)

func TestInstantiate(t *testing.T) {
	node := &node{}

	t.Logf("Can instantiate node: %v", node)
}

func TestInstantiateWithValue(t *testing.T) {
	val := &testvalue{}
	node := &node{value: val}

	t.Logf("Can instantiate node with value: %v", node)
}

func TestTraverseDepth1PartialFull(t *testing.T) {

	n1 := &node{}
	n1.children = make([]*node, 4)
	n1.children[0] = &node{}
	n1.children[1] = &node{}

	t.Logf("Length of children: ", len(n1.children))

	action := func(i int, el *node) {
		t.Logf("Was passed node: %v", el)
		return
	}

	n1.traverse(action)
}

func TestTraverseDepth1Empty(t *testing.T) {
	n1 := &node{children: make([]*node, 4)}

	action := func(i int, el *node) {
		t.Logf("Was passed node: %v", el)
		if el == nil {
			el = &node{}
		}
		return
	}

	n1.traverse(action)
}

func TestTraverseDepth1Full(t *testing.T) {

	n1 := &node{}
	n1.children = append(n1.children, &node{}, &node{}, &node{}, &node{})

	t.Logf("Length of children: ", len(n1.children))

	action := func(i int, el *node) {
		t.Logf("Was passed node: %v", el)
		return
	}

	n1.traverse(action)
}

func TestTraverseDepth2Full(t *testing.T) {

	recursion_limit := 100
	loops := 0
	fac := &o_node_fac{k: 4}

	n1 := &node{children: make([]*node, 0)}

	n1.children = append(n1.children, fac.New(), fac.New(), fac.New(), fac.New())

	t.Logf("Children: %v", n1.children)

	for _, el := range n1.children {
		t.Logf("setting up subnode: %v", el)
		el.children[0] = fac.New()
		el.children[1] = fac.New()
		el.children[2] = fac.New()
		el.children[3] = fac.New()
	}

	t.Logf("Length of children: ", len(n1.children))

	action := func(i int, el *node) {

		if el != nil {
			t.Logf("Was passed node: %v with value: %v", el, el.value)
		} else {
			t.Logf("Was passed node: %v", el)
		}

		loops++
		if loops > recursion_limit {
			t.Error("FAIL: Recursion Limit Reached")
			t.FailNow()
		}
		return
	}

	n1.traverse(action)
}

func TestNodeCtor(t *testing.T) {

	val, k := &testvalue{}, 20
	node := new_node(val, k)

	if len(node.children) != k {
		t.Error("FAIL: node ctor did not create child array size of k")
	}

	if node.value != val {
		t.Error("FAIL: node ctor did not create node with value set to parameter")
	}
}

func TestTraverseDepth2FullBreadthFirst(t *testing.T) {

	recursion_limit := 100
	loops := 0

	n1 := &node{children: make([]*node, 0)}

	n1.children = append(n1.children, &node{}, &node{}, &node{}, &node{})

	for _, el := range n1.children {
		t.Logf("setting up subnode: %v", el)
		el.children = make([]*node, 4)
		el.children = append(el.children, &node{}, &node{}, &node{}, &node{})
	}

	t.Logf("Length of children: ", len(n1.children))

	action := func(i int, el *node) {
		t.Logf("Was passed node: %v", el)

		loops++
		if loops > recursion_limit {
			t.Error("FAIL: Recursion Limit Reached")
			t.FailNow()
		}
		return
	}

	n1.traverse_breadth_first(action)
}

func TestMakeNode(t *testing.T) {
	fac := &o_node_fac{}
	n := fac.New()
	_ = n
}
