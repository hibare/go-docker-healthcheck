package healthcheck

import (
	"net/http"
	"time"

	"github.com/hibare/GoCommon/v2/pkg/slice"
	"github.com/rs/zerolog/log"
)

func Check(url string, statusCodes []int, timeout time.Duration) bool {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error().Err(err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err)
		return false
	}
	defer resp.Body.Close()

	if !slice.SliceContains(resp.StatusCode, statusCodes) {
		log.Error().Msgf("Expected status %v but received %d", statusCodes, resp.StatusCode)
		return false
	}

	return true
}
