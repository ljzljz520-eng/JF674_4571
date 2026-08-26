package signaling

import (
	"galleryline/domain"
	"sort"
)

type Session struct {
	Call     domain.CallSession
	messages []domain.SignalMessage
}

func NewSession(c domain.CallSession) *Session {
	return &Session{Call: c, messages: []domain.SignalMessage{}}
}
func (s *Session) Add(m domain.SignalMessage) { s.messages = append(s.messages, m) }
func (s *Session) Messages() []domain.SignalMessage {
	return append([]domain.SignalMessage(nil), s.messages...)
}
func (s *Session) Controls() []domain.SignalMessage {
	out := []domain.SignalMessage{}
	for _, m := range s.messages {
		if m.IsControl() {
			out = append(out, m)
		}
	}
	return out
}
func (s *Session) Last() domain.SignalMessage {
	if len(s.messages) == 0 {
		return domain.SignalMessage{}
	}
	return s.messages[len(s.messages)-1]
}
func (s *Session) Ordered() []domain.SignalMessage {
	out := s.Messages()
	sort.SliceStable(out, func(i, j int) bool { return out[i].Sequence < out[j].Sequence })
	return out
}
func (s *Session) Complete() bool             { return s.Call.IsFinished() }
func (s *Session) Participant(id string) bool { return id == s.Call.CallerID || id == s.Call.CalleeID }
func (s *Session) Other(id string) string {
	if id == s.Call.CallerID {
		return s.Call.CalleeID
	}
	return s.Call.CallerID
}
func (s *Session) Count(kind string) int {
	n := 0
	for _, m := range s.messages {
		if m.Kind == kind {
			n++
		}
	}
	return n
}
