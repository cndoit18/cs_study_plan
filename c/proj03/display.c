#include <stdio.h>
#include <time.h>
#include "display.h"

void DisplayEvent(struct event event)
{
	char start_buffer[100];
	char end_buffer[100];
	struct tm start_date;
	struct tm end_date;
	start_date = *localtime(&event.start_date);
	end_date = *localtime(&event.end_date);
	strftime(start_buffer, sizeof(start_buffer), "%m/%d/%Y  %I:%M%p", &start_date);
	strftime(end_buffer, sizeof(end_buffer), "%I:%M%p", &end_date);
	printf("%s  %s  %s\n", start_buffer, end_buffer, event.purpose);
}