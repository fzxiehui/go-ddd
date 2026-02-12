package job

import "github.com/robfig/cron/v3"

type Job struct {
	ID       string
	Name     string
	Spec     string // cron 表达式
	Enabled  bool
	EntryID  cron.EntryID
	Callback func()
}
