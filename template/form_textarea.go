package template

// FormTextarea is a templated form textarea.
type FormTextarea struct {
	ID         string
	Name       string
	Rows       int
	Label      *FormLabel
	Value      string
	Disabled   bool
	Required   bool
	Validation *FormValidation
}
