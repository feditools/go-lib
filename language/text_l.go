package liblanguage

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextLogin returns a translated phrase.
func (l *Localizer) TextLogin() *LocalizedString {
	lg := logger.WithField("func", "TextLogin")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Login",
			Other: "Login",
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

// TextLooksGood returns a translated phrase.
func (l *Localizer) TextLooksGood() *LocalizedString {
	lg := logger.WithField("func", "TextLooksGood")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "LooksGood",
			Other: "Looks Good!",
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
