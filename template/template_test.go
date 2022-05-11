package libtemplate

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	templates, err := testNewTestTemplates()
	if err != nil {
		t.Errorf("unexpected error creating template: %s", err.Error())

		return
	}
	if templates == nil {
		t.Error("expected templates, got: nil")

		return
	}

	result, err := testExecuteTemplate(templates, "test_test_func", "foo")
	if err != nil {
		t.Errorf("unexpected error creating template: %s", err.Error())

		return
	}
	if expected := "foo bar"; result != expected {
		t.Errorf("unexpected result\n\ngot:\n-------------\n%s\n\nwant:\n-------------\n%s\n", result, expected)

		return
	}
}

func testExecuteTemplate(templates *template.Template, name string, tmplVars interface{}) (string, error) {
	b := new(bytes.Buffer)
	err := templates.ExecuteTemplate(b, name, tmplVars)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func testNewTestTemplates() (*template.Template, error) {
	testFuncs := template.FuncMap{
		"testFunc": func(s string) string {
			return fmt.Sprintf("%s bar", s)
		},
	}

	tmpl, err := New(testFuncs)
	if err != nil {
		return nil, err
	}

	// add test templates
	file, err := os.Open("../test/templates/test.gohtml")
	if err != nil {
		return nil, err
	}
	tmplData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	_, err = tmpl.Parse(string(tmplData))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
