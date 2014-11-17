package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
	"math/rand"
	"strconv"
	"time"
)

type AuthService struct {
	db persister.DB
}

func NewAuthService(db persister.DB) *AuthService {
	return &AuthService{db}
}

func (self AuthService) RegisterUser(email string, password string) (*entity.User, error) {
	userPersister := persister.NewUserPersister(self.db)

	user := &entity.User{}

	hasher := md5.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	user.Name = "Unnamed"
	user.Email = email
	user.Hash = hash

	_, err := userPersister.Insert(user)

	return user, err
}

type SignInResponse struct {
	ApiKey string
	Email  string
	Name   string
	Type   string
}

func (self AuthService) SignIn(email string, password string) (*SignInResponse, error) {
	userPersister := persister.NewUserPersister(self.db)
	sessionPersister := persister.NewSessionPersister(self.db)

	hasher := md5.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	user, err := userPersister.GetByEmailAndHash(email, hash)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	rand.Seed(time.Now().UTC().UnixNano())

	// Generate an API key, this could probably be written better
	hasher = md5.New()
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	hasher.Write([]byte(strconv.FormatInt(rand.Int63(), 10)))
	apiKey := hex.EncodeToString(hasher.Sum(nil))

	session := &entity.Session{}
	session.UserId = user.Id
	session.ApiKey = apiKey
	err = sessionPersister.Insert(session)

	if err != nil {
		return nil, err
	}

	res := &SignInResponse{apiKey, user.Email, user.Name, "Admin"}

	return res, err
}
