package redisDB

import (
	"awesomeProject5/db"
	"awesomeProject5/logic"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"os"
	"strconv"
	"time"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func InitRedis() *redis.Client {
	var (
		host     = getEnv("REDIS_HOST", "localhost")
		port     = getEnv("REDIS_PORT", "6379")
		password = getEnv("REDIS_PASSWORD", "")
	)

	client := redis.NewClient(&redis.Options{
		Addr:        host + ":" + port,
		Password:    password,
		DB:          0,
		DialTimeout: 500 * time.Second,
		ReadTimeout: -1,
	})

	reslt, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("PING ERR", err)
	}

	log.Println("PING SUCESS ", reslt)

	return client
}

type Respons struct {
	UserId    int
	Location  int
	Timestamp time.Time
}

func StreamListener(rc *redis.Client, users map[int]*logic.User, pg db.Saver, logCh chan string) {
	ctx := context.Background()
	lastId, err := rc.Get(ctx, "LastID").Result()
	if err != nil {
		log.Println("GET ERR! ", err) //TODO: обработать момент если в ключе пусто
	}
	fmt.Println("LASTID: ", lastId)
	//read stream after last worked id
	prevStream, err := rc.XRange(ctx, "TestStream", lastId, "+").Result()
	if err != nil {
		log.Println("XRANGE ERR! ", err)
	}
	for _, message := range prevStream {
		fmt.Println("Stream", message.ID)

		//bufCh <- message.ID

		for k, v := range message.Values {
			fmt.Println("Message K:", k, "val: ", v)
		}
	}

	rsp := Respons{}

	//begin listen stream
	for {
		reslt, err := rc.XRead(ctx, &redis.XReadArgs{
			Streams: []string{"TestStream", "$"},
			Count:   1,
			Block:   0,
		}).Result()
		if err != nil {
			log.Println("XREAD ERR : ", err)
		}
		for _, stream := range reslt {
			for _, msg := range stream.Messages {
				fmt.Println("CYCLE MSGG", msg.ID)

				loc, err := strconv.Atoi(msg.Values["location"].(string))
				userId, err := strconv.Atoi(msg.Values["userId"].(string))
				tim, err := strconv.Atoi(msg.Values["time"].(string))

				if err != nil {
					log.Println("PARSE ERR : ", err)
				}

				fmt.Println(loc, err)

				rsp.UserId = userId
				rsp.Location = loc
				rsp.Timestamp = time.Unix(int64(tim), 0)

				fmt.Println("FILLED STRUCT!", rsp)

				if userr, ok := users[rsp.UserId]; !ok {

					usr := &logic.User{
						Id:              userId,
						UsrLvl:          1,
						TempAchieves:    map[int]*logic.UserAchieve{},
						CurrentAchieves: map[int]*logic.UserAchieve{},
					}

					usr.AddAchieve(rsp.Timestamp, rsp.Location, logCh)

					pg.SaveUserData(*usr)

					users[userId] = usr
				} else {
					userr.AddAchieve(rsp.Timestamp, rsp.Location, logCh)

					pg.SaveUserData(*userr)
				}

				rc.Set(context.Background(), "LastID", msg.ID, 0)

			}
		}

	}
}
