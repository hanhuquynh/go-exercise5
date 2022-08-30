package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func Bai6(conn *redis.Client, ctx context.Context) {
	t := time.Now().Unix()

	expireTime := 10 * time.Second

	err := conn.Set(ctx, "demo_key", t, expireTime).Err()

	if err != nil {
		fmt.Println(err)
	}

	v, err := conn.Get(ctx, "demo_key").Result()

	if err == redis.Nil {
		fmt.Println("Key doesn't exist!")
	} else {
		fmt.Println("Value: ", v)
	}

	time.Sleep(12 * time.Second)

	v1, err := conn.Get(ctx, "demo_key").Result()

	if err == redis.Nil {
		timeNow := time.Now().Second()
		err = conn.Set(ctx, "demo_key", timeNow, 0).Err()
		if err != nil {
			fmt.Println(err)
		}

		v2, err := conn.Get(ctx, "demo_key").Result()

		if err == redis.Nil {
			fmt.Println("Key doesn't exist!")
		} else {
			fmt.Println("Value after 12s: ", v2)
		}
	} else {
		fmt.Println("Value: ", v1)
	}
}
