package roundtripper

// import (
// 	"net/http"
// 	"time"

// 	"github.com/chinese-room-solutions/fakebro/agent"
// 	"golang.org/x/exp/maps"
// )

// type RoundTripper struct {
// 	Pool    *agent.ActiveAgentPool
// 	Headers map[string]string
// }

// func New(
// 	network string,
// 	address string,
// 	dialTimeout time.Duration,
// 	limit int,
// 	condition func(*agent.Agent) bool,
// 	headers map[string]string,
// ) *RoundTripper {
// 	return &RoundTripper{
// 		Pool:    agent.NewActiveAgentPool(network, address, dialTimeout, limit, condition),
// 		Headers: headers,
// 	}
// }

// func (rt *RoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
// 	// Get an agent from the pool
// 	a, err := rt.Pool.Get()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Merge the custom headers with the agent headers with the precedence of the custom headers
// 	headers := map[string]string{}
// 	maps.Copy(headers, a.Headers)
// 	maps.Copy(headers, rt.Headers)

// 	for header, value := range headers {
// 		r.Header.Set(header, value)
// 	}

// 	return a.T.RoundTrip(r)
// }
