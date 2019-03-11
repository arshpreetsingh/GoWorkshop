package server

import (
	"fmt"

	"github.com/wizelineacademy/GoWorkshop/bounce/models"
	"github.com/wizelineacademy/GoWorkshop/proto/bounce"
	"golang.org/x/net/context"
)

type Server struct {
	db models.Datastore
}

func (s *Server) GetData(ctx context.Context, in *bounce.BounceGetData) (*bounce.BounceGetDataResponse, error) {
	bounces, err := s.db.BounceGetData()
	//return bounces, err
	response := new(bounce.BounceGetDataResponse)
	fmt.Println("this is userID", bounces)
	return response, err
}
