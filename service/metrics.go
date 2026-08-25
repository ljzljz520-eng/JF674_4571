package service

import (
	"fmt"
	"galleryline/domain"
	"sort"
)

type Metrics struct {
	Calls, Connected, Ended, Rejected int
	Durations                         int64
}

func CollectMetrics(calls []domain.CallSession, records []domain.CallRecord) Metrics {
	m := Metrics{Calls: len(calls)}
	for _, c := range calls {
		switch c.Status {
		case "connected":
			m.Connected++
		case "ended":
			m.Ended++
		case "rejected":
			m.Rejected++
		}
	}
	for _, r := range records {
		m.Durations += r.Duration
	}
	return m
}
func (m Metrics) AverageDuration() int64 {
	if m.Ended == 0 {
		return 0
	}
	return m.Durations / int64(m.Ended)
}
func (m Metrics) SuccessRate() float64 {
	if m.Calls == 0 {
		return 0
	}
	return float64(m.Ended) / float64(m.Calls)
}
func (m Metrics) Labels() map[string]string {
	return map[string]string{"calls": itoa(m.Calls), "connected": itoa(m.Connected), "ended": itoa(m.Ended), "rejected": itoa(m.Rejected), "duration": itoa64(m.Durations)}
}
func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	return fmt.Sprintf("%d", v)
}
func itoa64(v int64) string { return fmt.Sprintf("%d", v) }
func SortRecords(in []domain.CallRecord) []domain.CallRecord {
	out := append([]domain.CallRecord(nil), in...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Duration > out[j].Duration })
	return out
}
