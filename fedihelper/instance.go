package fedihelper

import (
	"context"
	"net/url"
)

type Instance interface {
	GetActorURI() (actorURI string)
	GetOAuthClientID() (clientID string)
	GetOAuthClientSecret() (clientSecret string, err error)
	GetDomain() (domain string)
	GetID() (id int64)
	GetServerHostname() (hostname string)
	GetSoftware() (software string)

	SetActorURI(actorURI string)
	SetOAuthClientID(clientID string)
	SetOAuthClientSecret(clientSecret string) error
	SetDomain(domain string)
	SetInboxURI(inboxURI string)
	SetServerHostname(hostname string)
	SetSoftware(software string)
}

// GenerateInstanceFromDomain created an Instance object by querying the apis of the federated instance.
func (f *FediHelper) GenerateInstanceFromDomain(ctx context.Context, domain string) (Instance, error) {
	l := logger.WithField("func", "GenerateFediInstanceFromDomain")

	// get host meta
	hostMeta, err := f.FetchHostMeta(ctx, domain)
	if err != nil {
		l.Errorf("get host meta: %s", err.Error())

		return nil, err
	}
	hostMetaURIString, err := f.WebfingerURIFromHostMeta(hostMeta)
	if err != nil {
		l.Errorf("get webfinger uri: %s", err.Error())

		return nil, err
	}
	hostMetaURI, err := url.Parse(hostMetaURIString)
	if err != nil {
		l.Errorf("parsing host meta uri: %s", err.Error())

		return nil, err
	}

	// get nodeinfo endpoints from well-known location
	wkni, err := f.GetWellknownNodeInfo(ctx, hostMetaURI.Host)
	if err != nil {
		l.Errorf("get nodeinfo: %s", err.Error())

		return nil, err
	}

	// check for nodeinfo 2.0 schema
	nodeinfoURI, err := findNodeInfo20URI(wkni)
	if err != nil {
		return nil, err
	}
	if nodeinfoURI == nil {
		return nil, NewError("missing nodeinfo 2.0 uri")
	}

	// get nodeinfo from
	nodeinfo, err := f.GetNodeInfo20(ctx, hostMetaURI.Host, nodeinfoURI)
	if err != nil {
		fhErr := NewErrorf("get nodeinfo 2.0: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	// get actor uri
	webfinger, err := f.FetchWellknownWebFinger(ctx, hostMetaURI.Host, domain, domain)
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
		fhErr := NewErrorf("can't fetch actor: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	instance, err := f.NewInstanceHandler(ctx)
	if err != nil {
		fhErr := NewErrorf("new instance: %s", err.Error())
		l.Error(fhErr.Error())

		return nil, fhErr
	}

	instance.SetActorURI(actorURI.String())
	instance.SetInboxURI(actor.Endpoints.SharedInbox)
	instance.SetDomain(domain)
	instance.SetServerHostname(hostMetaURI.Host)
	instance.SetSoftware(nodeinfo.Software.Name)

	return instance, nil
}

func (f *FediHelper) GetOrCreateInstance(ctx context.Context, domain string) (Instance, error) {
	// check if instance exists
	instance, found, err := f.GetInstanceHandler(ctx, domain)
	if err != nil {
		return nil, err
	}
	if found {
		return instance, err
	}

	// create new instance
	newInstance, err := f.GenerateInstanceFromDomain(ctx, domain)
	if err != nil {
		return nil, err
	}

	// save instance to database
	err = f.CreateInstanceHandler(ctx, newInstance)
	if err != nil {
		return nil, err
	}

	return newInstance, nil
}
