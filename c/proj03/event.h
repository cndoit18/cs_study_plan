#include <time.h>

#ifndef EVENT_H
#define EVENT_H

struct event
{
	char *purpose;
	time_t end_date;
	time_t start_date;
};

time_t InputDate(char *prompt);
time_t InputTime(char *prompt, time_t date);
void InputString(char *str, int max);
int InputInt(char *prompt);
struct event InputEvent();
#endif