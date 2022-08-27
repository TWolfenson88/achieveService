package redisDB

import (
	"context"
	"github.com/go-redis/redis/v9"
	"testing"
	"time"
)

type mockRedis struct {}

func (m mockRedis) XRange(ctx context.Context, stream, start, stop string) ([]redis.XMessage, error) {
	result := []redis.XMessage{
		{
			ID:     "10-0",
			Values: map[string]interface{}{"key1":"val1", "key2":"val2"},
		},
	}
	return result, nil
}

func (m mockRedis) XRead(ctx context.Context, a *redis.XReadArgs) ([]redis.XStream, error) {
	time.Sleep(1*time.Millisecond)
	result := []redis.XStream{{
		Stream:   "testStream",
		Messages: []redis.XMessage{
			{
				ID:     "10-0",
				Values: map[string]interface{}{"key1":"val1", "key2":"val2"},
			},
		},
	}}
	return result, nil
}

func (m mockRedis) Get(ctx context.Context, key string) (string, error) {
	return "getted", nil
}

func (m mockRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return "setted", nil
}

func TestRdb_WriteLastID(t *testing.T) {
	testRedis := &mockRedis{}

	writeLastID(testRedis)
}

func Test_streamListner(t *testing.T) {

	mockRed := &mockRedis{}
	bufCh := make(chan string, 4)
	stopCh := make(chan struct{})
	for i := 0; i < 2; i++ {
		go ChReader(bufCh, i, workk)
	}
	go func() {
		time.Sleep(10*time.Millisecond)
		stopCh <- struct{}{}
	}()

	t.Run("listen", func(t *testing.T) {
		streamListner(bufCh,stopCh, mockRed)

	})
}