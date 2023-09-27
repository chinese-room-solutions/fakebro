package agent

import (
	"time"
)

type AgentPool struct {
	Agents      chan *Agent
	DialTimeout time.Duration
	Size        int
	Roller      *Roller
}

// NewAgentPool creates a new active agent pool.
// dialTimeout is the timeout for dialing the server.
// size is the number of agents in the pool.
func NewAgentPool(
	dialTimeout time.Duration,
	size int,
	roller *Roller,
) *AgentPool {
	p := AgentPool{
		DialTimeout: dialTimeout,
		Size:        size,
		Roller:      roller,
	}
	p.Roll()

	return &p
}

// Roll rotates the agents for the pool.
func (p *AgentPool) Roll() {
	p.Close()

	p.Agents = make(chan *Agent, p.Size)
	for i := 0; i < p.Size; i++ {
		agent := p.Roller.Roll(time.Now().UnixNano(), p.DialTimeout)
		p.Agents <- agent
	}
}

// Get an agent from the pool and do some work using it. If the agent is bad (blocked, not responsive), it will be replaced.
// Throw ErrBadAgent in the `work` function if the agent is bad.
func (p *AgentPool) Do(work func(agent *Agent) (bool, error)) error {
	agent := <-p.Agents
	bad, err := work(agent)
	if bad {
		agent.Stop()
		agent = p.Roller.Roll(time.Now().UnixNano(), p.DialTimeout)
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
