package src

import (
	"testing"
)

type testStruct struct {
	name  string
	value float32
}

type testContainerStruct struct {
	value interface{}
}

func TestPointerToInt(t *testing.T) {
	p1 := &testStruct{name: "test", value: 34.5}

	node := &testContainerStruct{value: p1}

	t.Logf("Arbitrary pointer assignment succeeded; %v", node)
}
