package view

import (
	"github.com/ericmieth/ub-shift/templates/funcs"

	"html/template"
)

func ParseTemplates() *template.Template {
	funcMap := template.FuncMap{
		"dateDay":           funcs.DateDay,
		"dateMonth":         funcs.DateMonth,
		"dateWeekdayID":     funcs.DateWeekdayID,
		"dateFillMonth":     funcs.DateFillMonth,
		"dateNextMonth":     funcs.DateNextMonth,
		"datePreviousMonth": funcs.DatePreviousMonth,
		"dateCurrentMonth":  funcs.DateCurrentMonth,
	}
	return template.Must(
		template.New("content").
			Funcs(funcMap).
			ParseGlob("templates/layouts/*.tmpl"),
	)

}
