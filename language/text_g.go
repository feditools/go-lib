package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextGeneral returns a translated phrase.
func (l *Localizer) TextGeneral() *LocalizedString {
	lg := logger.WithField("func", "TextGeneral")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "General",
			Other: "General",
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
