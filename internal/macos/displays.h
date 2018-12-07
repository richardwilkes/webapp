#import <Quartz/Quartz.h>

typedef struct {
	CGRect bounds;
	CGRect usableBounds;
	double scalingFactor;
	int    isMain;
} Display;

Display *displays(unsigned long *qty);
