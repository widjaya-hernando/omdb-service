package models

import (
	pb "github.com/widjaya-hernando/prod-trans-user-service/gen"
)

type User struct {
	ID        uint64 `json:"id"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	//Products  []Product `json:"products" gorm:"foreignKey:user_id"`
}

func UserToGrpcMessage(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		UserName: user.UserName,
		Phone:    user.Phone,
		Email:    user.Email,
		Address:  user.Address,
	}
}
