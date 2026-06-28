package client

import (
	"errors"
	"testing"

	api "github.com/voltariafinance/go-sdk/v2"
	option "github.com/voltariafinance/go-sdk/v2/option"
)

func TestNewVoltaria_LiveRoutesToProduction(t *testing.T) {
	c, err := NewVoltaria("live_abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got, want := c.baseURL, api.Environments.Production; got != want {
		t.Fatalf("baseURL = %q, want %q", got, want)
	}
}

func TestNewVoltaria_SandboxRoutesToSandbox(t *testing.T) {
	c, err := NewVoltaria("sandbox_abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got, want := c.baseURL, api.Environments.Sandbox; got != want {
		t.Fatalf("baseURL = %q, want %q", got, want)
	}
}

func TestNewVoltaria_UnknownPrefixReturnsErrInvalidApiKey(t *testing.T) {
	c, err := NewVoltaria("nope_abc123")
	if !errors.Is(err, ErrInvalidApiKey) {
		t.Fatalf("err = %v, want ErrInvalidApiKey", err)
	}
	if c != nil {
		t.Fatalf("client = %v, want nil", c)
	}
}

func TestNewVoltaria_EmptyKeyReturnsErrInvalidApiKey(t *testing.T) {
	c, err := NewVoltaria("")
	if !errors.Is(err, ErrInvalidApiKey) {
		t.Fatalf("err = %v, want ErrInvalidApiKey", err)
	}
	if c != nil {
		t.Fatalf("client = %v, want nil", c)
	}
}

func TestNewVoltaria_ExplicitBaseURLOverridesPrefix(t *testing.T) {
	const custom = "https://api.staging.voltaria.io"
	c, err := NewVoltaria("live_abc123", option.WithBaseURL(custom))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got, want := c.baseURL, custom; got != want {
		t.Fatalf("baseURL = %q, want %q (explicit option must win)", got, want)
	}
}

func TestNewVoltaria_ExplicitBaseURLRescuesUnknownPrefix(t *testing.T) {
	const custom = "https://api.staging.voltaria.io"
	c, err := NewVoltaria("nope_abc123", option.WithBaseURL(custom))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got, want := c.baseURL, custom; got != want {
		t.Fatalf("baseURL = %q, want %q", got, want)
	}
}
