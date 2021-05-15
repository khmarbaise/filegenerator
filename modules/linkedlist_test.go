package modules

import (
	"testing"
)

func TestList_Size_ShouldBe_One(t *testing.T) {
	list := List{}
	list.Add(5)

	if size := list.Size(); size != 1 {
		t.Errorf("size: %v", list.Size())
	}
}

func TestList_Size_ShouldBe_Two(t *testing.T) {
	list := List{}
	list.Add(1)
	list.Add(2)

	if size := list.Size(); size != 2 {
		t.Errorf("size: %v", list.Size())
	}
}
