package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextHomeWeb returns a translated phrase.
func (l *Localizer) TextHomeWeb() *LocalizedString {
	lg := logger.WithField("func", "TextHomeWeb")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "HomeWeb",
			Other: "Home",
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

// TextHomePageBody returns a translated phrase.
func (l *Localizer) TextHomePageBody() *LocalizedString {
	lg := logger.WithField("func", "TextHomeWeb")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "HomePageBody",
			Other: "Home Page Body",
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

// TextHowToJoin returns a translated phrase.
func (l *Localizer) TextHowToJoin() *LocalizedString {
	lg := logger.WithField("func", "TextHowToJoin")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "HowToJoin",
			Other: "How to Join",
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
