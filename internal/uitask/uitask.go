package uitask

import "sync"

var (
	nextID uint32 = 1
	lock   sync.Mutex
	tasks  = make(map[uint32]func())
)

// Record a task for later execution on the UI thread.
func Record(task func()) uint32 {
	lock.Lock()
	id := nextID
	nextID++
	tasks[id] = task
	lock.Unlock()
	return id
}

// Dispatch a UI task. Should only be called on the UI thread.
func Dispatch(id uint32) {
	lock.Lock()
	task := tasks[id]
	delete(tasks, id)
	lock.Unlock()
	if task != nil {
		task()
	}
}
