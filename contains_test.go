package lib

import (
	"fmt"
	"testing"
)

//revive:disable:add-constant

func TestContainsString(t *testing.T) {
	t.Parallel()

	stack := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	tables := []struct {
		x string
		n bool
	}{
		{"one", true},
		{"four", true},
		{"foo", false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Checking stack for %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ContainsString(stack, table.x)
			if result != table.n {
				t.Errorf("Sum was incorrect, got: %v, want: %v.", result, table.n)
			}
		})
	}
}

func TestContainsOneOfStrings(t *testing.T) {
	t.Parallel()

	stack := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	tables := []struct {
		x []string
		n bool
	}{
		{[]string{"one", "two", "three"}, true},
		{[]string{"foo", "five", "bar"}, true},
		{[]string{"foo", "bar", "fizz"}, false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Checking stack for %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ContainsOneOfStrings(stack, table.x)
			if result != table.n {
				t.Errorf("Sum was incorrect, got: %v, want: %v.", result, table.n)
			}
		})
	}
}

//revive:enable:add-constant
