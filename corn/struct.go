package corn

import (
	"time"
)

type Event struct {
	Func func()
	Time time.Time
}

type Corn struct {
	events data
}
