package arrays_test

import (
	"testing"

	"github.com/gabrieldebem/arrays"
)

func TestFilterWithInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	odd := arrays.Filter[int](numbers, func(n int) bool {
		return n%2 == 0
	})

	if len(odd) != 4 {
		t.Errorf("Expected length of 4, got %d", len(odd))
	}
}

func TestFilterWithString(t *testing.T) {
	words := []string{"gopher", "monkey", "robot", "ninja"}
	filtered := arrays.Filter[string](words, func(s string) bool {
		return len(s) > 5
	})

	if len(filtered) != 2 {
		t.Errorf("Expected length of 2, got %d", len(filtered))
	}
}

func TestMapWithInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mapped := arrays.Map[int, int](numbers, func(n int) int {
		return n * 2
	})

	if len(mapped) != 9 {
		t.Errorf("Expected length of 9, got %d", len(mapped))
	}

	for i, v := range mapped {
		if v != numbers[i]*2 {
			t.Errorf("Expected %d, got %d", numbers[i]*2, v)
		}
	}
}

func TestMapceWithString(t *testing.T) {
	words := []string{"gopher", "monkey", "robot", "ninja"}
	mapped := arrays.Map[string, string](words, func(s string) string {
		return s + "!"
	})

	if len(mapped) != 4 {
		t.Errorf("Expected length of 4, got %d", len(mapped))
	}

	for i, v := range mapped {
		if v != words[i]+"!" {
			t.Errorf("Expected %s, got %s", words[i]+"!", v)
		}
	}
}

func TestReduceceWithInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reduced := arrays.Reduce[int](numbers, func(a, b int) int {
		return a + b
	})

	if reduced != 45 {
		t.Errorf("Expected 45, got %d", reduced)
	}
}

func TestReduceceWithString(t *testing.T) {
	words := []string{"gopher", "monkey", "robot", "ninja"}
	reduced := arrays.Reduce[string](words, func(a, b string) string {
		return a + " " + b
	})

	if reduced != "gopher monkey robot ninja" {
		t.Errorf("Expected 'gopher monkey robot ninja', got '%s'", reduced)
	}
}
