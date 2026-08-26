package signaling

import (
	"encoding/json"
	"fmt"
	"galleryline/domain"
	"sort"
)

func Encode(m domain.SignalMessage) ([]byte, error) {
	if e := ValidateMessage(m); e != nil {
		return nil, e
	}
	return json.Marshal(m)
}
func Decode(b []byte) (domain.SignalMessage, error) {
	var m domain.SignalMessage
	if e := json.Unmarshal(b, &m); e != nil {
		return m, e
	}
	return m, ValidateMessage(m)
}
func EncodeBatch(ms []domain.SignalMessage) ([]byte, error) {
	for _, m := range ms {
		if e := ValidateMessage(m); e != nil {
			return nil, e
		}
	}
	return json.Marshal(ms)
}
func DecodeBatch(b []byte) ([]domain.SignalMessage, error) {
	var ms []domain.SignalMessage
	if e := json.Unmarshal(b, &ms); e != nil {
		return nil, e
	}
	for _, m := range ms {
		if e := ValidateMessage(m); e != nil {
			return nil, e
		}
	}
	return ms, nil
}
func MustEncode(m domain.SignalMessage) []byte {
	b, e := Encode(m)
	if e != nil {
		panic(e)
	}
	return b
}
func MessageSummary(m domain.SignalMessage) string {
	return fmt.Sprintf("%d %s %s->%s", m.Sequence, m.Kind, m.From, m.To)
}
func IsAudio(m domain.SignalMessage) bool                             { return m.Payload == "audio" || m.Kind == "candidate" }
func WithSequence(m domain.SignalMessage, n int) domain.SignalMessage { m.Sequence = n; return m }
func Order(ms []domain.SignalMessage) []domain.SignalMessage {
	out := append([]domain.SignalMessage(nil), ms...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Sequence < out[j].Sequence })
	return out
}
