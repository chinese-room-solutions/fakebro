package agent

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestSelectOne(t *testing.T) {
// 	data := []struct {
// 		condition func(a *Agent) bool
// 		expect    *Agent
// 	}{
// 		{
// 			func(a *Agent) bool { return true },
// 			Agents[0],
// 		},
// 		{
// 			func(a *Agent) bool { return false },
// 			nil,
// 		},
// 		{
// 			func(a *Agent) bool { return a.Name == "firefox" },
// 			Agents[0],
// 		},
// 	}

// 	for _, d := range data {
// 		a := SelectOne(d.condition)
// 		require.Equal(t, d.expect, a)
// 	}
// }

// func TestSelectMany(t *testing.T) {
// 	data := []struct {
// 		condition func(a *Agent) bool
// 		expect    []*Agent
// 	}{
// 		{
// 			func(a *Agent) bool { return true },
// 			Agents[:4],
// 		},
// 		{
// 			func(a *Agent) bool { return false },
// 			nil,
// 		},
// 		{
// 			func(a *Agent) bool { return a.Name == "firefox" },
// 			Agents[:4],
// 		},
// 	}

// 	for _, d := range data {
// 		a := SelectMany(d.condition)
// 		require.Equal(t, d.expect, a)
// 	}
// }

// func TestDialTLS(t *testing.T) {
// 	network := "tcp"
// 	remoteResource := "indeed.com:443"

// 	for _, a := range Agents {
// 		conn, err := a.DialTLS(context.Background(), network, remoteResource)
// 		require.NoError(t, err)
// 		require.NotNil(t, conn)
// 	}
// }
