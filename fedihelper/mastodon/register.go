package mastodon

import (
	"context"
	"net/http"
	"net/url"

	"github.com/feditools/go-lib/fedihelper"

	mastodon "github.com/mattn/go-mastodon"
)

// RegisterApp registers fedihelper with mastodon and returns the client id and client secret.
func (h *Helper) RegisterApp(ctx context.Context, redirectURI *url.URL, instance fedihelper.Instance) (clientID string, clientSecret string, err error) {
	l := logger.WithField("func", "RegisterApp")
	v, serr, _ := h.registerAppGroup.Do(instance.GetDomain(), func() (interface{}, error) {
		app, merr := mastodon.RegisterApp(ctx, &mastodon.AppConfig{
			Client: http.Client{
				Transport: h.transport.Client().Transport(),
			},
			Server:       "https://" + instance.GetServerHostname(),
			ClientName:   h.appClientName,
			Scopes:       "read:accounts",
			Website:      h.appWebsite,
			RedirectURIs: redirectURI.String(),
		})

		if merr != nil {
			l.Errorf("registering app: %s", err.Error())

			return nil, merr
		}

		keys := []string{
			app.ClientID,
			app.ClientSecret,
		}

		return &keys, nil
	})

	if serr != nil {
		l.Errorf("singleflight: %s", err.Error())

		return "", "", serr
	}
	keys, ok := v.(*[]string)
	if !ok {
		return "", "", fedihelper.NewError("invalid response type from single flight")
	}

	return (*keys)[0], (*keys)[1], nil
}
