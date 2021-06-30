#include "event.h"

#ifndef SCHEDULER_H
#define SCHEDULER_H

struct scheduler
{
	struct scheduler *next;
	struct scheduler *prev;
	struct event event;
};

#endif