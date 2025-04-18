// Package healthcheck provides a simple health check for a given URL.
package healthcheck

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/hibare/GoCommon/v2/pkg/slice"
)

// Check performs a health check on the given URL.
// It returns true if the response status code is in the list of expected status codes.
// It returns false if the request fails or if the response status code is not in the list of expected status codes.
// The function takes a URL, a list of expected status codes, and a timeout duration as parameters.
// The default timeout is 10 seconds.
func Check(url string, statusCodes []int, timeout time.Duration) bool {
	ctx := context.Background()

	client := http.Client{
		Timeout: timeout,
	}

	req, cErr := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if cErr != nil {
		slog.Error("Failed to create request", "error", cErr)
		return false
	}

	resp, cErr := client.Do(req)
	if cErr != nil {
		slog.Error("Failed to send request", "error", cErr)
		return false
	}
	defer func() {
		if bErr := resp.Body.Close(); bErr != nil {
			slog.Error("Failed to close response body", "error", bErr)
		}
	}()

	if !slice.SliceContains(resp.StatusCode, statusCodes) {
		slog.Error("Healthcheck failed", "expected", statusCodes, "received", resp.StatusCode)
		return false
	}

	return true
}
