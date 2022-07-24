package language

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestLocalizer_TextEditBlock(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputStrings: []string{"example.com"},
			outputString: "Edit Block example.com",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputStrings: []string{"example2.com"},
			outputString: "Edit Block example2.com",
			outputLang:   language.English,
		},
	}

	langMod, _ := New()
	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf(testTranslatedTo, i, table.inputLang)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			localizer, err := langMod.NewLocalizer(table.inputLang.String())
			if err != nil {
				t.Errorf(testCantGetLocalizer, i, table.inputLang, err.Error())

				return
			}

			testTextWith1String(t, i, localizer.TextEditBlockDomain, table)
		})
	}
}
