package mastodon

import (
	"github.com/feditools/go-lib/log"
)

type empty struct{}

var logger = log.WithPackageField(empty{})
