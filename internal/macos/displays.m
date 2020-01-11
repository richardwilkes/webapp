// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#import "displays.h"

Display *displays(unsigned long *qty) {
	NSArray *s = [NSScreen screens];
	unsigned long count = [s count];
	*qty = count;
	Display *d = (Display *)malloc(sizeof(Display) * count);
	for (unsigned int i = 0; i < count; i++) {
		NSScreen *screen = [s objectAtIndex: i];
		CGDirectDisplayID dID = (CGDirectDisplayID)[[[screen deviceDescription] objectForKey:@"NSScreenNumber"] unsignedIntValue];
		d[i].bounds = CGDisplayBounds(dID);
		d[i].usableBounds = [screen visibleFrame];
		d[i].scalingFactor = [screen backingScaleFactor];
		d[i].isMain = CGDisplayIsMain(dID);
		CGRect b = [screen frame];
		d[i].usableBounds.origin.y = d[i].bounds.origin.y + (b.origin.y + b.size.height - (d[i].usableBounds.origin.y + d[i].usableBounds.size.height));
	}
	return d;
}
