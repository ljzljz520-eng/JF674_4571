package signaling

import "galleryline/domain"

func Offer(c domain.CallSession) domain.SignalMessage {
	return domain.SignalMessage{CallID: c.ID, From: c.CallerID, To: c.CalleeID, Kind: "offer", Payload: "audio"}
}
func Answer(c domain.CallSession) domain.SignalMessage {
	return domain.SignalMessage{CallID: c.ID, From: c.CalleeID, To: c.CallerID, Kind: "answer", Payload: "audio"}
}
func Hangup(c domain.CallSession, from string) domain.SignalMessage {
	return domain.SignalMessage{CallID: c.ID, From: from, To: other(c, from), Kind: "hangup"}
}
func other(c domain.CallSession, id string) string {
	if id == c.CallerID {
		return c.CalleeID
	}
	return c.CallerID
}
