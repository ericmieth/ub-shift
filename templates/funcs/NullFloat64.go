package funcs

import (
	"database/sql"
	"html/template"
	"strconv"
)

// returns the value of an sql.NullFloat64 if its valid and an empty
// string otherwise

func NullFloat64(b sql.NullFloat64) template.HTML {
	if b.Valid {
		return template.HTML(strconv.FormatFloat(b.Float64, 'f', -1, 64))
	} else {
		return ""
	}
}
