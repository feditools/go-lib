package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextUnauthorized returns a translated phrase.
func (l *Localizer) TextUnauthorized() *LocalizedString {
	lg := logger.WithField("func", "TextUnauthorized")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Unauthorized",
			Other: "Unauthorized",
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

// TextUpdate returns a translated phrase.
func (l *Localizer) TextUpdate(count int) *LocalizedString {
	lg := logger.WithField("func", "TextSystem")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Update",
			One:   "Update",
			Other: "Updates",
		},
		PluralCount: count,
	})
	if err != nil {
		lg.Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}
