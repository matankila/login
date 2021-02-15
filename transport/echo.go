package transport

import (
	"com.poalim.bank.hackathon.login/endpoints"
	error2 "com.poalim.bank.hackathon.login/error"
	"com.poalim.bank.hackathon.login/global"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type server struct {
	eps endpoints.Endpoints
}

type Server interface {
	Init(interface{})
}

func NewServer(eps endpoints.Endpoints) Server {
	return &server{
		eps: eps,
	}
}

// init framework endpoints & custom error handler
func (s server) Init(framework interface{}) {
	var e *echo.Echo

	e, ok := framework.(*echo.Echo)
	if !ok {
		panic("something went wrong, not right framework")
	}

	e.HTTPErrorHandler = CustomErrorHandler
	e.POST("/auth/login", echo.HandlerFunc(s.eps.Login))
	e.POST("/auth/register", echo.HandlerFunc(s.eps.Register))
	e.POST("/auth/activate/user/:userID", echo.HandlerFunc(s.eps.Activate))
	e.GET("/health", echo.HandlerFunc(s.eps.Health))
}

// This error handler write to log & returns error response to user.
func CustomErrorHandler(err error, c echo.Context) {
	resp := map[string]interface{}{}
	code := http.StatusInternalServerError
	if he, ok := err.(*error2.HTTPError); ok {
		code = he.Code
	}

	m := global.GetRequestInfoData(c)
	c.Logger().Errorj(log.JSON{global.RequestInfo: m, global.ErrMessage: err.Error()})
	resp["error"] = err.Error()
	if err := c.JSON(code, resp); err != nil {
		panic(err)
	}
}
