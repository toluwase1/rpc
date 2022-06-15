package handlers

import (
	"rpc-client-server/database"
	"rpc-client-server/models"
)

var miningService database.MiningService = database.MiningServer{}

type Mining int

func (mining *Mining) Authorize(args *models.AuthorizeRequest, result *bool) (e error) {
	*result, e = miningService.Authorize(*args)
	return
}

func (mining *Mining) Subscribe(result *models.SubscribeResponse) (e error) {
	*result, e = miningService.Subscribe()
	return
}

func (mining *Mining) Notify(result *models.NotifyResponse) (e error) {
	*result, e = miningService.Notify()
	return
}
