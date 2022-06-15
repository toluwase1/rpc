package models

type AuthorizeRequest struct {
	Username string
	Password string
}

type NotifyResponse struct {
	JobId        string
	PreviousHash string
	CoinB1       string
	CoinB2       string
	MerkleBranch string
	Version      string
	NBit         string
	NTime        string
	CleanJobs    bool
}

type SubscribeResponse struct {
	Extranonce1      string
	Extranonce2_size int
	Subscriptions    []SubscriptionRequest
}

type SubscriptionRequest struct {
	SubscriptionType, Id string
}
