package global

import "github.com/labstack/echo/v4"

func GetRequestInfoData(c echo.Context) map[string]string {
	m := map[string]string{}
	m[Id] = c.Request().Header.Get(echo.HeaderXRequestID)
	m[Host] = c.Request().Host
	m[Uri] = c.Request().RequestURI
	m[Method] = c.Request().Method

	return m
}
