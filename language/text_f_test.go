package language

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestLocalizer_TextFediverse(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			outputString: "Fediverse",
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

			testText(t, i, localizer.TextFediverse, table)
		})
	}
}

func TestLocalizer_TextFollowing(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			outputString: "Following",
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

			testText(t, i, localizer.TextFollowing, table)
		})
	}
}
