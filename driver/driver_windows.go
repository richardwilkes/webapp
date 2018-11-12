package driver

import (
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/internal/windows"
)

// ForPlatform returns the driver for your platform.
func ForPlatform() webapp.Driver {
	return windows.Driver()
}
