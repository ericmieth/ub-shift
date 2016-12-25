package main

import (
	"database/sql"
	"html/template"
	"net/http"
)

// defines an function type which includes a database connection

type HandlerFuncArgs func(
	http.ResponseWriter,
	*http.Request,
	*sql.DB,
	*template.Template,
)

// a wrapper for http.HandlerFunc which includes some extra arguments

func argsHandler(
	f HandlerFuncArgs,
	db *sql.DB,
	t *template.Template,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, db, t)

	}
}
