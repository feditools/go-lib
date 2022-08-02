package fedihelper

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func (f *FediHelper) HandleCallback(r *http.Request, instance Instance, callbackURI *url.URL) (Account, int, error) {
	l := logger.WithField("func", "HandleCallback")

	switch SoftwareName(instance.GetSoftware()) {
	case SoftwareMastodon:
		// get code
		var code []string
		var ok bool
		if code, ok = r.URL.Query()["code"]; !ok || len(code[0]) < 1 {
			l.Debugf("missing code")

			return nil, http.StatusBadRequest, errors.New("missing code")
		}

		// retrieve access token
		var accessToken string
		accessToken, err := f.Helper(SoftwareMastodon).GetAccessToken(
			r.Context(),
			callbackURI,
			instance,
			code[0],
		)
		if err != nil {
			l.Errorf("get access token error: %s", err.Error())

			return nil, http.StatusInternalServerError, err
		}
		l.Debugf("access token: %s", accessToken)

		// retrieve current account
		account, err := f.Helper(SoftwareMastodon).GetCurrentAccount(
			r.Context(),
			instance,
			accessToken,
		)
		if err != nil {
			l.Errorf("get access token error: %s", err.Error())

			return nil, http.StatusInternalServerError, err
		}

		return account, 0, nil
	default:
		return nil, http.StatusNotImplemented, fmt.Errorf("no helper for '%s'", instance.GetSoftware())
	}
}
