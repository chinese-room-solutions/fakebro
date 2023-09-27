package roundtripper

import (
	"net/http"
	"time"

	"github.com/chinese-room-solutions/fakebro/agent"
	"github.com/chinese-room-solutions/fakebro/user_agent"
	"golang.org/x/exp/maps"
)

type RoundTripper struct {
	Pool         *agent.AgentPool
	Headers      map[string]string
	ErrorHandler func(error) bool
}

// New creates a new round tripper.
// dialTimeout is the timeout for dialing the server.
// poolSize is the number of agents in the pool.
// condition is a function that returns true if the token should be used.
// headers is a map of headers to add to the request.
// errorHandler is a function that returns true if error considered due to bad agent and it must be rotated.
func New(
	dialTimeout time.Duration,
	poolSize int,
	condition func(user_agent.TokenType) bool,
	headers map[string]string,
	errorHandler func(error) (badAgent bool),
) *RoundTripper {
	roller := agent.NewRoller(condition)
	pool := agent.NewAgentPool(dialTimeout, poolSize, roller)
	return &RoundTripper{
		Pool:         pool,
		Headers:      headers,
		ErrorHandler: errorHandler,
	}
}

func (rt *RoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	var resp *http.Response
	err := rt.Pool.Do(func(a *agent.Agent) (bool, error) {
		headers := map[string]string{}
		var err error
		maps.Copy(headers, a.Headers)
		maps.Copy(headers, rt.Headers)

		for header, value := range headers {
			r.Header.Set(header, value)
		}
		r.Header.Set("Host", r.Host)
		resp, err = a.T.RoundTrip(r)

		return rt.ErrorHandler(err), err
	})

	return resp, err
}
