package dao

import (
	"imdemo/imtwo/user/protocol"
	"log"
	"time"
)

var RD UsersRedis

type UsersRedis interface {
	AddUser(id int64,timeline *protocol.TimeLine)
	GetUsers()
	IsExist(userID int64) *protocol.TimeLine
}

type UsersRedisImpl struct {
	Users map[int64]*protocol.TimeLine
}

func (u *UsersRedisImpl)AddUser(id int64,timeline *protocol.TimeLine) {
	u.Users[id]=timeline
}

func (u *UsersRedisImpl)GetUsers()  {
	for  {
		time.Sleep(5*time.Second)
		for i, i2 := range u.Users {
			log.Println(i,i2)
		}

	}
}

//key userid value user of timeline
func (u *UsersRedisImpl)IsExist(userID int64) *protocol.TimeLine {
	timeline,ok:=u.Users[userID]
	if !ok {
		return nil
	}
	return timeline
}

func InitRedis() {
	RD= &UsersRedisImpl{Users: make(map[int64]*protocol.TimeLine)}
}