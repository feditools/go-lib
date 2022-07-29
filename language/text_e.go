package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextEditBlockDomain returns a translated phrase.
func (l *Localizer) TextEditBlockDomain(domain string) *LocalizedString {
	lg := logger.WithField("func", "TextEditBlockDomain")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "EditBlockDomain",
			Other: "Edit Block {{.Domain}}",
		},
		TemplateData: map[string]interface{}{
			"Domain": domain,
		},
	})
	if err != nil {
		lg.Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}

// TextErrorDatabase returns a translated phrase.
func (l *Localizer) TextErrorDatabase() *LocalizedString {
	lg := logger.WithField("func", "TextErrorDatabase")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ErrorDatabase",
			Other: "database error",
		},
	})
	if err != nil {
		lg.Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}
