package windows

import (
	"syscall"

	"github.com/richardwilkes/toolbox/errs"
)

func toUTF16PtrOrNilOnEmpty(in string) (*uint16, error) {
	if in == "" {
		return nil, nil
	}
	return toUTF16Ptr(in)
}

func toUTF16Ptr(in string) (*uint16, error) {
	out, err := syscall.UTF16PtrFromString(in)
	if err != nil {
		return nil, errs.NewWithCause("Unable to convert string to UTF16", err)
	}
	return out, nil
}

func fromBOOL(in BOOL) bool {
	if in == 0 {
		return false
	}
	return true
}

func toBOOL(in bool) BOOL {
	if in {
		return 1
	}
	return 0
}
