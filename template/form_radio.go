package template

import (
	"github.com/feditools/go-lib/language"
	"strconv"
)

// FormRadio is a templated form radio input.
type FormRadio struct {
	ID         string
	Name       string
	Values     map[string]*language.LocalizedString
	Selected   string
	Disabled   bool
	Required   bool
	Validation *FormValidation
}

func (f *FormRadio) GetFormInputs() *[]FormInput {
	formInputs := make([]FormInput, len(f.Values))

	index := int64(0)
	lastInput := int64(len(f.Values) - 1)
	for value, label := range f.Values {
		formInputs[index] = FormInput{
			ID:    f.ID + strconv.FormatInt(index, 10),
			Type:  FormInputTypeRadio,
			Name:  f.Name,
			Value: value,
			Label: &FormLabel{
				Class: "form-check-label",
				Text:  label,
			},
			Checked:  f.Selected == value,
			Required: f.Required,
		}

		if f.Validation != nil {
			validation := FormValidation{
				Valid: f.Validation.Valid,
			}
			if index == lastInput {
				validation.Response = f.Validation.Response
			}
			formInputs[index].Validation = &validation
		}

		index++
	}

	return &formInputs
}
