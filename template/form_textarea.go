package template

import liblanguage "github.com/feditools/go-lib/language"

// FormTextarea is a templated form textarea.
type FormTextarea struct {
	ID         string
	Name       string
	Rows       int
	Label      *liblanguage.LocalizedString
	LabelClass string
	Value      string
	Disabled   bool
	Required   bool
	Validation *FormValidation
}
