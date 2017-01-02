package view

import (
	"github.com/ericmieth/ub-shift/employee"

	"html/template"
	"log"
	"net/http"
	"time"

	"database/sql"
)

func EmployeeListHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()

	employees, err := employee.ListEmployees(db)
	if err != nil {
		log.Println(err)
	}

	data := struct {
		PageTitle    interface{}
		TemplateName string
		ContentData  interface{}
		StartTime    time.Time
	}{
		PageTitle:    "Mitarbeiter",
		TemplateName: "/employee/list/",
		ContentData:  employees,
		StartTime:    startTime,
	}

	renderView(w, t, data)
}
