package main

import (
	"com.poalim.bank.hackathon.login/endpoints"
	"com.poalim.bank.hackathon.login/global"
	"com.poalim.bank.hackathon.login/service"
	"com.poalim.bank.hackathon.login/transport"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	l = log.New(global.LoggerName)
	services = service.NewService(l)
	eps = endpoints.MakeEndpoints(services)
	server = transport.NewServer(eps)
)

func main() {
	errs := make(chan error)

	// handle signals
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <- c)
	}()

	// server setup
	go func() {
		e := echo.New()
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: strings.Replace(middleware.DefaultLoggerConfig.Format, global.OldTemplate, global.NewTemplate, 1),
		}))
		e.Logger.SetLevel(log.ERROR)
		e.Logger.SetPrefix(global.LoggerName)
		server.Init(e)
		errs <- e.Start(":" + global.Port)
	}()

	l.Fatal("exit due to: ", <- errs)
}
