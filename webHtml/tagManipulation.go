package webJs

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
	for _, v := range value {
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
	}
	return fmt.Sprintf("<%s%s>%s\n</%s>", tag, atr, value, tag)
}
