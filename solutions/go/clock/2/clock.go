package clock

import "fmt"

// Define the Clock type here.

type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	h += int(m / 60)
	h %= 24
	for m < 0 {
		h -= 1
		m += 60
	}
	for h < 0 {
		h += 24
	}
	return Clock{hours: h, minutes: m}
}

func (c Clock) Add(m int) Clock {
	minute := c.minutes + m
	hour := c.hours
	for minute >= 60 {
		hour += 1
		minute -= 60
		if hour >= 24 {
			hour -= 24
		}
	}

	return Clock{hours: hour, minutes: minute}
}

func (c Clock) Subtract(m int) Clock {
	minute := c.minutes - m
	hour := c.hours
	for minute < 0 {
		hour -= 1
		minute += 60
		if hour < 0 {
			hour += 24
		}
	}

	return Clock{hours: hour, minutes: minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
