// Package constants provides default values for the healthcheck package.
package constants

import (
	"net/http"
	"time"
)

var (
	// DefaultSuccessStatusCodes is the default status codes for a successful healthcheck.
	DefaultSuccessStatusCodes = []int{http.StatusOK}
)

const (
	// DefaultTimeout is the default timeout for the healthcheck.
	DefaultTimeout = 10 * time.Second
)
