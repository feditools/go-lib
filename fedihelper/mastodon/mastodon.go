package mastodon

import (
	"context"

	"github.com/feditools/go-lib/fedihelper"
	mastodon "github.com/mattn/go-mastodon"
	"golang.org/x/sync/singleflight"
)

// Helper is a mastodon helper.
type Helper struct {
	fedi      *fedihelper.FediHelper
	kv        fedihelper.KV
	transport *fedihelper.Transport

	appClientName string
	appWebsite    string
	externalURL   string

	registerAppGroup singleflight.Group
}

// New returns a new mastodon helper.
func New(k fedihelper.KV, t *fedihelper.Transport, appClientName, appWebsite, externalURL string) (*Helper, error) {
	return &Helper{
		kv:        k,
		transport: t,

		appClientName: appClientName,
		appWebsite:    appWebsite,
		externalURL:   externalURL,
	}, nil
}

// newClient return new mastodon API client.
func (h *Helper) newClient(ctx context.Context, instance fedihelper.Instance, accessToken string) (*mastodon.Client, error) {
	l := logger.WithField("func", "newClient")

	// get client secret
	clientSecret, err := instance.GetOAuthClientSecret()
	if err != nil {
		l.Errorf("get client secret: %s", err.Error())

		return nil, err
	}

	// create client
	client := mastodon.NewClient(&mastodon.Config{
		Server:       "https://" + instance.GetServerHostname(),
		ClientID:     instance.GetOAuthClientID(),
		ClientSecret: clientSecret,
		AccessToken:  accessToken,
	})

	// apply custom transport
	client.Transport = h.transport.Client().Transport()

	return client, nil
}

// GetSoftware returns the software type of this module.
func (*Helper) GetSoftware() fedihelper.SoftwareName { return fedihelper.SoftwareMastodon }

// SetFedi adds the fedi module to a helper.
func (h *Helper) SetFedi(f *fedihelper.FediHelper) {
	h.fedi = f
}
