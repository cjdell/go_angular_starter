package services

import (
	"crypto/md5"
	"encoding/hex"
	"go_angular_starter/model/entity"
	"go_angular_starter/model/persister"
)

func RegisterUser(db persister.DB, email string, password string) error {
	userPersister := persister.NewUserPersister(db)

	user := &entity.User{}

	hasher := md5.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	user.Email = email
	user.Hash = hash

	err := userPersister.Insert(user)

	if err != nil {
		return err
	}

	return nil
}
