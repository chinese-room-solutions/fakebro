package user_agent

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUserAgent(t *testing.T) {
	for _, td := range []struct {
		seed            int64
		expectHeader    string
		expectClient    string
		expectedVersion string
	}{
		{
			seed:            1695575963768825450,
			expectHeader:    "Mozilla/5.0 (X11; Linux x86_64; rv:105.0) Gecko/20100101 Firefox/105.0",
			expectClient:    "Firefox",
			expectedVersion: "105.0",
		},
		{
			seed:            1695290468933462166,
			expectHeader:    "Mozilla/5.0 (iPad; CPU OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/99.0 Mobile/15E148 Safari/605.1.15",
			expectClient:    "FxiOS",
			expectedVersion: "99.0",
		},
		{
			seed:            1695290138144165281,
			expectHeader:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6.1 Safari/605.1.15",
			expectClient:    "Safari",
			expectedVersion: "15.6.1",
		},
		{
			seed:            1695288692615889799,
			expectHeader:    "Mozilla/5.0 (iPad; CPU OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/105.0 Mobile/15E148 Safari/605.1.15",
			expectClient:    "FxiOS",
			expectedVersion: "105.0",
		},
		{
			seed:            1695249310195281554,
			expectHeader:    "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5.2 Mobile/15E148 Safari/605.1.15",
			expectClient:    "Safari",
			expectedVersion: "16.5.2",
		},
		{
			seed:            1695247073946548752,
			expectHeader:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.3) rv:102.0) Gecko/20100101 Firefox/102.0",
			expectClient:    "Firefox",
			expectedVersion: "102.0",
		},
		{
			seed:            1695246322582242447,
			expectHeader:    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0",
			expectClient:    "Firefox",
			expectedVersion: "102.0",
		},
		{
			seed:            1693587396081377813,
			expectHeader:    "Mozilla/5.0 (Android 11; Mobile SM-G973U; rv:102.0) Gecko/102.0 Firefox/102.0",
			expectClient:    "Firefox",
			expectedVersion: "102.0",
		},
		{
			seed:            1693588302512633204,
			expectHeader:    "Mozilla/5.0 (Android 12; Mobile SM-A536B; rv:102.0) Gecko/102.0 Firefox/102.0",
			expectClient:    "Firefox",
			expectedVersion: "102.0",
		},
		{
			seed:            1693588721744517187,
			expectHeader:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6.1 Safari/605.1.15",
			expectClient:    "Safari",
			expectedVersion: "15.6.1",
		},
	} {
		t.Run(strconv.FormatInt(td.seed, 10), func(t *testing.T) {
			ua := NewUserAgent(20, td.seed)
			require.NotNil(t, ua)
			require.Equal(t, td.expectHeader, ua.Header)
			require.Equal(t, td.expectClient, ua.Client)
		})
	}
}
