package dao

import (
	"imdemo/imtwo/user/timeline"
	"log"
	"time"
)

var RD UsersRedis

type UsersRedis interface {
	AddUser(id int64,timeline *timeline.TimeLine)
	GetUsers()
	IsExist(userID int64) *timeline.TimeLine
}

type UsersRedisImpl struct {
	Users map[int64]*timeline.TimeLine
}

func (u *UsersRedisImpl)AddUser(id int64,timeline *timeline.TimeLine) {
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
func (u *UsersRedisImpl)IsExist(userID int64) *timeline.TimeLine {
	timeline,ok:=u.Users[userID]
	if !ok {
		return nil
	}
	return timeline
}

func InitRedis() {
	RD= &UsersRedisImpl{Users: make(map[int64]*timeline.TimeLine)}
}