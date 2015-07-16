package tree

import (
	"testing"
)

func TestInstantiateTree(t *testing.T) {
	tree := &Tree{}

	t.Logf("Can instantiate tree: %v", tree)
}

func TestAddOneToEmpty(t *testing.T) {
	tree := &Tree{K: 2}

	key := "1"
	value := &testvalue{}

	tree.Add(key, value)

	if tree.root == nil {
		t.Error("Failed; tree root is nil")
	}
}

func TestAddOneToSizeOne(t *testing.T) {
	tree := &Tree{K: 2}

	tree.Add("1", &testvalue{})

	tree.Add("2", &testvalue{})

	if len(tree.root.children) < 1 || tree.root.children[0] == nil {
		t.Error("Failed; root has no first child")
	}
}

func TestAddOneToSizeTwo(t *testing.T) {
	tree := &Tree{K: 2}

	tree.Add("1", &testvalue{})

	tree.Add("2", &testvalue{})

	tree.Add("3", &testvalue{})

	if len(tree.root.children) < 2 || tree.root.children[1] == nil {
		t.Error("Failed; root has no second child")
	}
}

func AddAWholeBunch(t *testing.T) {
	tree := &Tree{K: 2}

	tree.Add("1", &testvalue{})

	tree.Add("2", &testvalue{})

	tree.Add("3", &testvalue{})

	if len(tree.root.children) < 2 || tree.root.children[1] == nil {
		t.Error("Failed; root has no second child")
	}
}
