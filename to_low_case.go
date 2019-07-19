package main

import "strings"

func toLowerCase(str string) string {
	if str == "" {
		return ""
	}

	var newstring []string
	for _, cint := range str {
		if cint >= 65 && cint <= 90 {
			c := string(cint + 32)
			newstring = append(newstring, c)
		} else {
			newstring = append(newstring, string(cint))
		}
	}

	return strings.Join(newstring, "")
}
