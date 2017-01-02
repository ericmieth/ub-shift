package funcs

import (
	"html/template"
)

// returns bootstrap glyphicons for booleans

func Checked(b bool) template.HTML {
	if b {
		return "<span class=\"glyphicon glyphicon-check\"></span>"
	} else {
		return "<span class=\"glyphicon glyphicon-unchecked\"></span>"
	}
}
