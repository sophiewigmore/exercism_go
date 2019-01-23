package leap

// IsLeapYear returns if a year is a leap year or not (boolean)
/**
* Leap Year Criteria:
*
* on every year that is evenly divisible by 4
* 	except every year that is evenly divisible by 100
*		unless the year is also evenly divisible by 400
 */
func IsLeapYear(year int) bool {
	if (year % 4) == 0 {
		if (year % 100) == 0 {
			if (year % 400) == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
