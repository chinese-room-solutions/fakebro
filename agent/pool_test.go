package agent

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// )

// // func TestRoll(t *testing.T) {
// // 	p := ActiveAgentPool{
// // 		Condition: func(a *Agent) bool {
// // 			return a.Name == "firefox" && a.Version == 55
// // 		},
// // 		Limit: 1,
// // 	}
// // 	p.Roll()
// // 	require.Len(t, p.Agents, 1)
// // 	require.NotNil(t, p.ActiveAgents)
// // }

// func TestGet(t *testing.T) {
// 	p := &ActiveAgentPool{
// 		Condition: func(a *Agent) bool {
// 			return a.Name == "firefox" && a.Version == 55
// 		},
// 		Limit: 1,
// 	}
// 	_, err := p.Get()
// 	require.Error(t, err)
// 	require.Equal(t, ErrNilPool, err)

// 	p = NewActiveAgentPool(
// 		"tcp", "indeed.com:443", 200*time.Millisecond, 2,
// 		func(a *Agent) bool {
// 			return a.Name == "firefox"
// 		},
// 	)
// 	agent, err := p.Get()
// 	require.NoError(t, err)
// 	require.NotNil(t, agent)
// }

// func TestPut(t *testing.T) {
// 	p := NewActiveAgentPool(
// 		"tcp", "indeed.com:443", 200*time.Millisecond, 2,
// 		func(a *Agent) bool {
// 			return a.Name == "firefox" && a.Version == 55
// 		},
// 	)
// 	aa := NewActiveAgent(Agents[0])
// 	err := p.Put(aa)
// 	require.NoError(t, err)
// 	require.Len(t, p.ActiveAgents, 1)
// 	require.NotNil(t, p.ActiveAgents)
// }

// func TestClose(t *testing.T) {
// 	p := NewActiveAgentPool(
// 		"tcp", "indeed.com:443", 200*time.Millisecond, 2,
// 		func(a *Agent) bool {
// 			return a.Name == "firefox" && a.Version == 55
// 		},
// 	)
// 	p.Close()
// 	require.Nil(t, p.ActiveAgents)
// 	require.Equal(t, p.Index, 0)
// }
