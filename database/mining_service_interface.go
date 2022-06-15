package database

import "rpc-client-server/models"

type MiningService interface {
	Authorize(request models.AuthorizeRequest) (bool, error)
	Subscribe() (models.SubscribeResponse, error)
}
