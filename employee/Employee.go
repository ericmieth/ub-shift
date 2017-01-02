package employee

import (
	"database/sql"
)

// contains all information for a account
type Employee struct {
	ID           int             // ID
	FirstName    string          // first name of the employee
	LastName     string          // last name of the employee
	Manager      bool            // true, if this account has special authorization
	MailAddress  string          // mail address
	Active       bool            // true, if this account was logged in at least once
	WorkingHours sql.NullFloat64 // qouta of working hours
	//Desks        []desk.Desk // list desks this employee is allowed to work
}
