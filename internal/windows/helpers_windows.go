package windows

import (
	"syscall"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
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

func mustToUTF16Ptr(in string) *uint16 {
	out, err := toUTF16Ptr(in)
	if err != nil {
		jot.Error(err)
		var empty [1]uint16
		out = &empty[0]
	}
	return out
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
