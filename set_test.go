package goset_test

import (
	"testing"

	goset "github.com/albertsundjaja/go-set"
)

func TestSet(t *testing.T) {
	s := goset.New[int64]()
	s.Add(1)
	if len(s) != 1 {
		t.Errorf("Expected set to contain 1 element but got %v", len(s))
	}
	if _, ok := s[1]; !ok {
		t.Errorf("Expected set to contain 1 as element but it's not in the set")
	}
	if !s.Contains(1) {
		t.Errorf("Expected Contains to return true for 1")
	}
	if s.Contains(2) {
		t.Errorf("Expected Contains to return false for element not in the set")
	}
	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("Expected Remove to remove the element, but it is still in the set")
	}
}
