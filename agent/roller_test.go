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
					"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:105.0) Gecko/20100101 Firefox/105.0",
				}, 1*time.Second,
			),
		},
		{
			seed: 1693588721744517187,
			expectedAgent: NewAgent(
				"Safari", "16.5", BaseTLSConfigs[3].Value, map[string]string{
					"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5 Safari/605.1.15",
				}, 1*time.Second,
			),
		},
	} {
		t.Run(strconv.FormatInt(td.seed, 10), func(t *testing.T) {
			r := NewRoller(func(t user_agent.TokenType) bool {
				if t > user_agent.StartSafari && t < user_agent.EndSafari && t != user_agent.Safari_16_5 ||
					t > user_agent.StartFirefoxMobile && t < user_agent.EndFirefoxMobile ||
					t > user_agent.StartFirefox && t < user_agent.EndFirefox && t != user_agent.Firefox_105_0 {
					return false
				}
				return true
			})
			require.NotNil(t, r)
			agent := r.Roll(td.seed, 1*time.Second)
			require.NotNil(t, agent)
			require.Equal(t, td.expectedAgent.Client, agent.Client)
			require.Equal(t, td.expectedAgent.Version, agent.Version)
			require.Equal(t, td.expectedAgent.TLSConfig, agent.TLSConfig)
			require.Equal(t, td.expectedAgent.Headers["User-Agent"], agent.Headers["User-Agent"])
		})
	}
}
