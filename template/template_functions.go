package template

import (
	"github.com/feditools/go-lib"
	"html/template"
)

const (
	funcNameDec                         = "dec"
	funcNameFormInputClass              = "formInputClass"
	funcNameFormInputLabelDisplayTop    = "formInputLabelDisplayTop"
	funcNameFormInputLabelDisplayBottom = "formInputLabelDisplayBottom"
	funcNameHTMLSafe                    = "htmlSafe"
	funcNameJSSafe                      = "jsSafe"
	funcNameInc                         = "inc"
)

var (
	funcDec = func(i int) int {
		i--

		return i
	}

	funcFormInputClassCheckType = []string{
		string(FormInputTypeCheckbox),
		string(FormInputTypeRadio),
	}
	funcFormInputClass = func(t string) string {
		switch {
		case lib.ContainsString(funcFormInputClassCheckType, t):
			return "form-check-input"
		default:
			return "form-control"
		}
	}

	formInputLabelDisplayBottomTypes = []string{
		string(FormInputTypeCheckbox),
		string(FormInputTypeRadio),
	}
	funcFormInputLabelDisplayBottom = func(t string) bool {
		return lib.ContainsString(formInputLabelDisplayBottomTypes, t)
	}

	formInputLabelDisplayTopTypes = []string{
		string(FormInputTypeFile),
		string(FormInputTypePassword),
		string(FormInputTypeText),
	}
	funcFormInputLabelDisplayTop = func(t string) bool {
		return lib.ContainsString(formInputLabelDisplayTopTypes, t)
	}

	funcHTMLSafe = func(html string) template.HTML {
		/* #nosec G203 */
		return template.HTML(html)
	}

	funcJSSafe = func(javascript string) template.JS {
		/* #nosec G203 */
		return template.JS(javascript)
	}

	funcInc = func(i int) int {
		i++

		return i
	}

	defaultFunctions = template.FuncMap{
		funcNameDec:                         funcDec,
		funcNameFormInputClass:              funcFormInputClass,
		funcNameFormInputLabelDisplayBottom: funcFormInputLabelDisplayBottom,
		funcNameFormInputLabelDisplayTop:    funcFormInputLabelDisplayTop,
		funcNameHTMLSafe:                    funcHTMLSafe,
		funcNameJSSafe:                      funcJSSafe,
		funcNameInc:                         funcInc,
	}
)
