package gapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	db "github.com/wsmbsbbz/simplebank/db/sqlc"
	"github.com/wsmbsbbz/simplebank/pb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreateAt:          timestamppb.New(user.CreateAt),
	}
}
