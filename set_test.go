package graphs

import (
	"testing"
)

func TestSetAdd(t *testing.T) {
	set := NewSet[string]()
	foo := "foo"

	if added := set.Add(foo); !added {
		t.Error("foo should not exist")
	}

	if added := set.Add(foo); added {
		t.Error("foo should exist")
	}
}

func TestSetLen(t *testing.T) {
	set := NewSet[string]()

	if set.Len() != 0 {
		t.Error("set length should be 0")
	}

	set.Add("foo")
	if set.Len() != 1 {
		t.Error("set length should be 1")
	}

	set.Add("bar")
	if set.Len() != 2 {
		t.Error("set length should be 2")
	}

	set.Add("bar")
	if set.Len() != 2 {
		t.Error("set length should be 2")
	}
}

func TestSetEquals(t *testing.T) {
	s1 := NewSet[string]()
	s2 := NewSet[string]()

	if s1.Equals(nil) {
		t.Error("no set is equal to a nil set")
	}

	if !s1.Equals(s2) {
		t.Error("two empty sets should be equal")
	}

	s1.Add("foo")
	s2.Add("foo")

	if !s1.Equals(s2) {
		t.Error("two sets with both one element should be equal")
	}

	s1.Add("moo")
	if s1.Equals(s2) {
		t.Error("two sets with different length should not be equal")
	}

	s2.Add("cow")
	if s1.Equals(s2) {
		t.Error("two sets with different elements should not be equal")
	}
}

func TestSetContains(t *testing.T) {
	set := NewSet[string]()
	set.Add("foo")

	if !set.Contains("foo") {
		t.Error("set should contain foo")
	}

	if set.Contains("bar") {
		t.Error("set should not contain bar")
	}
}

func TestSetMerge(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("foo")

	s2 := NewSet[string]()
	s2.Add("bar")

	s2.Merge(s1)
	if s2.Len() != 2 {
		t.Error("merged set should have two elements")
	}
}

func TestSetRemove(t *testing.T) {
	set := NewSet[string]()
	set.Add("foo")
	set.Add("bar")
	set.Remove("foo")

	if set.Len() != 1 {
		t.Error("set should contain one element")
	}

	if !set.Contains("bar") {
		t.Error("set should contain bar")
	}
}

func TestSetAny(t *testing.T) {
	set := NewSetWithElements[string]("foo", "bar")
	if e := set.Any(); e != "bar" && e != "foo" {
		t.Error("any should return bar or foo")
	}
}

func TestEach(t *testing.T) {
	set := NewSetWithElements[string]("foo", "bar", "baz")
	count := 0

	set.Each(func(element string, stop *bool) {
		count += 1
	})

	if count != 3 {
		t.Error("count should be 3")
	}
}

func TestIter(t *testing.T) {
	set := NewSetWithElements[string]("foo", "bar", "baz")
	count := 0

	for _ = range set.Iter() {
		count += 1
	}

	if count != 3 {
		t.Error("count should be 3")
	}
}
