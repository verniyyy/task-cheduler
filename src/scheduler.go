package src

import (
	mod_scheduler "github.com/carlescere/scheduler"
)

type Scheduler interface {
	RunEveryDayAt(t Task, hourTime string)
	RunEverySeconds(t Task, seconds int)
}

func NewScheduler() Scheduler {
	return &scheduler{}
}

type scheduler struct{}

func (s scheduler) RunEveryDayAt(t Task, hourTime string) {
	mod_scheduler.Every().Day().At(hourTime).Run(t.Job)
}

func (s scheduler) RunEverySeconds(t Task, seconds int) {
	mod_scheduler.Every(seconds).Seconds().Run(t.Job)
}
