package protocol

type UserInfo struct {
	UserId int64 `json:"user_id"`
	UserName string	`json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}


//map[userid]username
type FriendsInfo struct {
	UserId int64 `json:"user_id"`
	Friends map[int64]string
}


type SearchUserInfo struct {
	UserId int64`json:"user_id"`
	UserName string `json:"user_name"`
}

//type SearchUserResponse struct {
//	SearchUsers []*SearchUserInfo `json:"search_users"`
//}

//type UFriend struct {
//	UserId int64
//	FriendId int64
//}
type Letter struct {
	IPAddress string //收件人所在服务器ip
	From int64 //发件人
	To int64 // 收件人
	Message string //消息
	SendTime string //发送时间
}



//每个人所有的一个Timeline
type TimeLine struct {
	Owner int64
	InBox chan *Letter
}

