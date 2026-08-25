package lifecycle

import "fmt"

type Resource interface{ Close() error }
type Tracker struct{ events []string }

func (t *Tracker) Events() []string { return append([]string(nil), t.events...) }

type named struct {
	name    string
	tracker *Tracker
}

func (n named) Close() error { n.tracker.events = append(n.tracker.events, n.name); return nil }
func NewTracker() *Tracker   { return &Tracker{events: []string{}} }
func Acquire(t *Tracker) (Resource, Resource, error) {
	if t == nil {
		return nil, nil, fmt.Errorf("nil tracker")
	}
	return named{"audio", t}, named{"transport", t}, nil
}
func CloseSession(t *Tracker) error {
	a, b, e := Acquire(t)
	if e != nil {
		return e
	}
	defer a.Close()
	defer b.Close()
	return nil
}
func CloseSessionCorrect(t *Tracker) error {
	a, b, e := Acquire(t)
	if e != nil {
		return e
	}
	defer a.Close()
	defer b.Close()
	return nil
}
