package webapp

import (
	"time"

	"github.com/richardwilkes/cef/cef"
)

type taskFunc struct {
	task func()
}

func (t *taskFunc) Execute(self *cef.Task) {
	t.task()
}

// InvokeUITask a task on the UI thread. The task is put into the system event
// queue and will be run at the next opportunity.
func InvokeUITask(task func()) {
	cef.TaskRunnerGetForThread(cef.TIDUI).PostTask(cef.NewTask(&taskFunc{task: task}))
}

// InvokeUITaskAfter schedules a task to be run on the UI thread after waiting
// for the specified duration.
func InvokeUITaskAfter(task func(), after time.Duration) {
	cef.TaskRunnerGetForThread(cef.TIDUI).PostDelayedTask(cef.NewTask(&taskFunc{task: task}), int64(after/time.Millisecond))
}
