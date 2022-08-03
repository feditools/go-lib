package fedihelper

import (
	"context"
	"time"
)

type Account interface {
	GetActorURI() (actorURI string)
	GetDisplayName() (displayName string)
	GetID() (id int64)
	GetInstance() (instance Instance)
	GetLastInfoUpdate() (lastInfoUpdate time.Time)
	GetUsername() (username string)

	SetActorURI(actorURI string)
	SetAvatar(avatar string)
	SetAvatarStatic(avatarStatic string)
	SetBot(bot bool)
	SetDisplayName(displayName string)
	SetInstance(instance Instance)
	SetLastInfoUpdate(lastInfoUpdate time.Time)
	SetLocked(locked bool)
	SetMoved(account Account)
	SetURL(url string)
	SetUsername(username string)
}

// GenerateAccountFromUsername creates an Account object by querying the apis of the federated instance.
func (f *FediHelper) GenerateAccountFromUsername(ctx context.Context, username string, instance Instance) (Account, error) {
	l := logger.WithField("func", "GenerateFediAccountFromUsername")

	account, err := f.NewAccountHandler(ctx)
	if err != nil {
		l.Errorf("get host meta: %s", err.Error())

		return nil, err
	}

	// get host meta
	hostMeta, err := f.FetchHostMeta(ctx, instance.GetDomain())
	if err != nil {
		l.Errorf("get host meta: %s", err.Error())

		return nil, err
	}
	webfingerURI := hostMeta.WebfingerURI()
	if webfingerURI == "" {
		l.Errorf("host meta missing web finger url")

		return nil, NewError("host meta missing web finger url")
	}

	// get actor uri
	webfinger, err := f.FetchWebFinger(ctx, webfingerURI, username, instance.GetDomain())
	if err != nil {
		fhErr := NewErrorf("get wellknown webfinger: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}
	actorURI, err := webfinger.ActorURI()
	if err != nil {
		fhErr := NewErrorf("find actor url: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}
	if actorURI == nil {
		return nil, NewError("missing actor uri")
	}

	actor, err := f.FetchActor(ctx, actorURI)
	if err != nil {
		l.Errorf("decode json: %s", err.Error())

		return nil, err
	}

	account.SetActorURI(actorURI.String())
	account.SetUsername(actor.PreferredUsername)
	account.SetInstance(instance)
	account.SetDisplayName(actor.Name)

	return account, nil
}
