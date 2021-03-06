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
		return false, err
	}
	return true, nil
}

func (db MiningServer) Subscribe() (response models.SubscribeResponse, e error) {
	var subscription = models.Subscription{
		SubId1:    stringGenerator(20),
		SubId2:    stringGenerator(20),
		Extra:     stringGenerator(10),
		CreatedAt: time.Now(),
	}
	_, err := db.subscriptionRepository.SaveSubscription(subscription)
	if err != nil {
		e = err
		return
	}
	response = models.SubscribeResponse{
		Extra1:     subscription.Extra,
		ExtraSize2: 4,
		Subscriptions: []models.SubscriptionRequest{
			{
				"mining.set_difficulty",
				subscription.SubId1,
			},
			{
				"mining.notify",
				subscription.SubId2,
			},
		},
	}
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
