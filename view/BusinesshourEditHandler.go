package view

import (
	"github.com/ericmieth/ub-shift/counter"

	"log"

	"errors"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"database/sql"
)

func BusinesshourEditHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()

	var e error
	var err []error

	counterID, e := strconv.Atoi(path.Base(r.URL.Path))
	if e != nil {
		http.Redirect(w, r, "/counter/list/", http.StatusFound)
	}

	// POST request
	if r.Method == "POST" {

		inputAction := strings.TrimSpace(r.FormValue("action"))
		inputWeekdayID, e := strconv.Atoi(
			strings.TrimSpace(r.FormValue("weekday")),
		)
		if e != nil {
			log.Println(e)
			err = append(err, e)
		}
		r.ParseForm()

		switch inputAction {
		case "assign":
			inputBusinessHourID, e := strconv.Atoi(
				strings.TrimSpace(r.FormValue("businesshourID")),
			)
			if e != nil {
				err = append(err, e)
			}
			err2 := counter.AssignBusinessHour(
				db,
				counterID,
				inputWeekdayID,
				inputBusinessHourID,
			)
			if checkErrors(err2) {
				err = append(err, err2...)
			}

		case "create-assign":
			inputOpening, e := time.Parse(
				"15:04",
				strings.TrimSpace(r.FormValue("opening")),
			)
			if e != nil {
				err = append(
					err,
					errors.New("ungültiges Zeitformat im Feld Beginn."),
				)
				// break if format is invalid
				break
			}

			inputClosing, e := time.Parse(
				"15:04",
				strings.TrimSpace(r.FormValue("closing")),
			)
			if e != nil {
				err = append(
					err,
					errors.New("ungültiges Zeitformat im Feld Ende."),
				)
				// break if format is invalid
				break
			}

			if inputClosing.Sub(inputOpening) <= 0 {
				err = append(
					err,
					errors.New("Das Ende muss nach dem Beginn liegen."),
				)
				break
			}

			businessHourID, err2 := counter.CreateBusinessHour(
				db,
				inputOpening,
				inputClosing,
			)
			if checkErrors(err2) {
				err = append(err, err2...)
			}

			err2 = counter.AssignBusinessHour(
				db,
				counterID,
				inputWeekdayID,
				int(businessHourID),
			)
			if checkErrors(err2) {
				err = append(err, err2...)
			}

		case "remove":
			inputBusinessHourID, e := strconv.Atoi(
				strings.TrimSpace(r.FormValue("businesshourID")),
			)
			if e != nil {
				err = append(err, e)
			}
			err2 := counter.RemoveBusinessHour(
				db,
				counterID,
				inputWeekdayID,
				inputBusinessHourID,
			)
			if checkErrors(err2) {
				err = append(err, err2...)
			}

		}

		//err = counter.EditBranch(db, branchID, inputName, inputLocation)
		//if !checkErrors(err) {
		//	http.Redirect(w, r, "/branch/list/", http.StatusFound)
		//}

	}

	// GET request

	//c, e := counter.ReturnCounter(db, counterID)
	//if e != nil {
	//	err = append(err, e)
	//}

	// information for the counter
	c, e := counter.ReturnCounter(db, counterID)
	if e != nil {
		http.Redirect(w, r, "/counter/list/", http.StatusFound)
	}

	// information for each day
	d, err2 := counter.ReturnBusinessHoursForAllDays(db, counterID)
	if e != nil {
		err = append(err, err2...)
	}

	content := struct {
		Counter counter.Counter
		Days    []counter.BusinessHoursForADay
	}{
		Counter: c,
		Days:    d,
	}

	data := struct {
		PageTitle    interface{}
		TemplateName string
		ContentData  interface{}
		StartTime    time.Time
		Error        []error
	}{
		PageTitle:    "Geschäftszeiten bearbeiten",
		TemplateName: "/businesshour/edit/",
		ContentData:  content,
		StartTime:    startTime,
		Error:        err,
	}

	renderView(w, t, data)
}
