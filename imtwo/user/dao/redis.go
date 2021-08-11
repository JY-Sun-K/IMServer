package dao

import "imdemo/imtwo/user/timeline"

var Users =make(map[int64]*timeline.TimeLine)


func AddUser(id int64,timeline *timeline.TimeLine) {
	Users[id]=timeline
}