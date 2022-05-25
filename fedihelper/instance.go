package fedihelper

import (
	"context"
	"errors"
)

type Instance interface {
	IsOAuthSet() bool

	GetActorURI() (actorURI string)
	GetClientID() (clientID string)
	GetClientSecret() (clientID string, err error)
	GetDomain() (domain string)
	GetServerHostname() (hostname string)
	GetSoftware() (software string)

	SetActorURI(actorURI string)
	SetClientID(clientID string)
	SetClientSecret(clientID string) (err error)
	SetDomain(domain string)
	SetServerHostname(hostname string)
	SetSoftware(software string)
}

// GenerateFediInstanceFromDomain created a Instance object by querying the apis of the federated instance.
func (f *FediHelper) GenerateFediInstanceFromDomain(ctx context.Context, domain string, instance Instance) error {
	l := logger.WithField("func", "GenerateFediInstanceFromDomain")

	// get nodeinfo endpoints from well-known location
	wkni, err := f.GetWellknownNodeInfo(ctx, domain)
	if err != nil {
		l.Errorf("get nodeinfo: %s", err.Error())

		return err
	}

	// check for nodeinfo 2.0 schema
	nodeinfoURI, err := findNodeInfo20URI(wkni)
	if err != nil {
		return err
	}
	if nodeinfoURI == nil {
		return errors.New("missing nodeinfo 2.0 uri")
	}

	// get nodeinfo from
	nodeinfo, err := f.GetNodeInfo20(ctx, domain, nodeinfoURI)
	if err != nil {
		l.Errorf("get nodeinfo 2.0: %s", err.Error())

		return err
	}

	// get actor uri
	webfinger, err := f.GetWellknownWebFinger(ctx, domain, domain)
	if err != nil {
		return err
	}
	actorURI, err := FindActorURI(webfinger)
	if err != nil {
		return err
	}
	if actorURI == nil {
		return errors.New("missing actor uri")
	}

	instance.SetActorURI(actorURI.String())
	instance.SetDomain(domain)
	instance.SetServerHostname(nodeinfoURI.Host)
	instance.SetSoftware(nodeinfo.Software.Name)

	return nil
}