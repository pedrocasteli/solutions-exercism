/* This package is about leap years. Currently, it contains a function to determine if a given year is, in fact, a leap year */
package leap

// The IsLeapYear function receives a year (int) as parameter ands returns either true or false (bool) 
func IsLeapYear(year int) bool {
	if year % 4 == 0 {
        if year % 100 == 0 && year % 400 != 0 {
            return false
        }
        return true
    } else {
        return false
    }
	panic("Please implement the IsLeapYear function")
}
