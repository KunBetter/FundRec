package schedule

import (
	"fmt"
	"github.com/robfig/cron"
)

type Scheduler struct {
	cron *cron.Cron
	id   cron.EntryID
}

func NewScheduler(specTime string, job cron.Job) (s *Scheduler) {
	s = &Scheduler{cron: cron.New()}

	var err error
	s.id, err = s.cron.AddJob(specTime, job)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return s
}

func (s *Scheduler) Start() {
	s.cron.Start()
}
