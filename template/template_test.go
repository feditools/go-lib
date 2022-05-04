package template

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	templates, err := New(nil)
	if err != nil {
		t.Errorf("unexpected error creating template: %s", err.Error())
		return
	}
	if templates == nil {
		t.Error("expected templates, got: nil")
		return

	}
}

func addTestTemplates(templates *template.Template) error {
	// open it
	file, err := os.Open("../test/templates/test.gohtml")
	if err != nil {
		return err
	}

	// read it
	tmplData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// It can now be parsed as a string.
	_, err = templates.Parse(string(tmplData))
	if err != nil {
		return err
	}

	return nil
}

func testExecuteTemplate(templates *template.Template, name string, tmplVars interface{}) (string, error) {
	b := new(bytes.Buffer)
	err := templates.ExecuteTemplate(b, name, tmplVars)
	if err != nil {
		return "", err
	}
	return string(b.Bytes()), nil
}
