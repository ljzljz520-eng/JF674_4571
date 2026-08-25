package signaling

import (
	"fmt"
	"galleryline/domain"
	"sync"
)

type Router struct {
	mu     sync.Mutex
	queues map[string][]domain.SignalMessage
}

func NewRouter() *Router { return &Router{queues: map[string][]domain.SignalMessage{}} }
func (r *Router) Send(m domain.SignalMessage) error {
	if m.To == "" || m.From == "" {
		return fmt.Errorf("invalid route")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	m.Sequence = len(r.queues[m.To]) + 1
	r.queues[m.To] = append(r.queues[m.To], m)
	return nil
}
func (r *Router) Receive(id string) []domain.SignalMessage {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := append([]domain.SignalMessage(nil), r.queues[id]...)
	delete(r.queues, id)
	return out
}
func (r *Router) Pending(id string) int { r.mu.Lock(); defer r.mu.Unlock(); return len(r.queues[id]) }
func (r *Router) Broadcast(from string, targets []string, payload string) int {
	n := 0
	for _, t := range targets {
		if r.Send(domain.SignalMessage{From: from, To: t, Kind: "announce", Payload: payload}) == nil {
			n++
		}
	}
	return n
}
