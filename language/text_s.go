package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextSave returns a translated phrase.
func (l *Localizer) TextSave() *LocalizedString {
	lg := logger.WithField("func", "TextSave")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Save",
			Other: "Save",
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

// TextSetting returns a translated phrase.
func (l *Localizer) TextSetting(count int) *LocalizedString {
	lg := logger.WithField("func", "TextSetting")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Setting",
			One:   "Setting",
			Other: "Settings",
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

// TextSystem returns a translated phrase.
func (l *Localizer) TextSystem(count int) *LocalizedString {
	lg := logger.WithField("func", "TextSystem")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "System",
			One:   "System",
			Other: "Systems",
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
