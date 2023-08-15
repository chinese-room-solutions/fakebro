package roundtripper

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/chinese-room-solutions/fakebro/agent"
	"golang.org/x/exp/maps"
)

type RoundTripper struct {
	Pool    *agent.ActiveAgentPool
	Headers map[string]string
}

func New(
	address string,
	dialTimeout time.Duration,
	condition func(*agent.Agent) bool,
	limit int,
	headers map[string]string,
) *RoundTripper {
	return &RoundTripper{
		Pool:    agent.NewActiveAgentPool(address, dialTimeout, condition, limit),
		Headers: headers,
	}
}

func (rt *RoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	// Get an agent from the pool
	a, err := rt.Pool.Get()
	if err != nil {
		return nil, err
	}

	// Merge the custom headers with the agent headers with the precedence of the custom headers
	headers := map[string]string{}
	maps.Copy(&headers, a.Headers)
	maps.Copy(&headers, rt.Headers)

	for header, value := range headers {
		r.Header.Set(header, value)
	}

	transport := &http.Transport{
		DialTLS: a.GetConnection,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // this is insecure; only use if you know what you're doing
		},
	}

	return rt.inner.RoundTrip(r)
}
