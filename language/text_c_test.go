package language

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestLocalizer_TextClient(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x language.Tag
		c int
		n string
		l language.Tag
	}{
		{language.English, 1, "Client", language.English},
		{language.English, 2, "Clients", language.English},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Translating to %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			localizer, err := langMod.NewLocalizer(table.x.String())
			if err != nil {
				t.Errorf("[%d] can't get localizer for %s: %s", i, table.x, err.Error())
				return
			}

			result := localizer.TextClient(table.c)
			if result.String() != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result.String(), table.n)
			}
			if result.Language() != table.l {
				t.Errorf("[%d] got invalid language for %s, got: %v, want: %v,", i, table.x, result.Language(), table.l)
			}
		})
	}
}

func TestLocalizer_TextClientID(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x language.Tag
		c int
		n string
		l language.Tag
	}{
		{language.English, 1, "Client ID", language.English},
		{language.English, 2, "Client IDs", language.English},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Translating to %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			localizer, err := langMod.NewLocalizer(table.x.String())
			if err != nil {
				t.Errorf("[%d] can't get localizer for %s: %s", i, table.x, err.Error())
				return
			}

			result := localizer.TextClientID(table.c)
			if result.String() != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result.String(), table.n)
			}
			if result.Language() != table.l {
				t.Errorf("[%d] got invalid language for %s, got: %v, want: %v,", i, table.x, result.Language(), table.l)
			}
		})
	}
}

func TestLocalizer_TextClientSecret(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x language.Tag
		c int
		n string
		l language.Tag
	}{
		{language.English, 1, "Client Secret", language.English},
		{language.English, 2, "Client Secrets", language.English},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Translating to %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			localizer, err := langMod.NewLocalizer(table.x.String())
			if err != nil {
				t.Errorf("[%d] can't get localizer for %s: %s", i, table.x, err.Error())
				return
			}

			result := localizer.TextClientSecret(table.c)
			if result.String() != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result.String(), table.n)
			}
			if result.Language() != table.l {
				t.Errorf("[%d] got invalid language for %s, got: %v, want: %v,", i, table.x, result.Language(), table.l)
			}
		})
	}
}

func TestLocalizer_TextCreate(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x language.Tag
		n string
		l language.Tag
	}{
		{language.English, "Create", language.English},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Translating to %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			localizer, err := langMod.NewLocalizer(table.x.String())
			if err != nil {
				t.Errorf("[%d] can't get localizer for %s: %s", i, table.x, err.Error())
				return
			}

			result := localizer.TextCreate()
			if result.String() != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result.String(), table.n)
			}
			if result.Language() != table.l {
				t.Errorf("[%d] got invalid language for %s, got: %v, want: %v,", i, table.x, result.Language(), table.l)
			}
		})
	}
}
