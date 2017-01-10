package main

import (
	"github.com/ericmieth/ub-shift/view"
	"net/http"
)

func main() {

	// create da databse connection pool
	db := dbOpen()
	defer db.Close()

	// prepare template
	t := view.ParseTemplates()

	// create a http server
	http.NewServeMux()

	http.Handle(
		"/static/bootstrap/",
		http.StripPrefix(
			"/static/bootstrap/",
			http.FileServer(http.Dir("static/bootstrap/dist/")),
		),
	)

	http.Handle(
		"/static/fonts/",
		http.StripPrefix(
			"/static/fonts/",
			http.FileServer(http.Dir("static/fonts/")),
		),
	)

	http.Handle(
		"/static/bootstrap-modification/",
		http.StripPrefix(
			"/static/bootstrap-modification/",
			http.FileServer(
				http.Dir("static/bootstrap-modification/")),
		),
	)
	http.Handle(
		"/static/jquery/",
		http.StripPrefix(
			"/static/jquery/",
			http.FileServer(http.Dir("static/jquery/")),
		),
	)
	http.Handle(
		"/static/js/",
		http.StripPrefix(
			"/static/js/",
			http.FileServer(http.Dir("static/js/")),
		),
	)

	// calendar

	http.HandleFunc("/calendar/month/",
		argsHandler(view.CalendarViewMonthHandler, db, t),
	)

	// employees

	http.HandleFunc("/employee/list/",
		argsHandler(view.EmployeeListHandler, db, t),
	)

	// branches

	http.HandleFunc("/branch/list/",
		argsHandler(view.BranchListHandler, db, t),
	)

	http.HandleFunc("/branch/add/",
		argsHandler(view.BranchAddHandler, db, t),
	)

	http.HandleFunc("/branch/edit/",
		argsHandler(view.BranchEditHandler, db, t),
	)

	// counters

	http.HandleFunc("/counter/list/",
		argsHandler(view.CounterListHandler, db, t),
	)

	http.HandleFunc("/counter/add/",
		argsHandler(view.CounterAddHandler, db, t),
	)

	http.HandleFunc("/counter/edit/",
		argsHandler(view.CounterEditHandler, db, t),
	)

	http.HandleFunc("/",
		argsHandler(view.IndexHandler, db, t),
	)

	http.ListenAndServe(":9002", nil)
}
