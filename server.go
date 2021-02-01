package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/twinj/uuid"
)

var ctx = context.Background()
var rdb *redis.Client

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	uuid := CreateAuth(rdb)
	FetchAuth(uuid, rdb)
	FetchAuth("strangecode", rdb)
}

func main() {
	ExampleClient()
}

func CreateAuth(rds *redis.Client) string {
	ExpiresAt := time.Now().Add(time.Minute * 10).Unix()
	at := time.Unix(ExpiresAt, 0) //converting Unix to UTC
	now := time.Now()
	uuid := uuid.NewV4().String()
	err := rds.Set(ctx, uuid, "1", at.Sub(now)).Err()
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
	return uuid
}

// 토큰에 저장된 uuid를 Redis에서 찾음
func FetchAuth(uuid string, rds *redis.Client) {
	id, err := rds.Get(ctx, uuid).Result()
	// redis 에 저장된 적이 없거나 토큰의 유효 기간이 만료된 경우 에러 발생
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		fmt.Println("키가 없는 값일 때 무엇을 리턴할까?")
		fmt.Printf("non key id -> %v", id)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("user id : ", id)
	}
	userID, _ := strconv.Atoi(id)
	fmt.Println("user id : ", userID)
}
