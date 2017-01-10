package counter

import (
//"github.com/ericmieth/ub-shift/employee"
)

type Counter struct {
	ID     int    // ID
	Name   string // name of the desk
	Branch Branch // branch the desk belongs to
	//BusinessHours []BusinessHour      // contains the business hours for each day
	//Employees     []employee.Employee // which employees are allowed to work on this counter
	//DaysOff       []Day           // days, when the desk is closed irregulary
	//Shifts		[]shift.Shift   // contains all relating shifts
}
