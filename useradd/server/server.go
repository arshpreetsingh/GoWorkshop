package server

import (
	"fmt"

	"github.com/wizelineacademy/GoWorkshop/proto/useradd"
	"github.com/wizelineacademy/GoWorkshop/useradd/models"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) AddUser(ctx context.Context, in *useradd.AddDataRequest) (*useradd.AddDataResponse, error) {
	userID, err := models.AddUser(in.Value)
	if err != nil {
		println("we also got error here as well", err)
	}
	response := new(useradd.AddDataResponse)
	response.Result = "hello"
	fmt.Println("this is userID", userID)
	fmt.Println("this is response", response)
	return response, err
}
