package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	rd "github.com/hanhuquynh/redis"
)

var (
	connRedis = rd.ConnRedis()
	ctx       = context.Background()
)

type Name struct {
	Name string
}

func main() {
	route := gin.Default()

	route.GET("/names", getAllData)

	route.GET("/name/:index", getDataByIndex)

	route.DELETE("/name/:index", deleteData)

	route.POST("/name", postData)

	route.PUT("/name/:index", updateData)

	route.Run(":3000")
}

func getAllData(c *gin.Context) {
	count, _ := connRedis.LLen(ctx, "names").Result()

	for i := 0; i < int(count); i++ {
		data, _ := connRedis.LIndex(ctx, "names", int64(i)).Result()

		c.String(http.StatusOK, strconv.Itoa(i)+", "+data+"\n")
	}

}

func getDataByIndex(c *gin.Context) {
	index := c.Param("index")

	convIndex, _ := strconv.Atoi(index)

	data, err := connRedis.LIndex(ctx, "names", int64(convIndex)).Result()

	if err != nil {
		c.String(http.StatusOK, "Index doesn't exist")
		return
	}

	c.String(http.StatusOK, data)

}

func deleteData(c *gin.Context) {
	index := c.Param("index")

	convIndex, _ := strconv.Atoi(index)

	data, err := connRedis.LIndex(ctx, "names", int64(convIndex)).Result()

	if err != nil {
		c.String(http.StatusOK, "Index doesn't exist")
		return
	} else {
		_, err := connRedis.LRem(ctx, "names", 1, data).Result()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "err: " + err.Error(),
			})
		}

		c.String(http.StatusOK, "Delete index "+index+" successfully")
	}

}

func postData(c *gin.Context) {
	name := Name{
		Name: "Demo",
	}

	data, err := json.Marshal(name)

	if err != nil {
		c.String(http.StatusOK, "co loi: "+err.Error())
	}

	_, err = connRedis.LPush(ctx, "names", data).Result()

	if err != nil {
		c.String(http.StatusOK, "co loi: "+err.Error())
		return
	}

	c.String(http.StatusCreated, string(data)+" created")
}

func updateData(c *gin.Context) {
	index := c.Param("index")

	convIndex, _ := strconv.Atoi(index)

	_, err := connRedis.LIndex(ctx, "names", int64(convIndex)).Result()
	if err != nil {
		c.String(http.StatusOK, "Index doesn't exist")
		return
	} else {
		_, err := connRedis.LSet(ctx, "names", int64(convIndex), "Update data").Result()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "err: " + err.Error(),
			})
		}

		c.String(http.StatusOK, "Update index "+index+" successfully")
	}
}
