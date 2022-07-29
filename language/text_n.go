package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextNotification returns a translated phrase.
func (l *Localizer) TextNotification(count int) *LocalizedString {
	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Notification",
			One:   "Notification",
			Other: "Notifications",
		},
		PluralCount: count,
	})
	if err != nil {
		logger.WithField("func", "TextNotification").Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}
