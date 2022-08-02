package language

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestLocalizer_TextModeration(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			outputString: "Moderation",
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

			testText(t, i, localizer.TextModeration, table)
		})
	}
}

func TestLocalizer_TextModerator(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputCount:   1,
			outputString: "Moderator",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputCount:   2,
			outputString: "Moderators",
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

			testTextWithCount(t, i, localizer.TextModerator, table)
		})
	}
}
