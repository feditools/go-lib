package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextDashboard returns a translated phrase.
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
		lg.Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}

// TextDelete returns a translated phrase.
func (l *Localizer) TextDelete() *LocalizedString {
	lg := logger.WithField("func", "TextDelete")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Delete",
			Other: "Delete",
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

// TextDeleteBlockDomain returns a translated phrase.
func (l *Localizer) TextDeleteBlockDomain(domain string) *LocalizedString {
	lg := logger.WithField("func", "TextDeleteBlockDomain")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DeleteBlockDomain",
			Other: "Delete Block {{.Domain}}",
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

// TextDeleteBlockConfirmDomain returns a translated phrase.
func (l *Localizer) TextDeleteBlockConfirmDomain(domain string) *LocalizedString {
	lg := logger.WithField("func", "TextDeleteBlockConfirmDomain")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DeleteBlockConfirmDomain",
			Other: "Are you sure you want to delete the block for {{.Domain}}?",
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

// TextDemocrablock returns a translated phrase.
func (l *Localizer) TextDemocrablock() *LocalizedString {
	lg := logger.WithField("func", "TextDemocrablock")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Democrablock",
			Other: "Democrablock",
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

// TextDescription returns a translated phrase.
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
		lg.Warningf(missingTranslationWarning, err.Error())
	}

	return &LocalizedString{
		language: tag,
		string:   text,
	}
}

// TextDomain returns a translated phrase.
func (l *Localizer) TextDomain(count int) *LocalizedString {
	lg := logger.WithField("func", "TextDomain")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Domain",
			One:   "Domain",
			Other: "Domains",
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
