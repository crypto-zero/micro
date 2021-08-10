// Package handler is the handler for the `micro network dns` command
package handler

import (
	"github.com/crypto-zero/micro/v2/service/network/dns/provider"
)

// New returns a new handler
func New(provider provider.Provider, token string) *DNS {
	return &DNS{
		provider:    provider,
		bearerToken: token,
	}
}
