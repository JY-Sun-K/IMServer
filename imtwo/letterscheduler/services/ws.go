package services

import "imdemo/imtwo/letterscheduler/timeline"

var WS *timeline.WriteScheduler



func InitWs() {
	ws:=timeline.MakeWriteScheduler(2)
	WS=ws
	go WS.Run()
}
