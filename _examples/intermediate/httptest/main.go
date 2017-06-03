package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/sessions"
)

func main() {
	app := xeon.New()

	app.AttachSessionManager(sessions.New(sessions.Config{Cookie: "mysessionid"}))

	app.Get("/hello", func(ctx context.Context) {
		sess := ctx.Session()
		if !sess.HasFlash() {
			ctx.HTML("<h1> Unauthorized Page! </h1>")
			return
		}

		ctx.JSON(context.Map{
			"Message": "Hello",
			"From":    sess.GetFlash("name"),
		})
	})

	app.Post("/login", func(ctx context.Context) {
		sess := ctx.Session()
		if !sess.HasFlash() {
			sess.SetFlash("name", ctx.FormValue("name"))
		}

	})
	app.Run(xeon.Addr(":8080"))
}
