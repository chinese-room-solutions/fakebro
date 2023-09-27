package roundtripper

import (
	"net/http"
	"testing"
	"time"

	"github.com/chinese-room-solutions/fakebro/user_agent"
	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	client := &http.Client{
		Transport: New(
			1*time.Second,
			2,
			func(user_agent.TokenType) bool { return true },
			map[string]string{},
			func(error) bool { return false },
		),
	}

	resp, err := client.Get("https://nl.indeed.com")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}
