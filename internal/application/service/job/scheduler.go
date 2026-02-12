package job

import (
	"errors"
	"sync"

	domain "ddd/internal/domain/job"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
	jobs map[string]*domain.Job
	mu   sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron: cron.New(cron.WithSeconds()),
		jobs: make(map[string]*domain.Job),
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	ctx := s.cron.Stop()
	<-ctx.Done()
}

// 添加任务
func (s *Scheduler) AddJob(name string, spec string, callback func()) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.NewString()

	job := &domain.Job{
		ID:       id,
		Name:     name,
		Spec:     spec,
		Enabled:  true,
		Callback: callback,
	}

	entryID, err := s.cron.AddFunc(spec, callback)
	if err != nil {
		return "", err
	}

	job.EntryID = entryID
	s.jobs[id] = job

	return id, nil
}

// 删除任务
func (s *Scheduler) RemoveJob(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	job, ok := s.jobs[id]
	if !ok {
		return errors.New("job not found")
	}

	s.cron.Remove(job.EntryID)
	delete(s.jobs, id)

	return nil
}

// 列出任务
func (s *Scheduler) ListJobs() []*domain.Job {
	s.mu.Lock()
	defer s.mu.Unlock()

	var list []*domain.Job
	for _, j := range s.jobs {
		list = append(list, j)
	}
	return list
}
