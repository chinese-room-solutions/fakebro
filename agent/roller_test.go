package agent

import (
	"strconv"
	"testing"
	"time"

	"github.com/chinese-room-solutions/fakebro/user_agent"
	"github.com/stretchr/testify/require"
)

func TestRoll(t *testing.T) {
	for _, td := range []struct {
		seed          int64
		expectedAgent *Agent
	}{
		{
			seed: 1695575963768825450,
			expectedAgent: NewAgent(
				"Firefox", "105.0", BaseTLSConfigs[2].Value, map[string]string{
					"User-Agent":                "Mozilla/5.0 (X11; Linux x86_64; rv:105.0) Gecko/20100101 Firefox/105.0",
					"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
					"Accept-Language":           "en-US,en;q=0.5",
					"DNT":                       "1",
					"Upgrade-Insecure-Requests": "1",
					"Connection":                "keep-alive",
					"Sec-Fetch-Dest":            "document",
					"Sec-Fetch-Mode":            "navigate",
					"Sec-Fetch-Site":            "none",
					"Sec-Fetch-User":            "?1",
				}, 1*time.Second,
			),
		},
	} {
		t.Run(strconv.FormatInt(td.seed, 10), func(t *testing.T) {
			r, err := NewRoller(func(t user_agent.TokenType) bool {
				if t > user_agent.StartSafari && t < user_agent.EndSafari ||
					t > user_agent.StartFirefoxMobile && t < user_agent.EndFirefoxMobile ||
					t > user_agent.StartFirefox && t < user_agent.EndFirefox && t != user_agent.Firefox_105_0 {
					return false
				}
				return true
			})
			require.NoError(t, err)
			agent, err := r.Roll(td.seed, 1*time.Second)
			require.NoError(t, err)
			require.Equal(t, td.expectedAgent.Client, agent.Client)
			require.Equal(t, td.expectedAgent.Version, agent.Version)
			require.Equal(t, td.expectedAgent.TLSConfig, agent.TLSConfig)
			require.Equal(t, td.expectedAgent.Headers, agent.Headers)
		})
	}
}
