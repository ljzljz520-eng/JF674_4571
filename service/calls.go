package service

import (
	"fmt"
	"galleryline/auth"
	"galleryline/domain"
	"galleryline/signaling"
	"galleryline/storage"
	"time"
)

type CallManager struct {
	store  *storage.Store
	dir    *Directory
	router *signaling.Router
}

func NewCallManager(s *storage.Store, d *Directory, r *signaling.Router) *CallManager {
	return &CallManager{store: s, dir: d, router: r}
}
func (m *CallManager) Dial(p auth.Principal, caller, callee, id string) (domain.CallSession, error) {
	target, e := m.dir.Find(callee)
	if e != nil {
		return domain.CallSession{}, e
	}
	if !auth.CanDial(p, target) {
		return domain.CallSession{}, domain.ErrUnauthorized
	}
	if !target.Online {
		return domain.CallSession{}, domain.ErrOffline
	}
	c := domain.CallSession{ID: id, CallerID: caller, CalleeID: callee, Status: "ringing", StartedAt: time.Unix(0, 0)}
	if e = m.store.SaveCall(c); e != nil {
		return c, e
	}
	return c, m.router.Send(signaling.Offer(c))
}
func (m *CallManager) Accept(id string) error {
	c, e := m.store.GetCall(id)
	if e != nil {
		return e
	}
	if c.Status != "ringing" {
		return domain.ErrInvalidTransition
	}
	c.Status = "connected"
	c.ConnectedAt = time.Unix(1, 0)
	if e = m.store.SaveCall(c); e != nil {
		return e
	}
	return m.router.Send(signaling.Answer(c))
}
func (m *CallManager) End(id, from string) error {
	c, e := m.store.GetCall(id)
	if e != nil {
		return e
	}
	if c.IsFinished() {
		return domain.ErrInvalidTransition
	}
	c.Status = "ended"
	c.EndedAt = time.Unix(11, 0)
	if e = m.store.SaveCall(c); e != nil {
		return e
	}
	r := domain.CallRecord{ID: fmt.Sprintf("record-%s", id), Caller: c.CallerID, Callee: c.CalleeID, Outcome: "completed", Duration: int64(c.EndedAt.Sub(c.ConnectedAt).Seconds()), StartedAt: c.StartedAt}
	if e = m.store.SaveRecord(r); e != nil {
		return e
	}
	return m.router.Send(signaling.Hangup(c, from))
}
func (m *CallManager) Reject(id string) error {
	c, e := m.store.GetCall(id)
	if e != nil {
		return e
	}
	if c.Status != "ringing" {
		return domain.ErrInvalidTransition
	}
	c.Status = "rejected"
	return m.store.SaveCall(c)
}
func (m *CallManager) History() ([]domain.CallRecord, error) { return m.store.ListRecords() }
