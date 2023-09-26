package agent

import (
	"errors"
	"time"
)

type AgentPool struct {
	Agents      chan *Agent
	DialTimeout time.Duration
	Size        int
	Roller      *Roller
}

var ErrBadAgent = errors.New("bad agent")

// NewAgentPool creates a new active agent pool.
// dialTimeout is the timeout for dialing the server.
// size is the number of agents in the pool.
func NewAgentPool(
	dialTimeout time.Duration,
	size int,
	roller *Roller,
) (*AgentPool, error) {
	p := AgentPool{
		DialTimeout: dialTimeout,
		Size:        size,
		Roller:      roller,
	}
	if err := p.Roll(); err != nil {
		return nil, err
	}

	return &p, nil
}

// Roll rotates the agents for the pool.
func (p *AgentPool) Roll() error {
	p.Close()

	p.Agents = make(chan *Agent, p.Size)
	for i := 0; i < p.Size; i++ {
		agent, err := p.Roller.Roll(time.Now().UnixNano(), p.DialTimeout)
		if err != nil {
			return err
		}
		p.Agents <- agent
	}

	return nil
}

// Get an agent from the pool and do some work using it. If the agent is bad (blocked, not responsive), it will be replaced.
// Throw ErrBadAgent in the `work` function if the agent is bad.
func (p *AgentPool) Do(work func(agent *Agent) error) error {
	agent := <-p.Agents
	err := work(agent)
	if err == ErrBadAgent {
		agent.Stop()
		agent, err = p.Roller.Roll(time.Now().UnixNano(), p.DialTimeout)
		if err != nil {
			return err
		}
	}
	p.Agents <- agent

	return err
}

// Close the pool.
func (p *AgentPool) Close() {
	if p.Agents != nil {
		close(p.Agents)
		for aa := range p.Agents {
			aa.Stop()
		}
		p.Agents = nil
	}
}
