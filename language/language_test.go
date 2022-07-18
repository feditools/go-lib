package language

import (
	"testing"

	"golang.org/x/text/language"
)

const (
	testCantGetLocalizer      = "[%d] can't get localizer for %s: %s"
	testGotInvalidLanguage    = "[%d] got invalid language for %s, got: %v, want: %v"
	testGotInvalidTranslation = "[%d] got invalid translation for %s, got: %v, want: %v"
	testTranslatedTo          = "[%d] Translating to %s"
)

func TestNew(t *testing.T) {
	langMod, err := New()
	if err != nil {
		t.Errorf("can't get new language module: %s", err.Error())

		return
	}

	if langMod == nil {
		t.Errorf("language module is nil")

		return
	}

	if langMod.langBundle == nil {
		t.Errorf("language module's bundle is nil")

		return
	}

	if langMod.Language() != DefaultLanguage {
		t.Errorf("got invalid language, got: %v, want: %v,", langMod.Language().String(), DefaultLanguage.String())

		return
	}
}

// text testers

type testTextTable struct {
	inputLang    language.Tag
	inputCount   int
	inputStrings []string

	outputString string
	outputLang   language.Tag
}

func testText(t *testing.T, tid int, translate func() *LocalizedString, table testTextTable) {
	t.Helper()

	result := translate()
	testTextCheckResults(t, tid, result, table)
}

func testTextWithCount(t *testing.T, tid int, translate func(int) *LocalizedString, table testTextTable) {
	t.Helper()

	result := translate(table.inputCount)
	testTextCheckResults(t, tid, result, table)
}

func testTextWith1String(t *testing.T, tid int, translate func(string) *LocalizedString, table testTextTable) {
	t.Helper()

	result := translate(table.inputStrings[0])
	testTextCheckResults(t, tid, result, table)
}

func testTextCheckResults(t *testing.T, tid int, result *LocalizedString, table testTextTable) {
	t.Helper()

	if result.String() != table.outputString {
		t.Errorf(testGotInvalidTranslation, tid, table.inputLang, result.String(), table.outputString)
	}
	if result.Language() != table.outputLang {
		t.Errorf(testGotInvalidLanguage, tid, table.inputLang, result.Language(), table.outputLang)
	}
}
