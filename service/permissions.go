package service

import (
	"galleryline/auth"
	"galleryline/domain"
	"galleryline/storage"
	"time"
)

type PermissionService struct{ store *storage.Store }

func NewPermissionService(s *storage.Store) *PermissionService { return &PermissionService{store: s} }
func (p *PermissionService) Request(pr auth.Principal, action, id string) (domain.PermissionNotice, error) {
	n := auth.NoticeFor(pr, action)
	n.ID = id
	n.CreatedAt = time.Unix(0, 0)
	return n, p.store.SavePermission(n)
}
func (p *PermissionService) Grant(n domain.PermissionNotice, pr auth.Principal) error {
	n = auth.Approve(n, pr)
	return p.store.SavePermission(n)
}
