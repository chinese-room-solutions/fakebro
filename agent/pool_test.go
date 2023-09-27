package agent

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPoolRollClose(t *testing.T) {
	roller := NewRoller(nil)
	require.NotNil(t, roller)
	pool := NewAgentPool(1*time.Second, 3, roller)
	require.NotNil(t, pool)
	require.NotNil(t, pool.Agents)
	require.Len(t, pool.Agents, 3)
	pool.Close()
	require.Nil(t, pool.Agents)
}

func TestDo(t *testing.T) {
	network := "tcp"
	remoteResource := "indeed.com:443"
	roller := NewRoller(nil)
	require.NotNil(t, roller)
	pool := NewAgentPool(1*time.Second, 2, roller)
	require.NotNil(t, pool)
	err := pool.Do(func(agent *Agent) (bool, error) {
		conn, err := agent.DialTLS(context.Background(), network, remoteResource)
		require.NoError(t, err)
		require.NotNil(t, conn)
		return true, err
	})
	require.NoError(t, err)
}
