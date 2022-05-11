package liblanguage

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestLocalizer_TextOauth(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			outputString: "OAuth",
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

			testText(t, i, localizer.TextOauth, table)
		})
	}
}

func TestLocalizer_TextOauth20Client(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputCount:   1,
			outputString: "OAuth 2.0 Client",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputCount:   2,
			outputString: "OAuth 2.0 Clients",
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

			testTextWithCount(t, i, localizer.TextOauth20Client, table)
		})
	}
}

func TestLocalizer_TextOauth20Settings(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			outputString: "OAuth 2.0 Settings",
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

			testText(t, i, localizer.TextOauth20Settings, table)
		})
	}
}
