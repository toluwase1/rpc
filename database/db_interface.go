package database

import "rpc-client-server/models"

type AuthorizationRequestRepository interface {
	SaveAuthorizationRequest(request models.Authorization) (models.Authorization, error)
	FindAuthorizationRequestById(id int) (models.Authorization, error)
	FindAuthorizationRequestByUsername(username string) (models.Authorization, error)
	SaveSubscription(subscription models.Subscription) (models.Subscription, error)
}
