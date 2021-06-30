#include <stdio.h>
#include <stdlib.h>
#include "scheduler.h"
#include "display.h"
#include "event.h"

int main()
{
	struct scheduler *schedulers = (struct scheduler *)malloc(sizeof(struct scheduler));
	struct scheduler *current = schedulers;
	struct scheduler *temp;
	time_t now;
	int option;
	while (1)
	{
		do
		{
			printf("1 - Insert a new event\n");
			printf("2 - Display all events\n");
			printf("3 - Now?\n");
			printf("4 - Delete expired\n");
			printf("0 - Exit\n");
			option = InputInt("Please select an option: ");
			if (feof(stdin))
				return 0;
		} while (option < 0 || option > 4);

		switch (option)
		{
		case 1:
			temp = (struct scheduler *)malloc(sizeof(struct scheduler));
			temp->next = NULL;
			temp->prev = current;
			temp->event = InputEvent();
			current->next = temp;
			current = temp;
			break;
		case 2:
			printf("Schedule: \n");
			for (struct scheduler *c = schedulers->next; c != NULL; c = c->next)
			{
				DisplayEvent(c->event);
			}
			break;
		case 3:
			printf("Currently active events: \n");
			now = time(NULL);
			for (struct scheduler *c = schedulers->next;
				 c != NULL && now >= c->event.start_date && now <= c->event.end_date;
				 c = c->next)
			{
				DisplayEvent(c->event);
			}
			break;
		case 4:
			printf("Deleting: \n");
			now = time(NULL);
			for (struct scheduler *c = schedulers->next;
				 c != NULL && now > c->event.end_date;)
			{
				DisplayEvent(c->event);
				if (c->prev != NULL)
				{
					c->prev->next = c->next;
				}
				if (c->next != NULL)
				{
					c->next->prev = c->prev;
				}
				free(c->event.purpose);
				free(c);
				c = c->next;
			}
			break;
		case 0:
			exit(0);
		}
	}

	return 0;
}
