package template

import liblanguage "github.com/feditools/go-lib/language"

type FormLabel struct {
	Text  *liblanguage.LocalizedString
	Badge *Badge
	Class string
}
