package clock

// Clock type
type Clock struct {
	hours   int
	minutes int
}

// New : Creates a new clock with an initial time
func New(hours, minutes int) Clock {
	return Clock{hours, minutes}
}

// String : Returns Clock in String form
func String(Clock) string {

}

// Add Minutes to time on Clock c
func (c Clock) Add(minutes int) Clock {

}

// Subtract Minutes from tme on Clock c
func (c Clock) Subtract(minutes int) Clock {

}

// Compare Clocks
