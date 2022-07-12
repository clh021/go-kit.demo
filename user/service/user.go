package service

import (
	"context"
	"demo/user/model"
	"fmt"
)

// UserService is the server API for User service.
type UserService interface {
	Create(context.Context, *model.CreateReq) (*model.CreateResp, error)
	Delete(context.Context, *model.DeleteReq) (*model.DeleteResp, error)
}

// NewUserService new service
func NewUserService() UserService {
	var svc = &userService{}
	{
		// middleware
	}
	return svc
}

// implement
type userService struct {}

func (s *userService) Create(ctx context.Context, req *model.CreateReq) (*model.CreateResp, error) {
	fmt.Println("Create")
	resp := model.NewCreateResp()
	resp.Code = 200
	resp.Msg = "success"
	resp.Data.Id = "12"
	resp.Data.Name = req.Name
	resp.Data.Age = req.Age
	return resp, nil
}

func (s *userService) Delete(context.Context, *model.DeleteReq) (*model.DeleteResp, error) {
	fmt.Println("Delete")
	return model.NewDeleteResp(), nil
}
