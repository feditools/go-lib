package lib

import (
	"testing"
)

func TestSplitAccount(t *testing.T) {
	t.Parallel()

	tables := []struct {
		name     string
		account  string
		username string
		domain   string
		error    error
	}{
		{"testing test@example.com", "test@example.com", "test", "example.com", nil},
		{"testing @test@example.com", "@test@example.com", "test", "example.com", nil},
		{"testing example.com", "example.com", "", "", ErrInvalidAccountFormat},
	}

	for i, table := range tables {
		i := i
		table := table

		t.Run(table.name, func(t *testing.T) {
			t.Parallel()

			username, domain, err := SplitAccount(table.account)
			if err != table.error {
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
