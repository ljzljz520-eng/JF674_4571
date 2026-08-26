package auth

import "galleryline/domain"

func NoticeFor(p Principal, action string) domain.PermissionNotice {
	return domain.PermissionNotice{ExtensionID: p.ID, Action: action, Message: "permission required: " + action, Granted: p.Role == "admin"}
}
func Approve(n domain.PermissionNotice, p Principal) domain.PermissionNotice {
	if p.Role == "admin" || p.ID == n.ExtensionID {
		n.Granted = true
	}
	return n
}
