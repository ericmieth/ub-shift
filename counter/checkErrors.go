package counter

import ()

// returns true 'err' if contains at least one error
func checkErrors(err []error) bool {
	for _, e := range err {
		if e != nil {
			return true
		}
	}
	return false
}
