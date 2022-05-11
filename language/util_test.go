package liblanguage

import (
	"fmt"
	"testing"
)

func TestIsEmptyYaml(t *testing.T) {
	t.Parallel()

	tables := []struct {
		input  string
		output bool
	}{
		{"", true},
		{"---", true},
		{"---\n", true},
		{"---\nvalid: yaml", false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] running isEmptyYaml", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := isEmptyYaml([]byte(table.input))
			if result != table.output {
				t.Errorf("[%d] got invalid response, got: %v, want: %v,", i, result, table.output)
			}
		})
	}
}
