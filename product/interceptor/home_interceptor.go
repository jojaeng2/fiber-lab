package interceptor

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStore = session.New()

func AuthInterceptor(ctx *fiber.Ctx) error {
	session, err := SessionStore.Get(ctx)
	if err != nil {
		return err
	}

	value := session.Get("user")
	fmt.Println(value)
	if value == nil {
		return ctx.Redirect("http://localhost:3000/123")
	}
	return ctx.Next()
}
