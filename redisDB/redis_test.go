package redisDB

import (
	"testing"
)

func Test_streamListener(t *testing.T) {

	client := InitRedis()

	StreamListener(client)

}
