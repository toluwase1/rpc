package models

import "time"

type Authorization struct {
	Id               int
	Username         string
	LastLoggedInTime time.Time
}
