package webHtml

import (
	"fmt"
	"strings"
)

func contains(slice []string, atribut, value string) (bool, bool, int) {
	var exist = false
	var index = -1
	for ind, val := range slice {
		if val == atribut+value {
			return true, true, ind
		}
		if strings.Contains(val, atribut) {
			exist = true
		}
	}
	return exist, false, index
}

/*
Atr is function that formatted slice with given attributes into formatted string used in html tag.
*/
func Atr(atr ...string) (out string) {
	for _, val := range atr {
		out += fmt.Sprintf(" %s", val)
	}
	return
}

/*
Val is function that takes slice of string and convert it single string used in HTML.
*/
func Val(value ...string) (out string) {
	for _, val := range value {
		out += fmt.Sprintf("\n%s", val)
	}
	return
}

/*
CreateTag is function complete tag with given arguments.
*/
func CreateTag(tag, atr, value string) string {
	if len(atr) != 0 {
		if atr[0] != ' ' {
			return fmt.Sprintf("<%s %s>%s\n</%s>", tag, atr, value, tag)
		}
		return fmt.Sprintf("<%s>%s\n</%s>", tag, value, tag)
	}
	return fmt.Sprintf("<%s%s>%s\n</%s>", tag, atr, value, tag)
}

type (
	El struct {
		Html bool
		Tag  string
		Atr  map[string]any
		Val  ElVal
	}
	ElSlice struct {
		Tags []El
	}
	ElStr struct {
		Text string
	}
	ElVal interface {
		isVal()
	}
)

func (ElSlice) isVal() {}
func (ElStr) isVal()   {}

func NewEl(html bool, tag string, atr map[string]any, val ElVal) *El {
	return &El{Html: html, Tag: tag, Atr: atr, Val: val}
}

func (e *El) SetHtml(html bool) {
	e.Html = html
}

func (e *El) SetTag(tag string) {
	e.Tag = tag
}

func (e *El) SetAtr(atr map[string]any) {
	e.Atr = atr
}

func (e *El) SetVal(val ElVal) {
	e.Val = val
}

func (e *El) GenCode() string {
	var tagEl, attributes, val = "", "", ""
	switch content := e.Val.(type) {
	case ElStr:
		val += content.Text
	case ElSlice:
		for _, tag := range content.Tags {
			val += tag.GenCode()
		}
	}
	if e.Tag != "" && e.Html {
		tagEl = e.Tag
		for key, value := range e.Atr {
			attributes += fmt.Sprintf(` %s="%s"`, key, value)
		}
		return CreateTag(tagEl, attributes, val)
	}
	return val
}
