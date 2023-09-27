package agent

import (
	"errors"
	"slices"
	"time"

	"github.com/chinese-room-solutions/fakebro/user_agent"
	"golang.org/x/exp/maps"
)

type Roller struct {
	AllowedTokens []user_agent.TokenType
}

var (
	ErrUnfulfilledHeaderCondition = errors.New("no headers fulfill the condition")
	ErrUnfulfilledTLSCondition    = errors.New("no TLS configs fulfill the condition")
)

func NewRoller(condition func(user_agent.TokenType) bool) (*Roller, error) {
	allowedTokens := []user_agent.TokenType{}
	if condition != nil {
		for i := 0; i < int(user_agent.TotalTokens); i++ {
			if condition(user_agent.TokenType(i)) {
				allowedTokens = append(allowedTokens, user_agent.TokenType(i))
			}
		}
	}

	return &Roller{
		AllowedTokens: allowedTokens,
	}, nil
}

func (r *Roller) Roll(seed int64, dialTimeout time.Duration) (*Agent, error) {
	var ua = user_agent.NewUserAgent(20, seed, r.AllowedTokens...)
	var headers = map[string]string{}
	var tlsConfig *TLSConfig

	for _, c := range BaseHeaders {
		if slices.Contains[[]string, string](c.Clients, ua.Client) {
			maps.Copy(headers, c.Value)
			break
		}
	}
	if len(headers) == 0 {
		return nil, ErrUnfulfilledHeaderCondition
	}
	headers["User-Agent"] = ua.Header

	for _, c := range BaseTLSConfigs {
		if slices.Contains[[]string, string](c.Clients, ua.Client) &&
			slices.Contains[[]string, string](c.Versions, ua.Version) {
			tlsConfig = c
			break
		}
	}
	if tlsConfig == nil {
		return nil, ErrUnfulfilledTLSCondition
	}

	return NewAgent(
		ua.Client,
		ua.Version,
		tlsConfig.Value,
		headers,
		dialTimeout,
	), nil
}
