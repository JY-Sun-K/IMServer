package main

import (
	"fmt"
	"imdemo/imone/common"
	"imdemo/imone/dao"
	"imdemo/imone/services"
	"imdemo/imone/transports"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	common.MsgChan = make(chan string)
	errChan := make(chan error)
	dao.InitRedis()
	wsService := services.MakeWSServiceImpl(&dao.RedisHubImpl{})
	r:=transports.MakeHttpHandler(wsService)
	go func() {
		errChan <- http.ListenAndServe(":8080", r)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		for  {
			log.Println(<-common.MsgChan)
		}
	}()

	error := <-errChan
	log.Println(error)
}
