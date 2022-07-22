package language

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestLocalizer_TextBlock(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputCount:   1,
			outputString: "Block",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputCount:   2,
			outputString: "Blocks",
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

			testTextWithCount(t, i, localizer.TextBlock, table)
		})
	}
}

func TestLocalizer_TextBlockExists(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputStrings: []string{"example.com"},
			outputString: "Block for domain example.com already exists.",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputStrings: []string{"example2.com"},
			outputString: "Block for domain example2.com already exists.",
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

			testTextWith1String(t, i, localizer.TextBlockExists, table)
		})
	}
}

func TestLocalizer_TextBlockSubdomain(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputCount:   1,
			outputString: "Block Subdomain",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputCount:   2,
			outputString: "Block Subdomains",
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

			testTextWithCount(t, i, localizer.TextBlockSubdomain, table)
		})
	}
}

func TestLocalizer_TextBlockedInstance(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputCount:   1,
			outputString: "Blocked Instance",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputCount:   2,
			outputString: "Blocked Instances",
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

			testTextWithCount(t, i, localizer.TextBlockedInstance, table)
		})
	}
}
