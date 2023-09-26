package agent

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPoolRollClose(t *testing.T) {
	r, err := NewRoller(nil)
	require.NoError(t, err)
	p, err := NewAgentPool(1*time.Second, 3, r)
	require.NoError(t, err)
	require.NotNil(t, p.Agents)
	require.Len(t, p.Agents, 3)
	p.Close()
	require.Nil(t, p.Agents)
}

func TestDo(t *testing.T) {
	network := "tcp"
	remoteResource := "indeed.com:443"
	roller, err := NewRoller(nil)
	require.NoError(t, err)
	p, err := NewAgentPool(1*time.Second, 3, roller)
	require.NoError(t, err)
	err = p.Do(func(agent *Agent) error {
		conn, err := agent.DialTLS(context.Background(), network, remoteResource)
		require.NoError(t, err)
		require.NotNil(t, conn)
		return ErrBadAgent
	})
	require.NoError(t, err)
}
