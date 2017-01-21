package counter

import ()

// contains the businesshours of a single day
type BusinessHoursForADay struct {
	Name                   string         // name of the day
	ID                     int            // name of the day
	AssignedBusinessHours  []BusinessHour // assigned businesshours
	RemainingBusinessHours []BusinessHour // existing businesshours without the assigned ones
}
