package redisDB

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

type Rdb struct {
	Client *redis.Client
}

type Mainerrr interface {
	StreamListner(bufCh chan string, stop chan struct{}, r DBaser)
}

type StreamListener interface {
	XRange(ctx context.Context, stream, start, stop string) ([]redis.XMessage, error)
	XRead(ctx context.Context, a *redis.XReadArgs) ([]redis.XStream, error)
}

type Storager interface {
	Get(ctx context.Context, key string) (string,error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string,error)
}

type DBaser interface {
	StreamListener
	Storager
}

func (r *Rdb) XRange(ctx context.Context, stream, start, stop string) ([]redis.XMessage, error) {
	return r.Client.XRange(ctx,stream,start,stop).Result()
}

func (r *Rdb) XRead(ctx context.Context, a *redis.XReadArgs) ([]redis.XStream, error) {
	return r.Client.XRead(ctx, a).Result()
}

func (r *Rdb) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *Rdb) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return r.Client.Set(ctx, key, value, expiration).Result()
}

func InitClient() Mainerrr {
	rdb := redis.NewClient(&redis.Options{
		Addr:    "192.168.10.205:6379",
		Password: "", // no password set
		DB:      0,  // use default DB
		DialTimeout: -1,
		ReadTimeout: -1, //TODO: понять хуй ли он стрим дропает через 5-10сек
	})

	return Rdb{Client: rdb}
}

func (rr Rdb) StreamListner(bufCh chan string, stop chan struct{}, r DBaser) {
	ctx := context.Background()
	lastId, err := r.Get(ctx, "LastID")
	if err != nil {
		log.Fatalln("ERR! ", err)
	}
	fmt.Println("LASTID: ", lastId)
	//read stream after last worked id
	prevStream, err := r.XRange(ctx, "TestStream", lastId, "+")
	if err != nil {
		log.Fatalln("ERR! ", err)
	}
	for _, message := range prevStream {
		fmt.Println("Stream", message.ID)

		bufCh <- message.ID

		for k, v := range message.Values {
			fmt.Println("Message K:", k, "val: ", v)
		}
	}

	//begin listen stream
	for {
		select {
		case <-stop:
			return
		default:
			reslt, err := r.XRead(ctx, &redis.XReadArgs{
				Streams: []string{"TestStream", "$"},
				Count:   1,
				Block:   0,
			})
			if err != nil{
				fmt.Println(err)
				break
			}

			for _, stream := range reslt {
				for _, message := range stream.Messages {
					fmt.Println("RESULT:", message.ID)

					bufCh <- message.ID
				}
			}
		}

	}

}


func writeLastID(r DBaser) {
	ctx := context.Background()
	lastId, err := r.Get(ctx, "LastID")
	if err != nil {
		log.Fatalln("ERR! ", err)
	}
	fmt.Println("LASTID: ", lastId)
}

func ChReader(bufCh chan string, num int, worker func(inp string))  {

	for {
		select {
		case rslt := <- bufCh:
			fmt.Println("jobber No", num, "reslt: ", rslt)
			worker(rslt)
		}
	}
}

func workk(inp string)  {
	fmt.Println("WORKER CALLED!", inp)
}

/*func CreateRedis() {
	bufCh := make(chan string, 4)
	stopCh := make(chan struct{})

	client := InitClient()

	for i := 0; i < 2; i++ {
		go ChReader(bufCh, i, workk)
	}

	streamListner(bufCh, stopCh, client)
	writeLastID(client)
}*/