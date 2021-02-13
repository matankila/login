package service

import (
	error2 "com.poalim.bank.hackathon.login/error"
	"com.poalim.bank.hackathon.login/global"
	"context"
	"github.com/labstack/gommon/log"
	"net/http"
)

type service struct {
	log *log.Logger
}

type Service interface {
	Login(ctx context.Context, mail, password string) (string, error)
	Register(ctx context.Context, mail, password, firstName, lastName string) error
	Health(ctx context.Context) error
    Activate(ctx context.Context, userId string) error
}

func NewService(log *log.Logger) Service {
	return &service{
		log: log,
	}
}

func (s service) Login(ctx context.Context, mail, password string) (string, error) {
	m := ctx.Value(global.Mapping).(map[string]string)

	s.log.Infoj(log.JSON{global.RequestInfo: m, global.Message: "trying to login..."})
	if mail == "test" {
		return "", error2.NewHTTPError(http.StatusBadRequest, "'test' as mail is wrong")
	}

	// TODO: add business logic
	return "DEADBEEF", nil
}

func (s service) Register(ctx context.Context, mail, password, firstName, lastName string) error {
	m := ctx.Value(global.Mapping).(map[string]string)
	s.log.Infoj(log.JSON{global.RequestInfo: m, global.Message: "user trying to register..."})
	// TODO: add business logic
	return nil
}

func (s service)Health(ctx context.Context) error {
	m := ctx.Value(global.Mapping).(map[string]string)
	s.log.Infoj(log.JSON{global.RequestInfo: m, global.Status: global.Ok})
	// TODO: add business logic
	return nil
}

func (s service)Activate(ctx context.Context, userId string) error {
	m := ctx.Value(global.Mapping).(map[string]string)
	s.log.Infoj(log.JSON{global.RequestInfo: m, global.Status: global.Ok})
	// TODO: add business logic
	return nil
}