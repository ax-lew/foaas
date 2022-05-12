package clock

import "time"

type DefaultClock struct {
}

func (d *DefaultClock) Now() time.Time {
	return time.Now()
}
