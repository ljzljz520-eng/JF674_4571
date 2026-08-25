package service

import (
	"galleryline/domain"
	"galleryline/storage"
)

type DeviceService struct{ store *storage.Store }

func NewDeviceService(s *storage.Store) *DeviceService { return &DeviceService{store: s} }
func (d *DeviceService) Add(v domain.Device) error     { v.Active = true; return d.store.SaveDevice(v) }
func (d *DeviceService) Activate(id string) error {
	v, e := d.store.GetDevice(id)
	if e != nil {
		return e
	}
	v.Active = true
	return d.store.SaveDevice(v)
}
