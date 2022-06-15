package models

type AuthorizeRequest struct {
	Username string
	Password string
}

type SubscribeResponse struct {
	Extra1        string
	ExtraSize2    int
	Subscriptions []SubscriptionRequest
}

type SubscriptionRequest struct {
	SubscriptionType string
	Id               string
}
