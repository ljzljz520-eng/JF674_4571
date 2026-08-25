package domain

import "strings"

func (e Extension) Validate() error {
	if strings.TrimSpace(e.ID) == "" || strings.TrimSpace(e.Nickname) == "" {
		return ErrNotFound
	}
	return nil
}
func (c CallSession) Validate() error {
	if c.ID == "" || c.CallerID == "" || c.CalleeID == "" {
		return ErrNotFound
	}
	if c.CallerID == c.CalleeID {
		return ErrInvalidTransition
	}
	return nil
}
func (d Device) Validate() error {
	if d.ID == "" || d.ExtensionID == "" {
		return ErrNotFound
	}
	return nil
}
func (n PermissionNotice) Validate() error {
	if n.ID == "" || n.ExtensionID == "" || n.Action == "" {
		return ErrNotFound
	}
	return nil
}
func NormalizeStatus(status string) string {
	switch strings.ToLower(status) {
	case "ringing", "connected", "ended", "rejected":
		return strings.ToLower(status)
	default:
		return "unknown"
	}
}
