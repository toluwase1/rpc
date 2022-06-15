package database

import (
	"math/rand"
	"rpc-client-server/models"
	"strings"
	"time"
)

type MiningServer struct {
	authorizationRequestRepository PostgresDb
	subscriptionRepository         PostgresDb
}

func (db MiningServer) Authorize(request models.AuthorizeRequest) (response bool, e error) {
	authReq, _ := db.authorizationRequestRepository.FindAuthorizationRequestByUsername(request.Username)
	authReq.Username = request.Username
	authReq.LastLoggedInTime = time.Now()
	_, err := db.authorizationRequestRepository.SaveAuthorizationRequest(authReq)
	if err != nil {
		e = err
		return
	}
	response = true
	return
}

func (db MiningServer) Subscribe() (response models.SubscribeResponse, e error) {
	var subscription = models.Subscription{
		SubscriptionId1: stringGenerator(20),
		SubscriptionId2: stringGenerator(20),
		Extranonce1:     stringGenerator(10),
		DateCreated:     time.Now(),
	}
	_, err := db.subscriptionRepository.SaveSubscription(subscription)
	if err != nil {
		e = err
		return
	}
	response = models.SubscribeResponse{
		Extranonce1:      subscription.Extranonce1,
		Extranonce2_size: 4,
		Subscriptions: []models.SubscriptionRequest{
			{
				"mining.set_difficulty",
				subscription.SubscriptionId1,
			},
			{
				"mining.notify",
				subscription.SubscriptionId2,
			},
		},
	}
	return
}

func (db MiningServer) Notify() (response models.NotifyResponse, e error) {
	return
}


func stringGenerator(n int) string {
	const alphabet = "abcdefghijklmnopgrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
