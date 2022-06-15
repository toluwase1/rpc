package models

import "time"

type Subscription struct {
	Id              int
	SubscriptionId1 string
	SubscriptionId2 string
	Extranonce1     string
	DateCreated     time.Time
}
