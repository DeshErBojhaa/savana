package data

import (
	"errors"
	"testing"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestMinHeapBuild(t *testing.T) {
	h := NewMinHeap(5)
	for i := 1; i <= 5; i++ {
		x, err := h.GetNearestSlot()
		if err != nil {
			t.Fatal(err)
		}
		if i != x {
			t.Fatalf("Expected %d got %d", i, x)
		}
		t.Logf("%s Expected %d Got %d", Success, i, x)
	}
}

func TestMinHeapInsert(t *testing.T) {
	h := NewMinHeap(4)
	for i := 1; i <= 4; i++ {
		x, err := h.GetNearestSlot()
		if err != nil {
			t.Fatal(err)
		}
		if i != x {
			t.Logf("Expected %d got %d", i, x)
		}
	}
	extra := []int{4, 2, 40, 30}
	for _, i := range extra {
		h.Insert(i)
	}

	match := []int{2, 4, 30, 40}
	for _, i := range match {
		x, err := h.GetNearestSlot()
		if err != nil {
			t.Fatal(err)
		}
		if i != x {
			t.Logf("Expected %d got %d", i, x)
		}
		t.Logf("%s Expected %d Got %d", Success, i, x)
	}
}

func TestEmptyHeap(t *testing.T) {
	h := NewMinHeap(1)
	_, err := h.GetNearestSlot()
	if err != nil {
		t.Fatal(err)
	}
	_, err = h.GetNearestSlot()
	if errors.Is(err, ErrHeapEmpty) {
		t.Logf("%s Expected %v Got %v", Success, err, ErrHeapEmpty)
	} else {
		t.Logf("%s Expected %v Got %v", Failed, err, ErrHeapEmpty)
	}
}

func TestHeapFull(t *testing.T) {
	h := NewMinHeap(1)
	err := h.Insert(100)
	if errors.Is(err, ErrHeapFull) {
		t.Logf("%s Expected %v Got %v", Success, err, ErrHeapEmpty)
	} else {
		t.Logf("%s Expected %v Got %v", Failed, err, ErrHeapEmpty)
	}
}
