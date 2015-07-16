package tree

import (
	"fmt"
)

type o_node_fac struct {
	current int
	k       int
}

type testvalue struct {
	Value
	_int int
}

func (t *testvalue) Less(x interface{}) (result bool) {
	result = true
	return
}

func (t *testvalue) String() string {
	return fmt.Sprintf("%v", t._int)
}

func (o *o_node_fac) New() (n *node) {
	n = o_node(o.current, o.k)
	o.current++
	return
}

func o_node(order int, k int) (n *node) {

	val := &testvalue{_int: order}

	n = new_node(val, k)

	return
}
