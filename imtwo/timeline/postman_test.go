package timeline

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	ws := MakeWriteScheduler(2)


	userA:=MakeTimeLine("小明")
	userB:=MakeTimeLine("小红")
	go ws.Run()
	AddUser("小明",userA)
	AddUser("小红",userB)
	l:=MakeLetter("小明","小红","helloWorld")
	ws.WriteChan<- l
	a:=<-userB.InBox
	fmt.Println(a)

}
