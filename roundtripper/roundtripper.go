package roundtripper

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/chinese-room-solutions/fakebro/agent"
	"github.com/chinese-room-solutions/fakebro/user_agent"
	"golang.org/x/exp/maps"
	"golang.org/x/net/http2"

	"github.com/quic-go/quic-go/http3"
)

type ErrUnsupportedALPN struct {
	ALPN string
}

func (e ErrUnsupportedALPN) Error() string {
	return fmt.Sprintf("unsupported ALPN: %v", e.ALPN)
}

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
		r.Header.Set("Host", strings.Split(r.Host, ":")[0])

		conn, err := a.DialTLS(context.Background(), "tcp", r.Host)
		if err != nil {
			return rt.ErrorHandler(err), err
		}

		switch a.ALPN {
		case "http/1.1":
			r.Proto = "HTTP/1.1"
			r.ProtoMajor = 1
			r.ProtoMinor = 1

			err := r.Write(conn)
			if err != nil {
				return rt.ErrorHandler(err), err
			}
			resp, err = http.ReadResponse(bufio.NewReader(conn), r)
			return rt.ErrorHandler(err), err
		case "h2", "":
			r.Proto = "HTTP/2.0"
			r.ProtoMajor = 2
			r.ProtoMinor = 0

			tr := http2.Transport{}
			cConn, err := tr.NewClientConn(conn)
			if err != nil {
				return rt.ErrorHandler(err), err
			}
			resp, err = cConn.RoundTrip(r)
			return rt.ErrorHandler(err), err
		case "h3":
			r.Proto = "HTTP/3.0"
			r.ProtoMajor = 3
			r.ProtoMinor = 0

			tr := http3.RoundTripper{
				// TLSClientConfig: a.TLSConfig,
			}
			_ = tr
			return false, ErrUnsupportedALPN{ALPN: a.ALPN}
		default:
			return false, ErrUnsupportedALPN{ALPN: a.ALPN}
		}

		// resp, err = a.T.RoundTrip(r)

		// return rt.ErrorHandler(err), err
	})

	return resp, err
}
