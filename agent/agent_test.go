package agent

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDialTLS(t *testing.T) {
	network := "tcp"
	remoteResource := "indeed.com:443"
	roller, err := NewRoller(nil)
	require.NoError(t, err)

	for _, td := range []struct {
		seed int64
	}{
		{
			seed: 1695575963768825450,
		},
		{
			seed: 1695249310195281554,
		},
	} {
		t.Run("", func(t *testing.T) {
			agent, err := roller.Roll(td.seed, 1*time.Second)
			require.NoError(t, err)
			require.NotNil(t, agent)
			conn, err := agent.DialTLS(context.Background(), network, remoteResource)
			require.NoError(t, err)
			require.NotNil(t, conn)
		})
	}
}
