package agent

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

var NilPool = errors.New("pool is closed or not initialized")

type ActiveAgentPool struct {
	Agents      chan *Agent
	Roller      *Roller
	Network     string
	Address     string
	DialTimeout time.Duration
	Condition   func(*Agent) bool
	Limit       int
	Size        int
	Index       int
	lock        sync.Mutex
}

// NewActiveAgentPool creates a new active agent pool.
// address is the address of the server.
// dialTimeout is the timeout for dialing the server.
// limit is the maximum number of active agents.
// condition is the condition for selecting agents.
func NewActiveAgentPool(
	network string,
	address string,
	dialTimeout time.Duration,
	limit int,
	condition func(*Agent) bool,
) *ActiveAgentPool {
	p := ActiveAgentPool{
		Network:     network,
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
	var aa *ActiveAgent
	if p.ActiveAgents == nil {
		return nil, NilPool
	}
	if len(p.ActiveAgents) == 0 && p.Size < p.Limit {
		p.lock.Lock()
		defer p.lock.Unlock()
		aa = NewActiveAgent(p.Agents[p.Index])
		p.Index = (p.Index + 1) % len(p.Agents)
		p.Size++
		return aa, nil
	} else {
		aa := <-p.ActiveAgents
		return aa, nil
	}
}

// Put an active agent back to the pool.
// To delete an active agent, just don't put it back and decrement the pool size.
// Don't put an active agent back to the pool if it is not created by the pool as it may lead to deadlock.
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
		close(p.ActiveAgents)
		for aa := range p.ActiveAgents {
			aa.Stop()
		}
		p.ActiveAgents = nil
	}
}
