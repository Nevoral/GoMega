package webHtml

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ParserGOMG(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file contents
	htmlBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	html := make([]byte, len(htmlBytes))

	// Copy the elements from the slice to the array
	copy(html, htmlBytes)
	findTag(html)

	// Convert the bytes to a string
	//html := string(htmlBytes)

	// Print the HTML contents
	//fmt.Println(html)

}

type Htmltag struct {
	counter int
	start   int
	stop    int
	tag     string
}

type rowHtml struct {
	row     []byte
	typeRow uint8
}

func findTag(html []byte) (value string) {
	var filterdChar = char{10, 13}
	var tag []rowHtml
	for ind := 0; ind < len(html); ind++ {
		if filterdChar.Contains(html[ind]) {
			continue
		}
		if html[ind] == byte('<') {
			if ind == 0 || len(tag[len(tag)-1].row) != 0 {
				tag = append(tag, rowHtml{
					row:     []byte{},
					typeRow: 1,
				})
			} else {
				tag[len(tag)-1].typeRow = 1
			}
			tag[len(tag)-1].row = append(tag[len(tag)-1].row, html[ind])
			for i := 1; i >= 0; i++ {
				if filterdChar.Contains(html[ind+i]) {
					continue
				}
				if html[ind+i] == byte('>') {
					tag[len(tag)-1].row = append(tag[len(tag)-1].row, html[ind+i])
					ind += i
					break
				} else if html[ind+i] == byte('<') {
					ind += i - 1
					break
				} else if html[ind+i] == byte('/') && i == 1 {
					tag[len(tag)-1].typeRow = 3
				}
				tag[len(tag)-1].row = append(tag[len(tag)-1].row, html[ind+i])
			}
		} else {
			if len(tag[len(tag)-1].row) != 0 {
				tag = append(tag, rowHtml{
					row:     []byte{},
					typeRow: 2,
				})
			} else {
				tag[len(tag)-1].typeRow = 2
			}
			var spaceCounter uint = 0
			for i := 0; i >= 0; i++ {
				if filterdChar.Contains(html[ind+i]) {
					continue
				}
				if html[ind+i] == byte('>') {
					tag[len(tag)-1].row = append(tag[len(tag)-1].row, html[ind+i])
					ind += i
					break
				} else if html[ind+i] == byte('<') {
					ind += i - 1
					break
				} else if html[ind+i] == byte(' ') {
					spaceCounter += 1
					continue
				}
				if spaceCounter == 1 {
					tag[len(tag)-1].row = append(tag[len(tag)-1].row, html[ind+i-1])
				}
				spaceCounter = 0
				tag[len(tag)-1].row = append(tag[len(tag)-1].row, html[ind+i])
			}
		}
	}
	createHtmlFile("./Pages/templates/test.html", tag)
	for _, val := range tag {
		fmt.Println(string(val.row))
	}
	return
}

type char []byte

func (list char) Contains(val byte) bool {
	for _, i := range list {
		if i == val {
			return true
		}
	}
	return false
}

func createHtmlFile(path string, htmlFile []rowHtml) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	for _, val := range htmlFile {
		// Append the first message of bytes
		_, err = file.Write(val.row)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//func extractAtrTag()

func convToFunc(bodyTag string) string {

	switch bodyTag {
	case "html":
		fmt.Println(bodyTag)
	case "head":
		fmt.Println(bodyTag)
	case "body":
		fmt.Println(bodyTag)
	case "div":
		fmt.Println(bodyTag)
	default:
		// Handle unknown tags
		fmt.Printf("Unknown tag: %s\n", bodyTag)
	}
	return ""
}
