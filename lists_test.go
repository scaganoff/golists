package lists

import (
	"testing"
)

func TestInclude(t *testing.T) {

	list := StringSlice{"foo", "bar", "baz"}

	if !list.Includes("foo") {
		t.Fatal("Expected list to include foo")
	}

	if list.Includes("biffo") {
		t.Fatal("Didn't Expecte list to include biffo")
	}
}

func TestEquals(t *testing.T) {
	l1 := StringSlice{"foo", "bar", "baz"}
	l2 := StringSlice{"foo", "baz", "biffo"}
	l3 := StringSlice{"baz", "foo", "bar"}

	if l1.Equals(l2) {
		t.Fatal("Failed equals test 1")
	}

	if !l1.Equals(l3) {
		t.Fatal("Failed equals test 1")
	}
}

func TestSubtract(t *testing.T) {

	l1 := StringSlice{"foo", "bar", "baz"}
	l2 := StringSlice{"foo", "baz", "biffo"}
	expect := StringSlice{"bar"}

	l3 := l1.Subtract(l2)
	if !(l3.Equals(expect)) {
		t.Fatal("Unexpected result")
	}
}
