package agent

import (
	"errors"
	"math/rand"
	"time"

	"golang.org/x/exp/maps"
)

type Roller struct {
	FilteredTLSConfigs []*TLSConfig
}

var ErrUnfulfilledCondition = errors.New("no TLS configs fulfill the condition")

func NewRoller(condition func(client, version string) bool) (*Roller, error) {
	configs := []*TLSConfig{}
	for _, c := range BaseTLSConfigs {
		for _, v := range c.Versions {
			if condition(c.Client, v) {
				configs = append(configs, &TLSConfig{
					Client:  c.Client,
					Version: v,
					Value:   c.Value,
				})
			}
		}
	}
	if len(configs) == 0 {
		return nil, ErrUnfulfilledCondition
	}

	return &Roller{configs}, nil
}

func (r *Roller) Roll(network, address string, dialTimeout time.Duration) *Agent {
	var headers map[string]string
	i := rand.Intn(len(r.FilteredTLSConfigs))
	maps.Copy(headers, BaseHeaders[r.FilteredTLSConfigs[i].Client])
	j := rand.Intn(len(IdentityHeadersPool[r.FilteredTLSConfigs[i].Client]))
	maps.Copy(headers, IdentityHeadersPool[r.FilteredTLSConfigs[i].Client][j])

	return NewAgent(
		r.FilteredTLSConfigs[i].Client,
		r.FilteredTLSConfigs[i].Version,
		r.FilteredTLSConfigs[i].Value,
		headers,
		dialTimeout,
	)
}
