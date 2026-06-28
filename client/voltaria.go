package client

import (
	"errors"
	"strings"

	api "github.com/voltariafinance/go-sdk/v2"
	core "github.com/voltariafinance/go-sdk/v2/core"
	option "github.com/voltariafinance/go-sdk/v2/option"
)

// ErrInvalidApiKey is returned by NewVoltaria when the provided API key does
// not carry a recognised environment prefix ("live_" or "sandbox_") and no
// explicit base URL was supplied to override prefix routing.
var ErrInvalidApiKey = errors.New("voltaria: invalid API key: expected a \"live_\" or \"sandbox_\" prefix")

const (
	livePrefix    = "live_"
	sandboxPrefix = "sandbox_"
)

// NewVoltaria constructs a Voltaria SDK client, deriving the API base URL from
// the API key prefix:
//
//   - "live_"    -> production (https://api.voltaria.io)
//   - "sandbox_" -> sandbox    (https://api.sandbox.voltaria.io)
//
// The API key is also used as the bearer token. An explicitly supplied base
// URL option (option.WithBaseURL) always overrides prefix routing: since the
// generated request options are applied in order and the last write wins, the
// derived base URL is prepended so any caller-supplied option takes precedence.
//
// If the key carries no recognised prefix and no explicit base URL was given,
// ErrInvalidApiKey is returned.
func NewVoltaria(apiKey string, opts ...option.RequestOption) (*Client, error) {
	baseURL, ok := baseURLForApiKey(apiKey)

	derived := []option.RequestOption{
		option.WithToken(apiKey),
	}
	if ok {
		// Prepend the derived base URL so a caller-supplied WithBaseURL
		// (applied afterwards) overrides it.
		derived = append(derived, option.WithBaseURL(baseURL))
	} else if !hasExplicitBaseURL(opts) {
		// No recognised prefix and the caller did not override the base URL.
		return nil, ErrInvalidApiKey
	}

	merged := append(derived, opts...)
	return NewClient(merged...), nil
}

// baseURLForApiKey maps an API key prefix to its environment base URL.
func baseURLForApiKey(apiKey string) (string, bool) {
	switch {
	case strings.HasPrefix(apiKey, livePrefix):
		return api.Environments.Production, true
	case strings.HasPrefix(apiKey, sandboxPrefix):
		return api.Environments.Sandbox, true
	default:
		return "", false
	}
}

// hasExplicitBaseURL reports whether the caller supplied an option that sets a
// non-empty base URL, which is allowed to override prefix routing.
func hasExplicitBaseURL(opts []option.RequestOption) bool {
	probe := core.NewRequestOptions(opts...)
	return probe.BaseURL != ""
}
