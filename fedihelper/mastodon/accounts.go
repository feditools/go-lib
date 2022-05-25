package mastodon

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/feditools/go-lib/fedihelper"
)

// GetCurrentAccount retrieves the current federated account.
func (h *Helper) GetCurrentAccount(ctx context.Context, instance fedihelper.Instance, accessToken string) (fedihelper.Account, error) {
	l := logger.WithField("func", "GetCurrentAccount")

	// create mastodon client
	client, err := h.newClient(instance, accessToken)
	if err != nil {
		l.Errorf("creating client: %s", err.Error())

		return nil, err
	}

	// retrieve current account from
	retrievedAccount, err := client.GetAccountCurrentUser(ctx)
	if err != nil {
		l.Errorf("getting current account: %s", err.Error())

		return nil, err
	}

	// check if account is locked
	if retrievedAccount.Locked {
		return nil, fmt.Errorf("account '@%s@%s' locked", retrievedAccount.Username, instance.GetDomain())
	}

	// check if account is a bot
	if retrievedAccount.Bot {
		return nil, fmt.Errorf("account '@%s@%s' is a bot", retrievedAccount.Username, instance.GetDomain())
	}

	// check if account has moved
	if retrievedAccount.Moved != nil {
		return nil, fmt.Errorf("account '@%s@%s' has moved to '@%s'", retrievedAccount.Username, instance.GetDomain(), retrievedAccount.Moved.Acct)
	}

	// try to retrieve federated account
	account, err := h.fedi.GetAccountHandler(ctx, instance, retrievedAccount.Username)
	if err != nil {
		l.Errorf("db read: %s", err.Error())

		return nil, err
	}
	if account != nil {
		return account, nil
	}

	// do webfinger
	webFinger, err := h.fedi.GetWellknownWebFinger(ctx, retrievedAccount.Username, instance.GetDomain())
	if err != nil {
		l.Debugf("webfinger %s@%s: %s", retrievedAccount.Username, instance.GetDomain(), err.Error())

		return nil, err
	}
	actorURI, err := fedihelper.FindActorURI(webFinger)
	if err != nil {
		l.Debugf("webfinger %s@%s: %s", retrievedAccount.Username, instance.GetDomain(), err.Error())

		return nil, err
	}
	if actorURI == nil {
		msg := fmt.Sprintf("can't find actor uri for %s@%s", retrievedAccount.Username, instance.GetDomain())
		l.Debug(msg)

		return nil, errors.New(msg)
	}

	// create new federated account
	newFediAccount, err := h.fedi.NewAccountHandler(ctx)
	if err != nil {
		l.Warnf("new account: %s", err.Error())

		return nil, err
	}
	newFediAccount.SetActorURI(actorURI.String())
	newFediAccount.SetDisplayName(retrievedAccount.DisplayName)
	newFediAccount.SetInstance(instance)
	newFediAccount.SetLastFinger(time.Now())
	newFediAccount.SetUsername(retrievedAccount.Username)
	err = account.SetAccessToken(accessToken)
	if err != nil {
		l.Errorf("set access token: %s", err.Error())

		return nil, err
	}

	// write new federated account to database
	err = h.fedi.CreateAccountHandler(ctx, newFediAccount)
	if err != nil {
		l.Errorf("db create: %s", err.Error())

		return nil, err
	}

	return newFediAccount, nil
}
