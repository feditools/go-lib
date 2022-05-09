package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials"
)

// Credential contains feditools grpc token
type Credential struct {
	token string
}

// NewCredential creates a new feditools grpc token
func NewCredential(token string) *Credential {
	return &Credential{
		token: token,
	}
}

// GetRequestMetadata add the token to the request
func (c *Credential) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	ri, _ := credentials.RequestInfoFromContext(ctx)
	fmt.Printf("request info: %+v\n", ri)

	if err := credentials.CheckSecurityLevel(ri.AuthInfo, credentials.IntegrityOnly); err != nil {
		return nil, fmt.Errorf("security %s level too low: %s", ri.AuthInfo, err.Error())
	}

	return map[string]string{
		"authorization": c.token,
	}, nil
}

// RequireTransportSecurity returns false since transport security is not required
func (c *Credential) RequireTransportSecurity() bool {
	return false
}
