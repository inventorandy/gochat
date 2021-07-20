package controller

import (
	"fmt"
	"gochat/platform/internal/pbjson"
	"gochat/platform/internal/proto/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

// GetAllUsers fetches a list of all users on the system
func (c *AccountController) GetAllUsers(in *empty.Empty) (*pb.UserList, error) {
	// Call the Repository Method
	users, err := c.r.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Convert to Proto
	userList := &pb.UserList{}
	for _, user := range users {
		userPB := &pb.User{}
		if err := pbjson.ToProto(user, userPB); err != nil {
			return nil, fmt.Errorf("controller error: %s", err.Error())
		}
		userList.Users = append(userList.Users, userPB)
	}

	// Return the List
	return userList, nil
}
