package counter

import (
//"github.com/ericmieth/ub-shift/employee"
)

type Counter struct {
	ID            int            // ID
	Name          string         // name of the desk
	Branch        Branch         // branch the desk is belonging to
	BusinessHours []BusinessHour // contains the business hours for each day
	//Shifts        []Shift        // list of all shifts belonging to a counter
	//Employees     []employee.Employee // which employees are allowed to work on this counter
	//DaysOff       []Day           // days, when the desk is closed irregulary
}
