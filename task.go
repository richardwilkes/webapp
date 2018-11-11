package webapp

import (
	"time"

	"github.com/richardwilkes/webapp/internal/uitask"
)

// InvokeUITask a task on the UI thread. The task is put into the system event
// queue and will be run at the next opportunity.
func InvokeUITask(task func()) {
	driver.Invoke(uitask.Record(task))
}

// InvokeUITaskAfter schedules a task to be run on the UI thread after waiting
// for the specified duration.
func InvokeUITaskAfter(f func(), after time.Duration) {
	driver.InvokeAfter(uitask.Record(f), after)
}
