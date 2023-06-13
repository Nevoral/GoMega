package main

import (
	gop "GoMega/webHtml"
)

func main() {
	stat := gop.NewPage()
	stat.PageCode += gop.Html(
		gop.Lang("cs-CZ"),
		gop.Head("",
			gop.Meta(gop.Charset("UTF-8")),
			gop.Meta(gop.Atr(
				gop.Name("viewport"),
				gop.Content("width=device-width, initial-scale=1.0")),
			),
			gop.Link(gop.Atr(
				gop.Rel("stylesheet"),
				gop.Type("text/css"),
				gop.Href(stat.Name+".css")),
			),
			gop.Title("",
				"GoMegaApp",
			),
		),
		gop.Body("",

			gop.Script(gop.Atr(
				gop.Id("loader"),
				gop.Src("loader.js"),
			),
			),
		),
	)
	stat.SaveHTMLPage("./Pages/templates", stat.Name+".html")
}
