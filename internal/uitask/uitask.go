package uitask

import "sync"

var (
	nextID uint64 = 1
	lock   sync.Mutex
	tasks  = make(map[uint64]func())
)

// Record a task for later execution on the UI thread.
func Record(task func()) uint64 {
	lock.Lock()
	id := nextID
	nextID++
	tasks[id] = task
	lock.Unlock()
	return id
}

// Dispatch a UI task. Should only be called on the UI thread.
func Dispatch(id uint64) {
	lock.Lock()
	task := tasks[id]
	delete(tasks, id)
	lock.Unlock()
	if task != nil {
		task()
	}
}
