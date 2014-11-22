package entity

import (
//	"time"
)

type Session struct {
	__table struct{} `db:"sessions"`

	Id     int64  `db:"id"`
	UserId int64  `db:"user_id"`
	ApiKey string `db:"api_key"`
	//Expires time.Time `db:"expires"`
}

func (self Session) GetId() int64 {
	return self.Id
}
