package service

import (
	"galleryline/domain"
	"galleryline/storage"
	"time"
)

type Directory struct{ store *storage.Store }

func NewDirectory(s *storage.Store) *Directory { return &Directory{store: s} }
func (d *Directory) Register(e domain.Extension) error {
	e.Online = true
	e.LastSeen = time.Unix(0, 0)
	return d.store.SaveExtension(e)
}
func (d *Directory) SetPresence(id string, online bool) error {
	e, err := d.store.GetExtension(id)
	if err != nil {
		return err
	}
	e.Online = online
	e.LastSeen = time.Unix(0, 0)
	return d.store.SaveExtension(e)
}
func (d *Directory) Find(id string) (domain.Extension, error) { return d.store.GetExtension(id) }
func (d *Directory) List() ([]domain.Extension, error)        { return d.store.ListExtensions() }
func (d *Directory) Online() ([]domain.Extension, error) {
	all, e := d.List()
	if e != nil {
		return nil, e
	}
	out := []domain.Extension{}
	for _, v := range all {
		if v.Online {
			out = append(out, v)
		}
	}
	return out, nil
}
