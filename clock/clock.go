package clock

import "time"

//go:generate mockery --name Clock --filename clock.go --quiet

type Clock interface {
	Now() time.Time
}
