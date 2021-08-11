package dao

import (
	"errors"
	"gorm.io/gorm"
	"imdemo/imtwo/user/protocol"
)

type UserDAO interface {
	IsExistUser(email string) (*protocol.UserInfo,error)
	CreateUser(userName string,hasedPassword string,email string) (int64,error)
	GetUserInfo(userid int64)(*protocol.FriendsInfo,error)
	SearchUser(userName string)([]*protocol.SearchUserInfo,error)
	AddUser(userId int64,friendId int64)error
}


type UserDAOImpl struct {
}

func (u *UserDAOImpl) IsExistUser(email string) (*protocol.UserInfo,error) {
	user:= User{}
	err:= DB.Where("email = ?", email).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil,err
	}
	if user.UserId==0 {
		return nil,nil
	}
	return &protocol.UserInfo{
		UserId:   user.UserId,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	},nil
}

func (u *UserDAOImpl)CreateUser(userName string,hasedPassword string,email string) (int64,error) {
	user := User{
		UserName:  userName,
		Email:     email,
		Password:  hasedPassword,
	}
	err:= DB.Create(&user).Error
	if err != nil {
		return 0,err
	}
	err= DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return 0,err
	}

	return user.UserId,nil
}

func (u *UserDAOImpl) GetUserInfo(userid int64)(*protocol.FriendsInfo,error) {
	friends:=make(map[int64]string)
	rows, err := DB.Table("friends").Select("friends.friend_id, users.user_name").Joins("join users on users.user_id = friends.friend_id").Where("friends.user_id=?",userid).Rows()
	for rows.Next() {
		var userId int64
		var userName string
		_ = rows.Scan(&userId, &userName)
		if _,ok:=friends[userId];!ok{
			friends[userId]=userName
		}else {
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return &protocol.FriendsInfo{
		UserId:  userid,
		Friends: friends,
	},nil

}

func (u *UserDAOImpl) SearchUser(userName string)([]*protocol.SearchUserInfo,error){
	users:=make([]*protocol.SearchUserInfo,0)
	rows,err:= DB.Table("users").Select("user_id,user_name").Where("user_name LIKE ?","%"+userName+"%").Rows()
	for rows.Next() {
		var userId int64
		var userName string
		_ = rows.Scan(&userId, &userName)
		user:=&protocol.SearchUserInfo{
			UserId:userId,
			UserName: userName,
		}
		users=append(users,user)

	}
	if err != nil {
		return nil, err
	}
	return users,nil
}

func (u *UserDAOImpl)AddUser(userId int64,friendId int64)error  {

	friends:=[]Friends{{UserId: userId,FriendId: friendId},{UserId: friendId,FriendId: userId}}
	err:= DB.Create(&friends).Error
	if err != nil {
		return err
	}
	return nil
}