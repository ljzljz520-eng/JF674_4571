package service

import (
	"fmt"
	"galleryline/auth"
	"galleryline/domain"
	"strings"
	"time"
)

type Operation struct {
	Code, Label, Description string
	RequiredRole             string
}
type OperationCatalog struct{ items map[string]Operation }

func NewOperationCatalog() *OperationCatalog {
	return &OperationCatalog{items: map[string]Operation{
		"dial":     {Code: "dial", Label: "Dial extension", Description: "Start an audio call", RequiredRole: "guide"},
		"accept":   {Code: "accept", Label: "Accept call", Description: "Connect incoming audio", RequiredRole: "desk"},
		"reject":   {Code: "reject", Label: "Reject call", Description: "Decline incoming call", RequiredRole: "desk"},
		"end":      {Code: "end", Label: "End call", Description: "Close active audio", RequiredRole: "guide"},
		"presence": {Code: "presence", Label: "Set presence", Description: "Publish online status", RequiredRole: "admin"},
		"records":  {Code: "records", Label: "View records", Description: "Review call history", RequiredRole: "admin"},
	}}
}
func (c *OperationCatalog) Get(code string) (Operation, bool) { v, ok := c.items[code]; return v, ok }
func (c *OperationCatalog) All() []Operation {
	out := []Operation{}
	for _, v := range c.items {
		out = append(out, v)
	}
	return out
}
func (c *OperationCatalog) Allowed(p auth.Principal, code string) bool {
	o, ok := c.Get(code)
	if !ok {
		return false
	}
	if p.Role == "admin" {
		return true
	}
	return p.Role == o.RequiredRole
}
func (c *OperationCatalog) Labels() map[string]string {
	m := map[string]string{}
	for k, v := range c.items {
		m[k] = v.Label
	}
	return m
}
func (c *OperationCatalog) Validate(code string) error {
	if strings.TrimSpace(code) == "" {
		return fmt.Errorf("empty operation")
	}
	if _, ok := c.Get(code); !ok {
		return fmt.Errorf("unknown operation %s", code)
	}
	return nil
}

type Schedule struct {
	Opening, Closing time.Time
	AllowedRoles     []string
}

func NewSchedule() Schedule {
	return Schedule{Opening: time.Unix(0, 0), Closing: time.Unix(86399, 0), AllowedRoles: []string{"guide", "desk", "device"}}
}
func (s Schedule) Open(at time.Time) bool { return !at.Before(s.Opening) && at.Before(s.Closing) }
func (s Schedule) RoleAllowed(role string) bool {
	for _, v := range s.AllowedRoles {
		if v == role {
			return true
		}
	}
	return false
}
func (s Schedule) Explain(p auth.Principal) string {
	if !s.RoleAllowed(p.Role) {
		return "role not allowed"
	}
	if !s.Open(time.Unix(3600, 0)) {
		return "closed"
	}
	return "ready"
}
func BuildNotice(p auth.Principal, action string) domain.PermissionNotice {
	return auth.NoticeFor(p, action)
}
