package mastodon

import (
	"context"
	"github.com/feditools/go-lib"
	"github.com/mattn/go-mastodon"
	"time"

	"github.com/feditools/go-lib/fedihelper"
)

// GetCurrentAccount retrieves the current federated account.
func (h *Helper) GetCurrentAccount(ctx context.Context, instance fedihelper.Instance, accessToken string) (fedihelper.Account, error) {
	l := logger.WithField("func", "GetCurrentAccount")

	// create mastodon client
	client, err := h.newClient(ctx, instance, accessToken)
	if err != nil {
		fhErr := fedihelper.NewErrorf("find actor url: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	// retrieve current account from
	retrievedAccount, err := client.GetAccountCurrentUser(ctx)
	if err != nil {
		fhErr := fedihelper.NewErrorf("getting current account: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	// try to retrieve federated account
	account, found, err := h.fedi.GetAccountHandler(ctx, instance, retrievedAccount.Username)
	if err != nil {
		fhErr := fedihelper.NewErrorf("get account: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}
	if found {
		return account, nil
	}

	// do webfinger
	webFinger, err := h.fedi.FetchWellknownWebFinger(ctx, instance.GetServerHostname(), retrievedAccount.Username, instance.GetDomain())
	if err != nil {
		fhErr := fedihelper.NewErrorf("webfinger %s@%s: %s", retrievedAccount.Username, instance.GetDomain(), err.Error())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}
	actorURI, err := webFinger.ActorURI()
	if err != nil {
		fhErr := fedihelper.NewErrorf("finding actor uri %s@%s: %s", retrievedAccount.Username, instance.GetDomain(), err.Error())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}
	if actorURI == nil {
		fhErr := fedihelper.NewErrorf("didn't find actor uri for %s@%s", retrievedAccount.Username, instance.GetDomain())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}

	// create new federated account
	newAccount, err := h.fedi.NewAccountHandler(ctx)
	if err != nil {
		fhErr := fedihelper.NewErrorf("new account: %s", err.Error())
		l.Warn(fhErr.Error())

		return nil, fhErr
	}
	if retrievedAccount.Moved == nil {
		movedAccount, err := h.movedAccountHelper(ctx, retrievedAccount.Moved)
		if err != nil {
			fhErr := fedihelper.NewErrorf("new account: %s", err.Error())
			l.Warn(fhErr.Error())

			return nil, fhErr
		}
		newAccount.SetMoved(movedAccount)
	}

	newAccount.SetActorURI(actorURI.String())
	newAccount.SetInstance(instance)
	newAccount.SetLastInfoUpdate(time.Now())
	populateAccount(newAccount, retrievedAccount)

	// write new federated account to database
	err = h.fedi.CreateAccountHandler(ctx, newAccount)
	if err != nil {
		fhErr := fedihelper.NewErrorf("db create: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	err = h.kv.SetAccessToken(ctx, newAccount.GetID(), accessToken)
	if err != nil {
		fhErr := fedihelper.NewErrorf("set access token: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	return newAccount, nil
}

func (h *Helper) movedAccountHelper(ctx context.Context, mastodonAccount *mastodon.Account) (fedihelper.Account, error) {
	l := logger.WithField("func", "movedAccountHelper")

	_, domain, err := lib.SplitAccount(mastodonAccount.Acct)
	if err != nil {
		return nil, err
	}

	instance, err := h.fedi.GetOrCreateInstance(ctx, domain)
	if err != nil {
		return nil, err
	}

	// try to retrieve account
	account, found, err := h.fedi.GetAccountHandler(ctx, instance, mastodonAccount.Username)
	if err != nil {
		fhErr := fedihelper.NewErrorf("get account: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}
	if found {
		return account, nil
	}

	// not found create new account
	newAccount, err := h.fedi.NewAccountHandler(ctx)
	if err != nil {
		fhErr := fedihelper.NewErrorf("new account: %s", err.Error())
		l.Warn(fhErr.Error())

		return nil, fhErr
	}

	// do webfinger
	webFinger, err := h.fedi.FetchWellknownWebFinger(ctx, instance.GetServerHostname(), mastodonAccount.Username, instance.GetDomain())
	if err != nil {
		fhErr := fedihelper.NewErrorf("webfinger %s@%s: %s", mastodonAccount.Username, instance.GetDomain(), err.Error())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}
	actorURI, err := webFinger.ActorURI()
	if err != nil {
		fhErr := fedihelper.NewErrorf("finding actor uri %s@%s: %s", mastodonAccount.Username, instance.GetDomain(), err.Error())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}
	if actorURI == nil {
		fhErr := fedihelper.NewErrorf("didn't find actor uri for %s@%s", mastodonAccount.Username, instance.GetDomain())
		l.Debug(fhErr.Error())

		return nil, fhErr
	}

	if mastodonAccount.Moved == nil {
		movedAccount, err := h.movedAccountHelper(ctx, mastodonAccount.Moved)
		if err != nil {
			fhErr := fedihelper.NewErrorf("new account: %s", err.Error())
			l.Warn(fhErr.Error())

			return nil, fhErr
		}
		newAccount.SetMoved(movedAccount)
	}

	// create new account
	newAccount.SetActorURI(actorURI.String())
	newAccount.SetInstance(instance)
	newAccount.SetLastInfoUpdate(time.Now())
	populateAccount(newAccount, mastodonAccount)

	// write new federated account to database
	err = h.fedi.CreateAccountHandler(ctx, newAccount)
	if err != nil {
		fhErr := fedihelper.NewErrorf("db create: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	return newAccount, nil
}

func populateAccount(fediAccount fedihelper.Account, mastodonAccount *mastodon.Account) {
	fediAccount.SetAvatar(mastodonAccount.Avatar)
	fediAccount.SetAvatarStatic(mastodonAccount.AvatarStatic)
	fediAccount.SetBot(mastodonAccount.Bot)
	fediAccount.SetDisplayName(mastodonAccount.DisplayName)
	fediAccount.SetLocked(mastodonAccount.Locked)
	fediAccount.SetURL(mastodonAccount.URL)
	fediAccount.SetUsername(mastodonAccount.Username)
}
