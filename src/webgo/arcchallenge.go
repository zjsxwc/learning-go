package main

import (
	"fmt"
	"github.com/hoisie/web"
	"math/rand"
	"time"
)

var form = `<form action="say" method="POST"><input name="said"><input type="submit"></form>`

var users = map[string]string{}

func main() {
	rand.Seed(time.Now().UnixNano())
	web.Config.CookieSecret = "8490sdfj489ufjijf984jojf94"
	web.Get("/", func(ctx *web.Context) string {
		ctx.Redirect(302, "/said")
		return ""
	})
	web.Get("/said", func() string { return form })
	web.Post("/say", func(ctx *web.Context) string {
		uid := fmt.Sprintf("%d\n", rand.Int63())
		ctx.SetSecureCookie("user", uid, 36000)
		users[uid] = ctx.Params["said"]
		return `<a href="/final">Click Here</a>`
	})
	web.Get("/final", func(ctx *web.Context) string {
		uid, _ := ctx.GetSecureCookie("user")
		return "You said " + users[uid]
	})
	web.Run(":9090")
}
