package language

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestLocalizer_TextDashboard(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x language.Tag
		c int
		n string
		l language.Tag
	}{
		{language.English, 1, "Dashboard", language.English},
		{language.English, 2, "Dashboards", language.English},
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

			result := localizer.TextDashboard(table.c)
			if result.String() != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result.String(), table.n)
			}
			if result.Language() != table.l {
				t.Errorf("[%d] got invalid language for %s, got: %v, want: %v,", i, table.x, result.Language(), table.l)
			}
		})
	}
}

func TestLocalizer_TextDescription(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x language.Tag
		c int
		n string
		l language.Tag
	}{
		{language.English, 1, "Description", language.English},
		{language.English, 2, "Descriptions", language.English},
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

			result := localizer.TextDescription(table.c)
			if result.String() != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result.String(), table.n)
			}
			if result.Language() != table.l {
				t.Errorf("[%d] got invalid language for %s, got: %v, want: %v,", i, table.x, result.Language(), table.l)
			}
		})
	}
}
