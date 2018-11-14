package webapp

import (
	"time"

	"github.com/richardwilkes/webapp/internal/cef"
)

// InvokeUITask a task on the UI thread. The task is put into the system event
// queue and will be run at the next opportunity.
func InvokeUITask(task func()) {
	cef.PostUITask(task)
}

// InvokeUITaskAfter schedules a task to be run on the UI thread after waiting
// for the specified duration.
func InvokeUITaskAfter(task func(), after time.Duration) {
	time.AfterFunc(after, func() { InvokeUITask(task) })
}
