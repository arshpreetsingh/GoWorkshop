package server

import (
	"github.com/wizelineacademy/GoWorkshop/proto/users"
	"github.com/wizelineacademy/GoWorkshop/useradd/models"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) AddUser(ctx context.Context, in *users.AddDataRequest) (*users.AddDataResponse, error) {
	userID, err := models.AddUser(in.value)
	response := new(users.AddDataResponse)
	response.result = "hello"
}
