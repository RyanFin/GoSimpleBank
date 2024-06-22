package gapi

import (
	db "RyanFin/GoSimpleBank/db/sqlc"
	"RyanFin/GoSimpleBank/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// convert the user from the database into the gRPC protobuf user equivalent
func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
