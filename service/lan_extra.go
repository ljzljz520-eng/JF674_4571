package service

import (
	"fmt"
	"galleryline/domain"
	"strings"
)

func (g *LanGuide) PortNumber() string   { return strings.TrimSpace(g.Port) }
func (g *LanGuide) AddressLabel() string { return fmt.Sprintf("LAN endpoint %s", g.Endpoint()) }
func (g *LanGuide) ReadyMessage() string {
	if !g.Valid() {
		return "LAN setup incomplete"
	}
	return "LAN ready at " + g.URL()
}
func (g *LanGuide) StepCount() int { return len(g.Steps) }
func (g *LanGuide) AddStep(step string) bool {
	step = strings.TrimSpace(step)
	if step == "" {
		return false
	}
	g.Steps = append(g.Steps, step)
	return true
}
func (g *LanGuide) ReplaceStep(index int, step string) bool {
	if index < 0 || index >= len(g.Steps) || strings.TrimSpace(step) == "" {
		return false
	}
	g.Steps[index] = step
	return true
}
func (g *LanGuide) Complete(done map[int]bool) bool {
	for i := range g.Steps {
		if !done[i] {
			return false
		}
	}
	return true
}
func (g *LanGuide) DeviceHint(d domain.Device) string {
	return d.Name + " at " + d.Location + " uses " + g.URL()
}
func (g *LanGuide) ExtensionHint(e domain.Extension) string {
	return e.Label() + " can connect via " + g.Endpoint()
}
func (g *LanGuide) Clone() *LanGuide {
	return &LanGuide{SSID: g.SSID, Address: g.Address, Port: g.Port, Steps: append([]string(nil), g.Steps...)}
}
