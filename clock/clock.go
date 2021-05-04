package clock

import (
	"fmt"
	"strconv"
)

const minsInHour = 60
const minsInDay = 24 * minsInHour

type Clock struct {
	mins int
}

func New(h, m int) *Clock {
	return &Clock{
		mins: h*minsInHour + m,
	}
}

func (c *Clock) Add(n int) *Clock {
	c.mins += n

	c.normalize()

	return c
}

func (c *Clock) Subtract(n int) *Clock {
	c.mins -= n

	c.normalize()

	return c
}

func (c *Clock) String() string {
	c.normalize()

	var h, m int

	h = c.mins / minsInHour
	m = c.mins - h*minsInHour

	return fmt.Sprintf("%s:%s", format(h), format(m))
}

func format(n int) string {
	s := strconv.Itoa(n)

	if len(s) == 1 {
		s = "0" + s
	}

	return s
}

func (c *Clock) normalize () {
	c.mins = c.mins % minsInDay

	if c.mins < 0 {
		c.mins += minsInDay
	}
}
