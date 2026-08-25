package signaling

import (
	"fmt"
	"galleryline/domain"
	"strings"
)

type Diagnostic struct {
	Messages int
	Offers   int
	Answers  int
	Hangups  int
	Invalid  int
}

func Inspect(messages []domain.SignalMessage) Diagnostic {
	d := Diagnostic{}
	for _, m := range messages {
		d.Messages++
		switch m.Kind {
		case "offer":
			d.Offers++
		case "answer":
			d.Answers++
		case "hangup":
			d.Hangups++
		default:
			d.Invalid++
		}
	}
	return d
}
func Healthy(d Diagnostic) bool {
	return d.Messages > 0 && d.Offers <= d.Messages && d.Answers <= d.Messages && d.Invalid == 0
}
func Render(d Diagnostic) string {
	return fmt.Sprintf("messages=%d offers=%d answers=%d hangups=%d invalid=%d", d.Messages, d.Offers, d.Answers, d.Hangups, d.Invalid)
}
func Explain(d Diagnostic) []string {
	out := []string{}
	if d.Messages == 0 {
		out = append(out, "no signaling traffic")
	}
	if d.Offers == 0 {
		out = append(out, "caller offer missing")
	}
	if d.Answers == 0 {
		out = append(out, "callee answer missing")
	}
	if d.Hangups == 0 {
		out = append(out, "session has no hangup")
	}
	if d.Invalid > 0 {
		out = append(out, fmt.Sprintf("%d invalid messages", d.Invalid))
	}
	return out
}
func ParseKind(raw string) string {
	v := strings.ToLower(strings.TrimSpace(raw))
	switch v {
	case "offer", "answer", "hangup", "candidate":
		return v
	default:
		return "unknown"
	}
}
func ValidateMessage(m domain.SignalMessage) error {
	if m.From == "" || m.To == "" {
		return fmt.Errorf("missing participant")
	}
	if ParseKind(m.Kind) == "unknown" {
		return fmt.Errorf("unknown kind")
	}
	if m.Sequence < 0 {
		return fmt.Errorf("negative sequence")
	}
	return nil
}
func Relay(r *Router, m domain.SignalMessage) error {
	if e := ValidateMessage(m); e != nil {
		return e
	}
	return r.Send(m)
}
func Drain(r *Router, id string) []domain.SignalMessage { return r.Receive(id) }
func Replay(r *Router, id string, messages []domain.SignalMessage) int {
	n := 0
	for _, m := range messages {
		m.To = id
		if Relay(r, m) == nil {
			n++
		}
	}
	return n
}
func Conversation(c domain.CallSession, r *Router) []domain.SignalMessage {
	return append(r.Receive(c.CalleeID), r.Receive(c.CallerID)...)
}
