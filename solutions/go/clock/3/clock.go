package clock

import "fmt"

// Define the Clock type here.

type Clock struct {
	hours   int
	minutes int
}

func Normalize(h, m int) (int, int) {
	h += int(m / 60)
	h %= 24
	m %= 60
	if m == 60 {
		m = 0
	}
	if m < 0 {
		h -= 1
		m += 60
	}
	if h < 0 {
		h += 24
	}
	return h, m
}
func New(h, m int) Clock {
	rh, rm := Normalize(h, m)
	return Clock{hours: rh, minutes: rm}
}

func (c Clock) Add(m int) Clock {
	minute := c.minutes + m
	rh, rm := Normalize(c.hours, minute)

	return Clock{hours: rh, minutes: rm}
}

func (c Clock) Subtract(m int) Clock {
	minute := c.minutes - m
	rh, rm := Normalize(c.hours, minute)

	return Clock{hours: rh, minutes: rm}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
