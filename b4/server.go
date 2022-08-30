package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	b5 "github.com/hanhuquynh/b5"
	"github.com/rs/xid"
)

var engine = b5.ConnectDB()

func main() {
	r := mux.NewRouter()
	port := 3001

	r.HandleFunc("/user-partner", GetUser).Methods(http.MethodGet)

	r.HandleFunc("/user-partner", PostUser).Methods(http.MethodPost)

	r.HandleFunc("/user-partner/{id}", DeleteUser).Methods(http.MethodDelete)

	r.HandleFunc("/user-partner/{id}", GetUserById).Methods(http.MethodGet)

	log.Printf("Server running on: "+"localhost:%v", port)

	log.Println(http.ListenAndServe(":"+strconv.Itoa(port), r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var u []b5.User

	err := engine.Find(&u)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	u := b5.User{
		Id:          xid.New().String(),
		UserId:      "5",
		PartnerId:   "5",
		AliasUserId: "5",
		Phone:       "0988776653",
		Created:     time.Now().UnixMilli(),
		Updated_at:  time.Now().UnixMilli(),
	}
	_, err := engine.Insert(&u)

	if err != nil {
		fmt.Fprintf(w, "Insert user: %v err %v", u, err)
	}
	fmt.Fprintf(w, "Insert user %+v successfully", u)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := string(params["id"])

	var u b5.User
	exist, _ := engine.Table("user").Exist(&b5.User{
		Id: id,
	})

	if exist {
		_, err := engine.Where("id = ?", id).Delete(&u)
		if err != nil {
			fmt.Fprintf(w, "err: %v", err)
		}
		fmt.Fprintf(w, "Delete user id: %v successfully", id)
	} else {
		fmt.Fprintf(w, "ID doesn't exist")
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := string(params["id"])
	var u b5.User
	exist, _ := engine.Table("user").Exist(&b5.User{
		Id: id,
	})

	if exist {
		_, err := engine.Where("id = ?", id).Get(&u)
		if err != nil {
			fmt.Fprintf(w, "err: %v", err)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	} else {
		fmt.Fprintf(w, "ID doesn't exist")
	}

}
