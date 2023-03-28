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

	value := session.Get("Key")
	fmt.Println(value)
	return ctx.Next()
}
