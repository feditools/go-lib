package template

import (
	"fmt"
	"testing"
)

func TestTemplateFunctions(t *testing.T) {
	t.Parallel()

	templates, err := testNewTestTemplates()
	if err != nil {
		t.Errorf("init: %s", err.Error())

		return
	}

	//revive:disable:add-constant
	tables := []struct {
		templateName string
		templateVars interface{}
		output       string
	}{
		{
			"test_dec",
			481,
			"480",
		},
		{
			"test_html_safe",
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,./;'[]\\`~!@#$%^&*()_+{}|:\"<>?",
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,./;'[]\\`~!@#$%^&*()_+{}|:\"<>?",
		},
		{
			"test_inc",
			48,
			"49",
		},
		{
			"test_passthrough",
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,./;'[]\\`~!@#$%^&*()_+{}|:\"<>?",
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,./;&#39;[]\\`~!@#$%^&amp;*()_&#43;{}|:&#34;&lt;&gt;?",
		},
	}
	//revive:enable:add-constant

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running making pagination", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result, err := testExecuteTemplate(templates, table.templateName, table.templateVars)
			if err != nil {
				t.Errorf("unexpected error creating template: %s", err.Error())

				return
			}
			if result != table.output {
				t.Errorf("[%d] unexpected result\n\ngot:\n-------------\n%s\n\nwant:\n-------------\n%s\n", i, result, table.output)

				return
			}
		})
	}
}
