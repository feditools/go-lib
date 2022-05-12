package http

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestGetPageLang(t *testing.T) {
	t.Parallel()

	//revive:disable:add-constant
	tables := []struct {
		name string

		query       string
		header      string
		defaultLang string

		output string
	}{
		{
			name:        "default",
			query:       "",
			header:      "",
			defaultLang: language.English.String(),
			output:      language.English.String(),
		},
		{
			name:        "query set",
			query:       "es",
			header:      "",
			defaultLang: language.Chinese.String(),
			output:      language.Spanish.String(),
		},
		{
			name:        "header set",
			query:       "",
			header:      "en;q=0.5,de",
			defaultLang: language.Hindi.String(),
			output:      language.German.String(),
		},
		{
			name:        "all set",
			query:       "no",
			header:      "ca;q=0.9,fr",
			defaultLang: language.Russian.String(),
			output:      language.Norwegian.String(),
		},
	}
	//revive:enable:add-constant

	for i, table := range tables {
		i := i
		table := table

		t.Run(fmt.Sprintf("[%d] %s", i, table.name), func(t *testing.T) {
			t.Parallel()

			o := GetPageLang(table.query, table.header, table.defaultLang)
			if o != table.output {
				t.Errorf("[%d] wrong lanauge: got '%s', want '%s'", i, o, table.output)
			}
		})
	}
}
