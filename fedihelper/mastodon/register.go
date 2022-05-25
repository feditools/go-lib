package mastodon

import (
	"context"
	"github.com/feditools/go-lib/fedihelper"
	"net/http"

	mastodon "github.com/mattn/go-mastodon"
)

// RegisterApp registers fedihelper with mastodon and returns the client id and client secret.
func (h *Helper) RegisterApp(ctx context.Context, instance fedihelper.Instance) (clientID string, clientSecret string, err error) {
	l := logger.WithField("func", "RegisterApp")
	var v interface{}
	v, err, _ = h.registerAppGroup.Do(instance.GetDomain(), func() (interface{}, error) {
		instanceToken := h.fedi.GetTokenHandler(ctx, instance)
		app, merr := mastodon.RegisterApp(ctx, &mastodon.AppConfig{
			Client: http.Client{
				Transport: h.fedi.HTTP().Transport(),
			},
			Server:       "https://" + instance.GetServerHostname(),
			ClientName:   h.appClientName,
			Scopes:       "read:accounts",
			Website:      h.appWebsite,
			RedirectURIs: h.externalURL + "/callback/oauth/" + instanceToken,
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

	if err != nil {
		l.Errorf("singleflight: %s", err.Error())

		return
	}

	keys := v.(*[]string)
	clientID = (*keys)[0]
	clientSecret = (*keys)[1]

	return
}
