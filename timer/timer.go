package timer

import "time"

type Timer struct {
	startTime time.Time
	endTime   time.Time
}

func (t *Timer) Start() {
	t.startTime = time.Now()
}

func (t *Timer) Stop() {
	t.endTime = time.Now()
}

func (t *Timer) Reset() {
	t.startTime = time.Now()
	t.endTime = time.Now()
}

func (t *Timer) Duration() time.Duration {
	return t.endTime.Sub(t.startTime)
}
