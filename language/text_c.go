package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextChatID returns a translated phrase.
func (l *Localizer) TextChatID(count int) *LocalizedString {
	lg := logger.WithField("func", "TextChatID")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ChatID",
			One:   "Chat ID",
			Other: "Chat IDs",
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

// TextChatIDOrUsername returns a translated phrase.
func (l *Localizer) TextChatIDOrUsername() *LocalizedString {
	lg := logger.WithField("func", "TextChatIDOrUsername")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ChatIDOrUsername",
			Other: "Chat ID or Username",
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

// TextClient returns a translated phrase.
func (l *Localizer) TextClient(count int) *LocalizedString {
	lg := logger.WithField("func", "TextClient")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Client",
			One:   "Client",
			Other: "Clients",
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

// TextClientID returns a translated phrase.
func (l *Localizer) TextClientID(count int) *LocalizedString {
	lg := logger.WithField("func", "TextClientID")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ClientID",
			One:   "Client ID",
			Other: "Client IDs",
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

// TextClientSecret returns a translated phrase.
func (l *Localizer) TextClientSecret(count int) *LocalizedString {
	lg := logger.WithField("func", "TextClientSecret")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ClientSecret",
			One:   "Client Secret",
			Other: "Client Secrets",
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

// TextClose returns a translated phrase.
func (l *Localizer) TextClose() *LocalizedString {
	lg := logger.WithField("func", "TextClose")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Close",
			Other: "Close",
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

// TextConfig returns a translated phrase.
func (l *Localizer) TextConfig(count int) *LocalizedString {
	lg := logger.WithField("func", "TextConfig")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Config",
			One:   "Config",
			Other: "Configs",
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

// TextCreate returns a translated phrase.
func (l *Localizer) TextCreate() *LocalizedString {
	lg := logger.WithField("func", "TextCreate")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Create",
			Other: "Create",
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
