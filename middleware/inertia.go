package middleware

import (
	"echo-gonertia-react-example/lib"
	"net/http"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func Inertia(i *inertia.Inertia) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			ctx := lib.NewEchoContext(req.Context(), c)
			ctx = inertia.SetProps(ctx, inertia.Props{
				"csrf_token": c.Get("csrf"),
			})
			i.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c.SetRequest(r)
				err := next(c)
				if err != nil {
					c.Error(err)
				}
			})).ServeHTTP(res, req.WithContext(ctx))

			return nil
		}
	}
}
