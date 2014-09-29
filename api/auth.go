package api

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
	"github.com/cjdell/go_angular_starter/services"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type AuthApi struct {
	db               persister.DB
	persister        *persister.UserPersister
	sessionPersister *persister.SessionPersister
}

func NewAuthApi(db *sqlx.DB) *AuthApi {
	authService := &AuthApi{}

	authService.db = db
	authService.persister = persister.NewUserPersister(db)
	authService.sessionPersister = persister.NewSessionPersister(db)

	return authService
}

type AuthSignInArgs struct {
	Email    string
	Password string
}

type AuthSignInReply struct {
	ApiKey string
}

func (self *AuthApi) SignIn(r *http.Request, args *AuthSignInArgs, reply *AuthSignInReply) error {
	hasher := md5.New()
	hasher.Write([]byte(args.Password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	user, err := self.persister.GetByEmailAndHash(args.Email, hash)

	if user == nil {
		return errors.New("User not found")
	}

	rand.Seed(time.Now().UTC().UnixNano())

	// Generate an API key, this could probably be written better
	hasher = md5.New()
	hasher.Write([]byte(user.Email))
	hasher.Write([]byte(args.Password))
	hasher.Write([]byte(strconv.FormatInt(rand.Int63(), 10)))
	reply.ApiKey = hex.EncodeToString(hasher.Sum(nil))

	session := &entity.Session{}
	session.UserId = user.Id
	session.ApiKey = reply.ApiKey
	self.sessionPersister.Insert(session)

	if err != nil {
		return err
	}

	return nil
}

type AuthRegisterArgs struct {
	Email    string
	Password string
}

type AuthRegisterReply struct {
}

func (self *AuthApi) Register(r *http.Request, args *AuthRegisterArgs, reply *AuthRegisterReply) error {
	return services.RegisterUser(self.db, args.Email, args.Password)
}
