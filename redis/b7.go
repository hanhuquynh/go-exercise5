package redis

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func Bai7(conn *redis.Client, ctx context.Context) {
	file, err := os.Open("redis/name.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		conn.LPush(ctx, "names", scanner.Text())
	}

	// log.Println(conn.LPop(ctx, "names"))

	// conn.Expire(ctx, "names", 5*time.Second)

	// conn.RPush(ctx, "names", "Dota2vn")

	// log.Println(conn.LRange(ctx, "names", 0, -1))
}
