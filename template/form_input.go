package template

type FormInputType string

const (
	// FormInputTypeCheckbox is a checkbox html input field.
	FormInputTypeCheckbox FormInputType = "checkbox"
	// FormInputTypeFile is a file html input field.
	FormInputTypeFile FormInputType = "file"
	// FormInputTypeHidden is a hidden html input field.
	FormInputTypeHidden FormInputType = "hidden"
	// FormInputTypePassword is a password html input field.
	FormInputTypePassword FormInputType = "password"
	// FormInputTypeText is a text html input field.
	FormInputTypeText FormInputType = "text"
)

// FormInput is a templated form input.
type FormInput struct {
	ID          string
	Type        FormInputType
	Name        string
	Placeholder string
	Label       *FormLabel
	Value       string
	Disabled    bool
	Required    bool
	Checked     bool
	Validation  *FormValidation
}

// FormValidation is a validation response to a form input.
type FormValidation struct {
	Valid    bool
	Response string
}
