package webapp

import (
	"sync"
	"time"
)

var (
	uiTaskNextID uint64 = 1
	uiTaskLock   sync.Mutex
	uiTaskMap    = make(map[uint64]func())
)

// InvokeUITask a task on the UI thread. The task is put into the system event
// queue and will be run at the next opportunity.
func InvokeUITask(task func()) {
	platformInvoke(recordUITask(task))
}

// InvokeUITaskAfter schedules a task to be run on the UI thread after waiting
// for the specified duration.
func InvokeUITaskAfter(f func(), after time.Duration) {
	platformInvokeAfter(recordUITask(f), after)
}

func recordUITask(task func()) uint64 {
	uiTaskLock.Lock()
	id := uiTaskNextID
	uiTaskNextID++
	uiTaskMap[id] = task
	uiTaskLock.Unlock()
	return id
}

func dispatchUITask(id uint64) {
	uiTaskLock.Lock()
	task := uiTaskMap[id]
	delete(uiTaskMap, id)
	uiTaskLock.Unlock()
	if task != nil {
		task()
	}
}
