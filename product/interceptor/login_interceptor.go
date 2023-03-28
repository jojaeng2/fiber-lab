package interceptor

import "github.com/gofiber/fiber/v2"

func LoginInterceptor(ctx *fiber.Ctx) error {
	session, err := SessionStore.Get(ctx)
	if err != nil {
		return err
	}
	session.Set("Key", ctx.Params("id"))
	if err := session.Save(); err != nil {
		panic("세션 저장소 문제 발생")
	}
	return ctx.Next()
}
