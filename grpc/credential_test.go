package grpc

//revive:disable:add-constant

import (
	"context"
	"testing"
)

func TestNewCredential(t *testing.T) {
	c := NewCredential("test")

	if c == nil {
		t.Errorf("expected credentials. got: nil")

		return
	}
	if c.token != "test" {
		t.Errorf("unexpected token. got: '%s', want '%s'", c.token, "test")
	}
}

func TestCredential_GetRequestMetadata(t *testing.T) {
	c := &Credential{
		token: "test",
	}

	headers, err := c.GetRequestMetadata(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if headers == nil {
		t.Errorf("expected headers. got: nil")

		return
	}
	resp, ok := headers["authorization"]
	if !ok {
		t.Errorf("expected authorization header. got: nil")

		return
	}
	if resp != "test" {
		t.Errorf("unexpected authorization header. got: '%s', want '%s'", resp, "test")
	}
}

func TestCredential_RequireTransportSecurity(t *testing.T) {
	c := &Credential{
		token: "test",
	}

	result := c.RequireTransportSecurity()
	if result {
		t.Errorf("unexpected bool: got: %v, want: %v", result, false)
	}
}

//revive:enable:add-constant
