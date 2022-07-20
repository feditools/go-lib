package language

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestLocalizer_TextAccount(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputCount:   1,
			outputString: "Account",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputCount:   2,
			outputString: "Accounts",
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

			testTextWithCount(t, i, localizer.TextAccount, table)
		})
	}
}

func TestLocalizer_TextAddOauth20Client(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputCount:   1,
			outputString: "Add OAuth 2.0 Client",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputCount:   2,
			outputString: "Add OAuth 2.0 Clients",
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

			testTextWithCount(t, i, localizer.TextAddOauth20Client, table)
		})
	}
}

func TestLocalizer_TextAdmin(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			outputString: "Admin",
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

			testText(t, i, localizer.TextAdmin, table)
		})
	}
}

func TestLocalizer_TextAllow(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			outputString: "Allow",
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

			testText(t, i, localizer.TextAllow, table)
		})
	}
}

func TestLocalizer_TextApplicationToken(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputCount:   1,
			outputString: "Application Token",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputCount:   2,
			outputString: "Application Tokens",
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

			testTextWithCount(t, i, localizer.TextApplicationToken, table)
		})
	}
}

func TestLocalizer_TextAuthorize(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			outputString: "Authorize",
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

			testText(t, i, localizer.TextAuthorize, table)
		})
	}
}

func TestLocalizer_TextAuthorizeApplicationDescription(t *testing.T) {
	t.Parallel()

	tables := []testTextTable{
		{
			inputLang:    language.English,
			inputStrings: []string{"test1"},
			outputString: "Authorize test1",
			outputLang:   language.English,
		},
		{
			inputLang:    language.English,
			inputStrings: []string{"test2"},
			outputString: "Authorize test2",
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

			testTextWith1String(t, i, localizer.TextAuthorizeApplicationDescription, table)
		})
	}
}
