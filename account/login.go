package account

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lilihx/chatRoom/common/config"
	"net/http"
)

func index(c echo.Context) error {
	html := "<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\">\n    <title>title</title>\n</head>\n<body class=\"body\">\n    <a href=\"https://github.com/login/oauth/authorize?client_id=%v&redirect_uri=%v\"> Github登录 </a>\n</body>\n</html>"
	s := fmt.Sprintf(html, config.Config.Github.Oauth.ClientId, config.Config.Github.Oauth.RedirectUrl)
	return c.HTML(http.StatusOK, s)
}

func oauth(c echo.Context) error {
	code := c.QueryParam("code")
	return c.String(http.StatusOK, code)
}
