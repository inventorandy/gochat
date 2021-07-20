package controller

import (
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// Controller interface
type Controller interface {
	CreateUser(in *pb.User) (*pb.User, error)
	GetUserByID(in *wrappers.StringValue) (*pb.User, error)
	GetUserByEmail(in *wrappers.StringValue) (*pb.User, error)
	Login(in *pb.LoginRequest) (*pb.LoginResponse, error)
	AuthenticateJWT(in *wrappers.StringValue) (*pb.User, error)
	GetAllUsers(in *empty.Empty) (*pb.UserList, error)
}
