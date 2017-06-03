package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/view"
)

func main() {
	app := xeon.New() // defaults to these

	// - standard html  | view.HTML(...)
	// - django         | view.Django(...)
	// - pug(jade)      | view.Pug(...)
	// - handlebars     | view.Handlebars(...)
	// - amber          | view.Amber(...)

	tmpl := view.HTML("./templates", ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	// default template funcs are:
	//
	// - {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// - {{ render "header.html" }}
	// - {{ render_r "header.html" }} // partial relative path to current page
	// - {{ yield }}
	// - {{ current }}
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
	app.AttachView(tmpl)

	app.Get("/", hi)

	// http://localhost:8080
	app.Run(xeon.Addr(":8080"), xeon.WithCharset("UTF-8")) // defaults to that but you can change it.
}

func hi(ctx context.Context) {
	ctx.ViewData("Title", "Hi Page")
	ctx.ViewData("Name", "Xeon") // {{.Name}} will render: Xeon
	// ctx.ViewData("", myCcustomStruct{})
	ctx.View("hi.html")
}
