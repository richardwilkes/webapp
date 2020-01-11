// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package macos

import "github.com/richardwilkes/cef/cef"

func (d *driver) OnPreKeyEvent(event *cef.KeyEvent, isKeyboardShortcut *int32) int32 {
	return 0
}

func (d *driver) OnKeyEvent(event *cef.KeyEvent) int32 {
	return 0
}
