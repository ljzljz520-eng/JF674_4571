package domain

import (
	"fmt"
	"strings"
	"time"
)

func NewExtension(id, nickname, role string) Extension {
	return Extension{ID: id, Nickname: nickname, Role: role, LastSeen: time.Unix(0, 0)}
}
func NewDevice(id, name, location, extension string) Device {
	return Device{ID: id, Name: name, Location: location, ExtensionID: extension}
}
func NewCall(id, caller, callee string) CallSession {
	return CallSession{ID: id, CallerID: caller, CalleeID: callee, Status: "ringing", StartedAt: time.Unix(0, 0)}
}
func NewPermission(id, ext, action string) PermissionNotice {
	return PermissionNotice{ID: id, ExtensionID: ext, Action: action, CreatedAt: time.Unix(0, 0)}
}
func (e Extension) Label() string { return fmt.Sprintf("%s (%s)", e.Nickname, e.ID) }
func (e Extension) RoleLabel() string {
	if e.Role == "" {
		return "unknown"
	}
	return strings.Title(e.Role)
}
func (d Device) Label() string        { return fmt.Sprintf("%s @ %s", d.Name, d.Location) }
func (c CallSession) Parties() string { return c.CallerID + " -> " + c.CalleeID }
func (c CallSession) StatusLabel() string {
	if c.Status == "" {
		return "unknown"
	}
	return strings.ToUpper(c.Status)
}
func (c CallSession) HasAudio() bool       { return c.Status == "connected" }
func (c CallSession) CanAccept() bool      { return c.Status == "ringing" }
func (c CallSession) CanEnd() bool         { return c.Status == "connected" || c.Status == "ringing" }
func (n PermissionNotice) Summary() string { return n.Action + " for " + n.ExtensionID }
func (r CallRecord) Summary() string       { return fmt.Sprintf("%s/%s %s", r.Caller, r.Callee, r.Outcome) }
func (m SignalMessage) IsControl() bool {
	return m.Kind == "offer" || m.Kind == "answer" || m.Kind == "hangup"
}
