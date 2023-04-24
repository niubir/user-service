package grpc

import (
	"github.com/niubir/user-service/grpc/server"
)

type UserService struct {
	server.UnsafeUserServiceServer
}
