package main

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strings"

	"github.com/ultravioletasdf/messenger/backend/db"
	"github.com/ultravioletasdf/messenger/backend/pb"
	"golang.org/x/crypto/bcrypt"

	"github.com/badoux/checkmail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
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

	// check if password is valid
	if between := isBetween(len(in.Password), 8, 32); !between {
		return nil, status.Error(codes.InvalidArgument, "Password must be between 8 and 32 characters")
	}

	// check if email is taken
	email, err := executor.GetFromEmail(ctx, in.Email)
	if err != sql.ErrNoRows && err != nil {
		return nil, status.Error(codes.Internal, "Couldn't check email: "+err.Error())
	}
	if email.ID != 0 {
		return nil, status.Error(codes.AlreadyExists, "Email is already being used")
	}

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "Couldn't hash password: "+err.Error())
	}

	// generate id
	id := idGenerator.Generate().Int64()

	// generate unique username
	username := strings.Split(in.Email, "@")[0]
	usernameModified := username
	existing, err := executor.CheckUsername(ctx, usernameModified)
	if err != sql.ErrNoRows && err != nil {
		existing = 0
	}
	for existing != 0 { // while username is taken
		existing, err = executor.CheckUsername(ctx, usernameModified)
		if err != nil {
			break
		}
		usernameModified = username + fmt.Sprint(1+rand.Intn(10000-1))
	}

	err = executor.CreateUser(ctx, db.CreateUserParams{ID: id, Username: usernameModified, Email: in.Email, Password: string(passwordHash)})
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create user: "+err.Error())
	}
	return &pb.User{Id: id, Email: in.Email, Username: usernameModified, DisplayName: usernameModified, Bio: "This user hasn't set a bio"}, nil
}
func (*usersServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.User, error) {
	user, err := executor.GetUser(ctx, in.Id)
	if err == sql.ErrNoRows && err != nil {
		return nil, status.Error(codes.NotFound, "User Not Found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong: "+err.Error())
	}
	return &pb.User{Id: user.ID, Email: user.Email, Username: user.Username, DisplayName: user.DisplayName.String, Bio: user.Bio.String}, nil
}
func (*usersServer) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.Session, error) {
	token := generateToken(32)
	p, _ := peer.FromContext(ctx)
	ip := p.Addr.String()
	user, err := executor.GetFromEmail(ctx, in.Email)
	if err == sql.ErrNoRows && err != nil {
		return nil, status.Error(codes.NotFound, "No user has that email")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong: "+err.Error())
	}
	// Check if hash is valid
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Incorrect Password")
	}

	err = executor.CreateSession(ctx, db.CreateSessionParams{Token: token, UserID: user.ID, Platform: in.Platform, Ip: ip})
	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong: "+err.Error())
	}
	return &pb.Session{Token: token}, nil
}
func (*usersServer) SignOut(ctx context.Context, in *pb.AuthorizationRequest) (*pb.Empty, error) {
	err := executor.DeleteSession(ctx, in.Token)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}
