package main

import (
	"echo-gonertia-react-example/config"
	"echo-gonertia-react-example/handler"
	appMiddleware "echo-gonertia-react-example/middleware"
	"encoding/gob"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func main() {
	e := echo.New()
	i := config.InitInertia()
	e.HTTPErrorHandler = handler.Error(i).Error
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	e.Use(appMiddleware.CustomCSRF())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("a_very_long_random_string_to_secure_session"))))
	e.Use(appMiddleware.Inertia(i))

	e.GET("/", handler.Home().Index(i)).Name = "home"
	e.GET("/post-example", handler.PostExample().Index(i)).Name = "post-example"
	e.POST("/post-example", handler.PostExample().Post(i)).Name = "post-example-post"

	e.Logger.Fatal(e.Start(":1323"))
}

func init() {
	gob.Register(inertia.ValidationErrors{})
}
