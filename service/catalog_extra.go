package service

import("strings";"galleryline/domain")
func(c *DeviceCatalog) Search(term string)[]domain.Device{term=strings.ToLower(strings.TrimSpace(term));out:=[]domain.Device{};for _,v:=range c.items{if strings.Contains(strings.ToLower(v.Name),term)||strings.Contains(strings.ToLower(v.Location),term){out=append(out,v)}};return out}
func(c *DeviceCatalog) ForExtension(id string)[]domain.Device{out:=[]domain.Device{};for _,v:=range c.items{if v.ExtensionID==id{out=append(out,v)}};return out}
func(c *DeviceCatalog) Has(id string)bool{_,ok:=c.items[id];return ok}
func(c *DeviceCatalog) Toggle(id string)error{v,ok:=c.Get(id);if !ok{return domain.ErrNotFound};v.Active=!v.Active;c.items[id]=v;return nil}
func(c *DeviceCatalog) Replace(d domain.Device)error{if d.Validate()!=nil{return domain.ErrNotFound};if !c.Has(d.ID){return domain.ErrNotFound};c.items[d.ID]=d;return nil}
func(c *DeviceCatalog) Locations()[]string{m:=map[string]bool{};for _,v:=range c.items{m[v.Location]=true};out:=[]string{};for x:=range m{out=append(out,x)};return out}
func(c *DeviceCatalog) ActiveAt(location string)[]domain.Device{out:=[]domain.Device{};for _,v:=range c.At(location){if v.Active{out=append(out,v)}};return out}
