package handler

import "github.com/nikitsingh/forky/backend/internal/service"

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}
