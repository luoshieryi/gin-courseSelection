package dao

import (
	"project/model"
	"project/util/helper"
	"time"
)

func CreateSession(userId int64) (string, error) {

	auth := helper.RandStr(32)

	session := model.Session{
		UserID:  userId,
		Auth:    auth,
		Expires: time.Now().Add(30 * 24 * time.Hour), //30天后过期
	}

	err := model.DB.Create(&session).Error

	return auth, err
}

func GetSessionByAuth(auth string) model.Session {
	session := model.Session{}

	model.DB.Find(&session, "auth = ?", auth)

	return session
}

func DeleteSessionByID(id int64) error {
	err := model.DB.Delete(&model.Session{ID: id}).Error
	return err
}

func DeleteSessionByUserID(userID int64) error {
	err := model.DB.Delete(&model.Session{UserID: userID}).Error
	return err
}
