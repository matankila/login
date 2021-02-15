package endpoints

import (
	"com.poalim.bank.hackathon.login/global"
	"com.poalim.bank.hackathon.login/service"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Endpoint func(c echo.Context) error

type Endpoints struct {
	Login    Endpoint
	Register Endpoint
	Health   Endpoint
	Activate Endpoint
}

type LoginReq struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
}

type RegisterResp struct {
	Message string `json:"message"`
}

type ActivationResp struct {
	Message string `json:"message"`
}

type HealthResp struct {
	Ok bool `json:"ok"`
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Login:    makeLoginEndpoint(s),
		Register: makeRegisterEndpoint(s),
		Health:   makeHealthEndpoint(s),
		Activate: makeUserActivationEndpoint(s),
	}
}

var (
	registerSuccessfulResp = RegisterResp{
		Message: "account activation sent to mail",
	}

	activationSuccessfully = ActivationResp{
		Message: "account activated successfully",
	}
)

func makeLoginEndpoint(s service.Service) Endpoint {
	return func(c echo.Context) error {
		ctx := initContext(c)
		r := new(LoginReq)
		if err := c.Bind(r); err != nil {
			return err
		}

		jwt, err := s.Login(ctx, r.Mail, r.Password)
		if err != nil {
			return err
		}

		resp := LoginResp{
			Token: jwt,
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func makeRegisterEndpoint(s service.Service) Endpoint {
	return func(c echo.Context) error {
		ctx := initContext(c)
		r := new(RegisterReq)
		if err := c.Bind(r); err != nil {
			return err
		}

		err := s.Register(ctx, r.Mail, r.Password, r.FirstName, r.LastName)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, registerSuccessfulResp)
	}
}

func makeHealthEndpoint(s service.Service) Endpoint {
	return func(c echo.Context) error {
		ctx := initContext(c)
		resp := &HealthResp{Ok: true}

		err := s.Health(ctx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func makeUserActivationEndpoint(s service.Service) Endpoint {
	return func(c echo.Context) error {
		ctx := initContext(c)
		err := s.Activate(ctx, c.Param(global.UserId))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, activationSuccessfully)
	}
}

// create a context from request and init it with map containing important info about the request,
// such as: host, uri, ip, x-request-id, etc..
func initContext(c echo.Context) context.Context {
	ctx := c.Request().Context()
	m := global.GetRequestInfoData(c)
	ctx = context.WithValue(ctx, global.Mapping, m)

	return ctx
}
