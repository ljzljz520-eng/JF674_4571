package domain

import "errors"

var ErrNotFound = errors.New("not found")
var ErrOffline = errors.New("extension offline")
var ErrUnauthorized = errors.New("unauthorized")
var ErrInvalidTransition = errors.New("invalid call transition")
var ErrBusy = errors.New("extension busy")
