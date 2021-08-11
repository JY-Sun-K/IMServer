package dao

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"imdemo/imone/common"
)

var RDB *redis.Client

func InitRedis()  {
	rdb :=redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
	RDB =rdb

}

type RedisHub interface {
	RegisterClient(client *common.Client)error
	UnRegisterClient(userId string)error
	QueryUserOnline(userId string)(bool,error)
}


type RedisHubImpl struct {

}





func (r *RedisHubImpl) RegisterClient(client *common.Client)error {
	cJson, err := json.Marshal(client)
	if err != nil {
		return err
	}
	err=RDB.Set(context.Background(),client.UserId,cJson,0).Err()
	if err!=nil {
		return err
	}
	return nil
}

func (r *RedisHubImpl) UnRegisterClient(userId string)error{
	err:=RDB.Del(context.Background(),userId).Err()
	if err!=nil {
		return err
	}
	return nil
}

func (r *RedisHubImpl) QueryUserOnline(userId string)(bool,error) {
	return true ,nil
}
