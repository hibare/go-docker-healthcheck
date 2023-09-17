package healthcheck

import (
	"testing"
	"time"
)

func TestCheckFunction(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		statusCodes []int
		timeout     time.Duration
		expected    bool
	}{
		{
			name:        "Valid URL and status code",
			url:         "https://example.com",
			statusCodes: []int{200},
			timeout:     5 * time.Second,
			expected:    true,
		},
		{
			name:        "Invalid URL",
			url:         "https://nonexistenturl12345.com",
			statusCodes: []int{200},
			timeout:     5 * time.Second,
			expected:    false,
		},
		{
			name:        "Invalid Status Code",
			url:         "https://example.com",
			statusCodes: []int{404},
			timeout:     5 * time.Second,
			expected:    false,
		},
		{
			name:        "Timeout",
			url:         "https://httpbin.org/delay/10", // Deliberately slow endpoint
			statusCodes: []int{200},
			timeout:     1 * time.Second, // Short timeout to force a timeout
			expected:    false,
		},
		{
			name:        "Timeout with Fast URL",
			url:         "https://example.com",
			statusCodes: []int{200},
			timeout:     1 * time.Millisecond, // Extremely short timeout
			expected:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Check(test.url, test.statusCodes, test.timeout)

			if result != test.expected {
				t.Errorf("Check(%s, %v, %v) returned %v, expected %v",
					test.url, test.statusCodes, test.timeout, result, test.expected)
			}
		})
	}
}
