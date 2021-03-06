package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextBlock returns a translated phrase.
func (l *Localizer) TextBlock(count int) *LocalizedString {
	lg := logger.WithField("func", "TextBlock")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Block",
			One:   "Block",
			Other: "Blocks",
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

// TextBlockExists returns a translated phrase.
func (l *Localizer) TextBlockExists(domain string) *LocalizedString {
	lg := logger.WithField("func", "TextBlockExists")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "BlockExists",
			Other: "Block for domain {{.Domain}} already exists.",
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

// TextBlockSubdomain returns a translated phrase.
func (l *Localizer) TextBlockSubdomain(count int) *LocalizedString {
	lg := logger.WithField("func", "TextBlockSubdomain")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "BlockSubdomain",
			One:   "Block Subdomain",
			Other: "Block Subdomains",
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

// TextBlocked returns a translated phrase.
func (l *Localizer) TextBlocked() *LocalizedString {
	lg := logger.WithField("func", "TextBlocked")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Blocked",
			Other: "Blocked",
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

// TextBlockedDomain returns a translated phrase.
func (l *Localizer) TextBlockedDomain(count int) *LocalizedString {
	lg := logger.WithField("func", "TextBlockedDomain")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "BlockedDomain",
			One:   "Blocked Domain",
			Other: "Blocked Domains",
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
