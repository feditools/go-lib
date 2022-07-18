package http

//revive:disable:add-constant

import (
	"fmt"
	"testing"
)

var testMimeTables = []struct {
	mime   Mime
	suffix Suffix
}{
	{Mime(""), Suffix("")},
	{mime: "image/gif", suffix: "gif"},
	{mime: "image/jpeg", suffix: "jpg"},
	{mime: "image/png", suffix: "png"},
	{mime: "image/svg+xml", suffix: "svg"},
	{mime: "image/webp", suffix: "webp"},
	{mime: "text/html", suffix: "html"},
}

func TestToMime(t *testing.T) {
	t.Parallel()

	for i, table := range testMimeTables {
		i := i
		table := table

		t.Run(fmt.Sprintf("[%d] ToMime %s", i, table.suffix), func(t *testing.T) {
			t.Parallel()

			m := ToMime(table.suffix)
			if m != table.mime {
				t.Errorf("[%d] wrong mime type: got '%s', want '%s'", i, m, table.mime)
			}
		})
	}
}

func TestToSuffix(t *testing.T) {
	t.Parallel()

	for i, table := range testMimeTables {
		i := i
		table := table

		t.Run(fmt.Sprintf("[%d] ToSuffix %s", i, table.mime), func(t *testing.T) {
			t.Parallel()

			s := ToSuffix(table.mime)
			if s != table.suffix {
				t.Errorf("[%d] wrong suffix: got '%s', want '%s'", i, s, table.suffix)
			}
		})
	}
}

//revive:enable:add-constant
