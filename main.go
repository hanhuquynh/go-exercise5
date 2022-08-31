package main

import (
	"context"

	rd "github.com/hanhuquynh/redis"
)

var (
	conn = rd.ConnRedis()
	ctx  = context.Background()
)

func main() {
	// rd.Bai6(conn, ctx)
	// rd.Bai7(conn, ctx)
	// b8.Gin()
	// b9.Bai9()
	//b9.B9Gin()

	// b5.CreateTable()
}
