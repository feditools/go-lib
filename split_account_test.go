package lib

import (
	"errors"
	"fmt"
	"testing"
)

func TestSplitAccount(t *testing.T) {
	t.Parallel()

	tables := []struct {
		account  string
		username string
		domain   string
		error    error
	}{
		{"test@example.com", "test", "example.com", nil},
		{"@test@example.com", "test", "example.com", nil},
		{"example.com", "", "", ErrInvalidAccountFormat},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running SplitAccount on %s", i, table.account)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			username, domain, err := SplitAccount(table.account)
			if !errors.Is(err, table.error) {
				t.Errorf("[%d] invalid error, got: '%v', want: '%v'", i, err, table.error)
			}
			if username != table.username {
				t.Errorf("[%d] invalid username, got: '%v', want: '%v'", i, username, table.username)
			}
			if domain != table.domain {
				t.Errorf("[%d] invalid domain, got: '%v', want: '%v'", i, domain, table.domain)
			}
		})
	}
}
