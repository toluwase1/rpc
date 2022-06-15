package models

import "time"

type Subscription struct {
	Id        int
	SubId1    string
	SubId2    string
	Extra     string
	CreatedAt time.Time
}
