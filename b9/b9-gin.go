package b9

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Posts2 struct {
	UserId int    `json:"user_id" form:"user_id"`
	Id     int    `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
}

func B9Gin() {
	r := gin.Default()

	r.GET("/", HomeHandlerB9)
	r.POST("/", PostHandlerB9)

	r.Run(":3003")
}

func HomeHandlerB9(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	var p []Posts

	json.Unmarshal(data, &p)

	c.IndentedJSON(http.StatusOK, p)

}

func PostHandlerB9(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	defer resp.Body.Close()

	var p Posts2

	if c.ShouldBind(&p) == nil {
		log.Println(p.UserId)
		log.Println(p.Id)
		log.Println(p.Title)
		log.Println(p.Body)
	}

	c.String(http.StatusCreated, "ok")

}
