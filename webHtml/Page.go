package webJs

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Page struct {
	StaticPath   string
	RelativePath string
	PageCode     string
}

func NewPage() (page *Page) {
	return &Page{
		StaticPath:   "",
		RelativePath: "",
		PageCode:     Doctype,
	}
}

func (p *Page) RenderHTMLPage() []byte {
	return []byte(p.PageCode)
}

func (p *Page) SaveHTMLPage(path, filename string) []byte {
	if !strings.Contains(filename, ".html") {
		filename += ".html"
	}
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
	html := p.RenderHTMLPage()
	_, err = file.WriteString(string(html))
	if err != nil {
		log.Printf("File coudn't be write: %e", err)
	}
	fmt.Println("HTML file generated successfully!")
	return html
}

type Prefetch struct {
	filename string
	asType   string
}

func PrefetchJs(files []string) (pref string) {
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
