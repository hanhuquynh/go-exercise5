package b9

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Posts struct {
	UserId int    `json:"user_id"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func Bai9() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/posts", CreatePosts)

	fmt.Println("Server running...")

	fmt.Println(http.ListenAndServe(":3002", nil))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://jsosnplaceholder.typicode.com/posts")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var posts []Posts

	err = json.Unmarshal(data, &posts)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	var post = Posts{
		UserId: 2001,
		Id:     2001,
		Title:  "Posts",
		Body:   "CreatePosts",
	}

	data, err := json.Marshal(post)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var p Posts

	err = json.Unmarshal(body, &p)

	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
