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

	http.HandleFunc("/view/month/", argsHandler(view.ViewMonthHandler, db, t))
	http.HandleFunc("/", argsHandler(view.IndexHandler, db, t))

	http.ListenAndServe(":9002", nil)
}
