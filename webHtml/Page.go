package webHtml

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Page struct {
	Name         string
	StaticPath   string
	RelativePath string
	PageCode     string
}

func NewPage() (page *Page) {
	return &Page{
		Name:         "",
		StaticPath:   "",
		RelativePath: "",
		PageCode:     Doctype,
	}
}

func (p *Page) RenderHTMLPageByte() []byte {
	return []byte(p.PageCode)
}

func (p *Page) RenderHTMLPageString() string {
	return p.PageCode
}

func (p *Page) SaveHTMLPage(path, filename string) {
	filenameB := filename
	if !strings.Contains(filename, ".html") {
		filename += ".html"
	} else {
		filenameB = filenameB[:len(filenameB)-6]
	}
	fmt.Println(filenameB)
	if len(path) != 0 {
		if string(path[len(path)-1]) != "/" && string(path[len(path)-1]) != "\\" {
			path += "/"
		}
	}
	file, err := os.Create(path + filename)

	if err != nil {
		log.Printf("File wasn't created: %e", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	_, err = file.WriteString(p.RenderHTMLPageString())
	if err != nil {
		log.Printf("File coudn't be write: %e", err)
	}
	GenerateTailwindCss(filename, path+"input.css", path+filename+".css", false)
	fmt.Println("HTML file generated successfully!")
}

func (p *Page) BasicSkeletonSite(PageName, title string) {
	p.Name = PageName
	p.PageCode += Html(
		Lang("cs-CZ"),
		Head("",
			Meta(Charset("UTF-8")),
			Meta(Atr(
				Name("viewport"),
				Content("width=device-width, initial-scale=1.0")),
			),
			Link(Atr(
				Rel("stylesheet"),
				Type("text/css"),
				Href(p.Name+".css")),
			),
			Title("",
				"GoMegaApp",
			),
			//Script(Src("https://cdn.tailwindcss.com")),
		),
		Body("",

			Script(Atr(
				Id("loader"),
				Src("loader.js"),
			),
				PrefetchJs(""),
				//Script(Type("qwick/json"),
				//	"{\"refs\": {}, \"ctx\": {}, \"objs\": [], \"subs\": []}",
				//),
				//Script("",
				//	"window.qwikevents.push(\"click\")",
				//),
			),
		),
	)
}

type Prefetch struct {
	filename string
	asType   string
}

func PrefetchJs(files ...string) (pref string) {
	for _, val := range files {
		pref += Link(Atr(
			Rel("prefetch"),
			Href(val),
			As("script"),
		))
	}
	return
}
func SerializedState() string {
	return fmt.Sprintf("")
}

//todo BASIC TEMPLATE HTML SIDE ALL NEEDED ATOMIC TAGS
