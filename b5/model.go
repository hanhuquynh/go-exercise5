package b5

import (
	"log"

	"github.com/hanhuquynh/grpc/pb"
)

type User struct {
	Id          string
	UserId      string `xorm:"pk"`
	PartnerId   string
	AliasUserId string
	Phone       string
	Created     int64
	Updated_at  int64
}

func ConvertPbUser(pbUser *pb.UserPartner) *User {
	return &User{
		Id:          pbUser.Id,
		UserId:      pbUser.UserId,
		PartnerId:   pbUser.PartnerId,
		AliasUserId: pbUser.AliasUserId,
		Phone:       pbUser.Phone,
		Created:     pbUser.Created,
		Updated_at:  pbUser.UpdatedAt,
	}
}

func ConvertUserPb(u User) *pb.UserPartner {
	return &pb.UserPartner{
		Id:          u.Id,
		UserId:      u.UserId,
		PartnerId:   u.PartnerId,
		AliasUserId: u.AliasUserId,
		Phone:       u.Phone,
		Created:     u.Created,
		UpdatedAt:   u.Updated_at,
	}
}

var engine = ConnectDB()

func (u *User) Insert() error {

	_, err := engine.Table("user").Insert(u)

	if err != nil {
		log.Printf("Insert %+v  %v: ", u, err)
		return err
	}
	log.Printf("Insert user %+v successfully!", u)

	return nil
}

func Read() ([]User, error) {
	var user []User

	err := engine.Find(&user)

	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func (u *User) Update() error {
	_, err := engine.Where("user_id = ?", u.UserId).Update(u)
	if err != nil {
		log.Fatalf("Update user %+v: %v", u, err)
	}
	return nil
}

func (u *User) Delete() error {
	_, err := engine.Where("user_id = ?", u.UserId).Delete(u)
	if err != nil {
		log.Fatalf("Delete user %+v: %v", u, err)
	}
	return nil
}
