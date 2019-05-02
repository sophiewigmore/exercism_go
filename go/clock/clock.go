package clock

import "fmt"

// Clock type
type Clock struct {
	hours   int
	minutes int
}

// New : Creates a new clock with an initial time
func New(hours, minutes int) Clock {
	for minutes < 0 {
		minutes = minutes + 60
		hours--
	}
	for hours < 0 {
		hours = 24 + hours
	}
	// Add roll over minutes to hour
	hours += minutes / 60
	// Convert minutes to Clock minutes
	minutes = minutes % 60
	// Convert hours to Clock hours
	hours = hours % 24

	return Clock{hours, minutes}
}

// String : Returns Clock in String form
func (c Clock) String() string {
	stringClock := fmt.Sprintf("%2.2d:%2.2d", c.hours, c.minutes)
	return stringClock
}

// Add Minutes to time on Clock c
func (c Clock) Add(minutes int) Clock {
	newMins := c.minutes + minutes
	newClock := New(c.hours, newMins)
	return newClock
}

// Subtract Minutes from tme on Clock c
func (c Clock) Subtract(minutes int) Clock {
	newMins := c.minutes - minutes
	newClock := New(c.hours, newMins)
	return newClock
}
