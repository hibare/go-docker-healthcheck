package healthcheck

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/hibare/GoCommon/v2/pkg/slice"
)

func Check(url string, statusCodes []int, timeout time.Duration) bool {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Failed to create request", "error", err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to send request", "error", err)
		return false
	}
	defer resp.Body.Close()

	if !slice.SliceContains(resp.StatusCode, statusCodes) {
		slog.Error("Healthcheck failed", "expected", statusCodes, "received", resp.StatusCode)
		return false
	}

	return true
}
