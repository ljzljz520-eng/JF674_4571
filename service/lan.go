package service

import (
	"fmt"
	"galleryline/domain"
	"sort"
	"strings"
)

type LanGuide struct {
	SSID, Address, Port string
	Steps               []string
}

func NewLanGuide(address, port string) *LanGuide {
	return &LanGuide{SSID: "GalleryLine", Address: address, Port: port, Steps: []string{"connect to the exhibit LAN", "open the operator console", "register an extension", "place a test call", "confirm audio and hang up"}}
}
func (g *LanGuide) Endpoint() string { return g.Address + ":" + g.Port }
func (g *LanGuide) URL() string      { return "http://" + g.Endpoint() }
func (g *LanGuide) Step(n int) string {
	if n < 1 || n > len(g.Steps) {
		return ""
	}
	return g.Steps[n-1]
}
func (g *LanGuide) Checklist() []string { return append([]string(nil), g.Steps...) }
func (g *LanGuide) Valid() bool         { return g.Address != "" && g.Port != "" && len(g.Steps) >= 4 }
func (g *LanGuide) Render() string      { return strings.Join(g.Steps, "\n") }

type Roster struct{ items map[string]domain.Extension }

func NewRoster() *Roster { return &Roster{items: map[string]domain.Extension{}} }
func (r *Roster) Add(e domain.Extension) error {
	if e.Validate() != nil {
		return domain.ErrNotFound
	}
	r.items[e.ID] = e
	return nil
}
func (r *Roster) Remove(id string) bool {
	if _, ok := r.items[id]; !ok {
		return false
	}
	delete(r.items, id)
	return true
}
func (r *Roster) Get(id string) (domain.Extension, bool) { v, ok := r.items[id]; return v, ok }
func (r *Roster) All() []domain.Extension {
	out := make([]domain.Extension, 0, len(r.items))
	for _, v := range r.items {
		out = append(out, v)
	}
	return domain.ExtensionSort(out)
}
func (r *Roster) Online() []domain.Extension { return domain.FindOnline(r.All()) }
func (r *Roster) SetOnline(id string, value bool) error {
	v, ok := r.Get(id)
	if !ok {
		return domain.ErrNotFound
	}
	v.Online = value
	r.items[id] = v
	return nil
}
func (r *Roster) Roles() []string {
	m := map[string]bool{}
	for _, v := range r.items {
		m[v.Role] = true
	}
	out := []string{}
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
func (r *Roster) Search(term string) []domain.Extension {
	term = strings.ToLower(term)
	out := []domain.Extension{}
	for _, v := range r.items {
		if strings.Contains(strings.ToLower(v.Nickname), term) || strings.Contains(strings.ToLower(v.ID), term) {
			out = append(out, v)
		}
	}
	return domain.ExtensionSort(out)
}
func (r *Roster) Summary() string {
	return fmt.Sprintf("%d extensions, %d online", len(r.items), len(r.Online()))
}
