package auth

import "galleryline/domain"

type Principal struct{ ID, Role string }

func CanDial(p Principal, target domain.Extension) bool {
	return p.ID != "" && target.ID != "" && (p.Role == "admin" || p.Role == "guide" || p.Role == "desk" || p.Role == "device")
}
func CanViewRecords(p Principal) bool { return p.Role == "admin" || p.Role == "desk" }
func CanManage(p Principal) bool      { return p.Role == "admin" }
func NormalizeRole(role string) string {
	switch role {
	case "guide", "desk", "device", "admin":
		return role
	default:
		return "guest"
	}
}
