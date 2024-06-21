package timer

import "time"

type Timer struct {
	StartTime time.Time
	EndTime   time.Time
}

func (t *Timer) Start() {
	t.StartTime = time.Now()
}

func (t *Timer) Stop() {
	t.EndTime = time.Now()
}

func (t *Timer) Reset() {
	t.StartTime = time.Now()
	t.EndTime = time.Now()
}

func (t *Timer) Duration() time.Duration {
	return t.EndTime.Sub(t.StartTime)
}
