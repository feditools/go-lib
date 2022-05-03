package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextDashboard returns a translated phrase
func (l *Localizer) TextDashboard(count int) *LocalizedString {
	lg := logger.WithField("func", "TextDashboard")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Dashboard",
			One:   "Dashboard",
			Other: "Dashboards",
		},
		PluralCount: count,
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return &LocalizedString{
		language: tag,
		string:   text,
	}
}

// TextDescription returns a translated phrase
func (l *Localizer) TextDescription(count int) *LocalizedString {
	lg := logger.WithField("func", "TextDescription")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Description",
			One:   "Description",
			Other: "Descriptions",
		},
		PluralCount: count,
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return &LocalizedString{
		language: tag,
		string:   text,
	}
}
