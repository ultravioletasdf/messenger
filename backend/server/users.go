package main

import (
	"context"
	"strings"

	"github.com/ultravioletasdf/messenger/backend/pb"

	"github.com/badoux/checkmail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type usersServer struct {
	pb.UnimplementedUsersServer
}

func (*usersServer) GetMe(ctx context.Context, in *pb.AuthorizationRequest) (*pb.User, error) {
	if in.Token == "abc" {
		return &pb.User{Id: 12345, Username: "myuser", DisplayName: "My Displa Name", Bio: "this is a bio"}, nil
	}

	return nil, status.Error(codes.NotFound, "User Not Found")
}

func (*usersServer) Create(ctx context.Context, in *pb.CreateRequest) (*pb.User, error) {
	// check if email is valid
	if err := checkmail.ValidateFormat(in.Email); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid email: "+err.Error())
	}

	// check if email is taken
	// hash password
	// create user in database
	username := in.Email[strings.LastIndex(in.Email, "@")+1:]
	return &pb.User{Id: 12345, Username: username, DisplayName: username, Bio: "this is a bio"}, nil
}
