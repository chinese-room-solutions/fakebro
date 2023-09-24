package agent

// import (
// 	"errors"
// 	"math/rand"
// 	"sync/atomic"
// 	"time"
// )

// type AgentPool struct {
// 	Agents      chan *Agent
// 	Roller      *Roller
// 	Network     string
// 	Address     string
// 	DialTimeout time.Duration
// 	Condition   func(*Agent) bool
// 	Limit       int
// 	Size        atomic.Int32
// 	Index       int
// }

// var ErrClosedPool = errors.New("pool is closed")

// // NewAgentPool creates a new active agent pool.
// // address is the address of the server.
// // dialTimeout is the timeout for dialing the server.
// // limit is the maximum number of active agents.
// // condition is the condition for selecting agents.
// func NewAgentPool(
// 	network string,
// 	address string,
// 	dialTimeout time.Duration,
// 	limit int,
// 	condition func(*Agent) bool,
// ) *AgentPool {
// 	p := AgentPool{
// 		Network:     network,
// 		Address:     address,
// 		DialTimeout: dialTimeout,
// 		Condition:   condition,
// 		Limit:       limit,
// 	}
// 	p.Roll()

// 	return &p
// }

// // Roll rotates the agents for the pool.
// func (p *AgentPool) Roll() {
// 	p.Close()
// 	p.Agents = SelectMany(p.Condition)
// 	rand.Shuffle(len(p.Agents), func(i, j int) {
// 		p.Agents[i], p.Agents[j] = p.Agents[j], p.Agents[i]
// 	})
// 	p.Agents = make(chan *Agent, p.Limit)
// }

// // Get an active agent from the pool.
// func (p *AgentPool) Get() *Agent {
// 	var aa *Agent
// 	if len(p.Agents) == 0 && int(p.Size.Load()) < p.Limit {
// 		aa = NewAgent(p.Agents[p.Index])
// 		p.Index = (p.Index + 1) % len(p.Agents)
// 		p.Size.Add(1)
// 		return aa
// 	}
// 	return <-p.Agents
// }

// // Put an active agent back to the pool.
// // To delete an active agent, just don't put it back and decrement the pool size.
// // Don't put an active agent back to the pool if it is not created by the pool as it may lead to deadlock.
// func (p *AgentPool) Put(aa *Agent) error {
// 	select {
// 	case p.Agents <- aa:
// 		return nil
// 	default:
// 		return ErrClosedPool
// 	}
// }

// // Close the pool.
// func (p *AgentPool) Close() {
// 	if p.Agents != nil {
// 		p.Index = 0
// 		close(p.Agents)
// 		for aa := range p.Agents {
// 			aa.Stop()
// 		}
// 		p.Agents = nil
// 	}
// }
