package service

import (
	"galleryline/domain"
	"sort"
)

type CallQueue struct{ items []domain.CallSession }

func NewCallQueue() *CallQueue { return &CallQueue{items: []domain.CallSession{}} }
func (q *CallQueue) Push(c domain.CallSession) {
	if c.Status == "ringing" {
		q.items = append(q.items, c)
		sort.SliceStable(q.items, func(i, j int) bool { return q.items[i].StartedAt.Before(q.items[j].StartedAt) })
	}
}
func (q *CallQueue) Pop() (domain.CallSession, bool) {
	if len(q.items) == 0 {
		return domain.CallSession{}, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}
func (q *CallQueue) Peek() (domain.CallSession, bool) {
	if len(q.items) == 0 {
		return domain.CallSession{}, false
	}
	return q.items[0], true
}
func (q *CallQueue) Remove(id string) bool {
	for i, v := range q.items {
		if v.ID == id {
			q.items = append(q.items[:i], q.items[i+1:]...)
			return true
		}
	}
	return false
}
func (q *CallQueue) Contains(id string) bool {
	for _, v := range q.items {
		if v.ID == id {
			return true
		}
	}
	return false
}
func (q *CallQueue) Len() int { return len(q.items) }
func (q *CallQueue) ForExtension(id string) []domain.CallSession {
	out := []domain.CallSession{}
	for _, v := range q.items {
		if v.CalleeID == id {
			out = append(out, v)
		}
	}
	return out
}
func (q *CallQueue) IDs() []string {
	out := make([]string, 0, len(q.items))
	for _, v := range q.items {
		out = append(out, v.ID)
	}
	return out
}
func (q *CallQueue) Clear() { q.items = nil }
func (q *CallQueue) Snapshot() []domain.CallSession {
	return append([]domain.CallSession(nil), q.items...)
}
