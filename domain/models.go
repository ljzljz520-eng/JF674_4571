package domain

import "time"

type Extension struct {
	ID, Nickname, Role string
	Online             bool
	LastSeen           time.Time
}
type CallSession struct {
	ID, CallerID, CalleeID, Status  string
	StartedAt, ConnectedAt, EndedAt time.Time
}
type PermissionNotice struct {
	ID, ExtensionID, Action, Message string
	Granted                          bool
	CreatedAt                        time.Time
}
type Device struct {
	ID, Name, Location string
	ExtensionID        string
	Active             bool
}
type CallRecord struct {
	ID, Caller, Callee, Outcome string
	Duration                    int64
	StartedAt                   time.Time
}
type SignalMessage struct {
	CallID, From, To, Kind, Payload string
	Sequence                        int
}

func (e Extension) Available() bool    { return e.Online && e.ID != "" }
func (c CallSession) IsFinished() bool { return c.Status == "ended" || c.Status == "rejected" }
func (c CallSession) Duration() time.Duration {
	if c.ConnectedAt.IsZero() {
		return 0
	}
	end := c.EndedAt
	if end.IsZero() {
		end = time.Now()
	}
	return end.Sub(c.ConnectedAt)
}
