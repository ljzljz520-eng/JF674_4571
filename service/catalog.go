package service

import("sort";"strings";"galleryline/domain")
type DeviceCatalog struct{items map[string]domain.Device}
func NewDeviceCatalog()*DeviceCatalog{return &DeviceCatalog{items:map[string]domain.Device{}}}
func(c *DeviceCatalog) Add(d domain.Device)error{if d.Validate()!=nil{return domain.ErrNotFound};c.items[d.ID]=d;return nil}
func(c *DeviceCatalog) Get(id string)(domain.Device,bool){v,ok:=c.items[id];return v,ok}
func(c *DeviceCatalog) Remove(id string)bool{if _,ok:=c.items[id];!ok{return false};delete(c.items,id);return true}
func(c *DeviceCatalog) All()[]domain.Device{out:=make([]domain.Device,0,len(c.items));for _,v:=range c.items{out=append(out,v)};sort.Slice(out,func(i,j int)bool{return out[i].ID<out[j].ID});return out}
func(c *DeviceCatalog) Active()[]domain.Device{out:=[]domain.Device{};for _,v:=range c.items{if v.Active{out=append(out,v)}};sort.Slice(out,func(i,j int)bool{return out[i].ID<out[j].ID});return out}
func(c *DeviceCatalog) At(location string)[]domain.Device{location=strings.ToLower(strings.TrimSpace(location));out:=[]domain.Device{};for _,v:=range c.items{if strings.ToLower(v.Location)==location{out=append(out,v)}};return out}
func(c *DeviceCatalog) Activate(id string)error{v,ok:=c.Get(id);if !ok{return domain.ErrNotFound};v.Active=true;c.items[id]=v;return nil}
func(c *DeviceCatalog) Deactivate(id string)error{v,ok:=c.Get(id);if !ok{return domain.ErrNotFound};v.Active=false;c.items[id]=v;return nil}
func(c *DeviceCatalog) Count()int{return len(c.items)}
func(c *DeviceCatalog) CountActive()int{return len(c.Active())}
func(c *DeviceCatalog) IDs()[]string{out:=[]string{};for id:=range c.items{out=append(out,id)};sort.Strings(out);return out}
