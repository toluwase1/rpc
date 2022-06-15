package database

import (
	"github.com/jinzhu/gorm"
	"rpc-client-server/models"
)

type PostgresDb struct {
	Db *gorm.DB
}

func (db *PostgresDb) SaveAuthorizationRequest(request models.Authorization) (res models.Authorization, e error) {
	if request.Id > 0 {
		if err := db.Db.Save(&request).Error; err != nil {
			e = err
			return
		}
	} else {
		if err := db.Db.Create(&request).Error; err != nil {
			e = err
			return
		}
	}
	res = request
	return
}

func (db *PostgresDb) FindAuthorizationRequestById(id int) (res models.Authorization, e error) {
	if err := db.Db.First(&res, id).Error; err != nil {
		e = err
		return
	}
	return
}

func (db *PostgresDb) FindAuthorizationRequestByUsername(username string) (res models.Authorization, e error) {
	if err := db.Db.Where("username = ?", username).First(&res).Error; err != nil {
		e = err
		return
	}
	return
}

func (db *PostgresDb) SaveSubscription(subscription models.Subscription) (res models.Subscription, e error) {
	if subscription.Id > 0 {
		if err := db.Db.Save(&subscription).Error; err != nil {
			e = err
			return
		}
	} else {
		if err := db.Db.Create(&subscription).Error; err != nil {
			e = err
			return
		}
	}
	res = subscription
	return
}
