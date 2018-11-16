#import <Quartz/Quartz.h>

typedef struct {
	CGRect bounds;
	CGRect usableBounds;
	int    isMain;
} Display;

Display *displays(unsigned long *qty);
