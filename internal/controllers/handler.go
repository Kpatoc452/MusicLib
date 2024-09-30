package controllers

import (
	"musiclib/internal/errors"
	"musiclib/internal/service"
)

type Handler struct{ 
	service *service.Service
	error errors.Error
}

func NewHandle(service *service.Service, error errors.Error) *Handler{
	return &Handler{
		service: service,
		error: error,
	}
}