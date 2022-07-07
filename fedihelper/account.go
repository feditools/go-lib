package fedihelper

import "time"

type Account interface {
	GetActorURI() (actorURI string)
	GetDisplayName() (displayName string)
	GetInstance() (instance Instance)
	GetLastFinger() (lastFinger time.Time)
	GetUsername() (username string)

	SetActorURI(actorURI string)
	SetDisplayName(displayName string)
	SetInstance(instance Instance)
	SetLastFinger(lastFinger time.Time)
	SetUsername(username string)
}
