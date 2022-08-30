package main

import (
	"context"

	"github.com/hanhuquynh/b3"
	rd "github.com/hanhuquynh/redis"
)

var (
	conn = rd.ConnRedis()
	ctx  = context.Background()
)

func main() {
	b3.Mux()
	// rd.Bai6(conn, ctx)
	//rd.Bai7(conn, ctx)
	// b8.Gin()
	// b9.Bai9()
	//b9.B9Gin()

	// b4.CreateTable()
}
