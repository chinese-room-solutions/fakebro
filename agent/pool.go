package agent

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

var NilPool = errors.New("pool is closed or not initialized")

type ActiveAgentPool struct {
	ActiveAgents chan *ActiveAgent
	Agents       []*Agent
	Address      string
	DialTimeout  time.Duration
	Condition    func(*Agent) bool
	Limit        int
	Size         int
	Index        int
	lock         sync.Mutex
}

// NewActiveAgentPool creates a new active agent pool.
// address is the address of the server.
// dialTimeout is the timeout for dialing the server.
// condition is the condition for selecting agents.
// limit is the maximum number of active agents.
func NewActiveAgentPool(
	address string,
	dialTimeout time.Duration,
	condition func(*Agent) bool,
	limit int,
) *ActiveAgentPool {
	p := ActiveAgentPool{
		Address:     address,
		DialTimeout: dialTimeout,
		Condition:   condition,
		Limit:       limit,
	}
	p.Roll()

	return &p
}

// Roll rotates the agents for the pool.
func (p *ActiveAgentPool) Roll() {
	p.Close()
	p.Agents = SelectMany(p.Condition)
	rand.Shuffle(len(p.Agents), func(i, j int) {
		p.Agents[i], p.Agents[j] = p.Agents[j], p.Agents[i]
	})
	p.ActiveAgents = make(chan *ActiveAgent, p.Limit)
}

// Get an active agent from the pool.
func (p *ActiveAgentPool) Get() (*ActiveAgent, error) {
	var err error
	var aa *ActiveAgent
	if p.ActiveAgents == nil {
		return nil, NilPool
	}
	if len(p.ActiveAgents) == 0 && p.Size < p.Limit {
		p.lock.Lock()
		defer p.lock.Unlock()
		aa, err = NewActiveAgent(p.Agents[p.Index], p.Address, p.DialTimeout)
		p.Index = (p.Index + 1) % len(p.Agents)
		if err != nil {
			return nil, err
		}
		p.Size++
		return aa, nil
	} else {
		aa := <-p.ActiveAgents
		return aa, nil
	}
}

// Put an active agent back to the pool.
// To delete an active agent, just don't put it back and decrement the pool size.
func (p *ActiveAgentPool) Put(aa *ActiveAgent) error {
	select {
	case p.ActiveAgents <- aa:
		return nil
	default:
		return NilPool
	}
}

// Close the pool.
func (p *ActiveAgentPool) Close() {
	if p.ActiveAgents != nil {
		p.Index = 0
		for aa := range p.ActiveAgents {
			aa.Stop()
		}
		close(p.ActiveAgents)
		p.ActiveAgents = nil
	}
}
