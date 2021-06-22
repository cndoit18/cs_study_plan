#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

/*
 * Name :  cndoit18
 * 
 * Simple Wumpus game in 1D
*/

/* Add any #defines here */
#define Empty 0
#define Wumpus 1
#define End 2
#define CaveSize 20
#define ArraySize (CaveSize + 2)

#define Left 0
#define Right 1

/* Add any function prototypes here */
void CreateWorld(int cave[]);
int *GetEmptyRoot(int cave[]);
void DisplayWorld(int cave[], int *agent, int agentDir);
int DifferenceByDirection(int dir);
bool DisplayStatus(int cave[], int *agent);

int main()
{
	int cave[ArraySize];
	int *agentRoom;
	int agentDirection;

	/* Seed the random number generator */
	srand(time(NULL));

	CreateWorld(cave);

	agentRoom = GetEmptyRoot(cave);
	agentDirection = rand() % 2;

	char command[20];
	while (true)
	{
		if (DisplayStatus(cave, agentRoom))
		{
			break;
		}

		// DisplayWorld(cave, agentRoom, agentDirection);
		printf("Command: ");
		scanf("%20s", command);
		if (strcmp(command, "quit") == 0)
		{
			break;
		}
		else if (strcmp(command, "move") == 0)
		{
			int direction = DifferenceByDirection(agentDirection);
			if (*(agentRoom + direction) != End)
				agentRoom += direction;
		}
		else if (strcmp(command, "turn") == 0)
		{
			agentDirection = !agentDirection;
		}
		else if (strcmp(command, "fire") == 0)
		{
			int direction = DifferenceByDirection(agentDirection);
			for (int offset = 1; *(agentRoom + offset * direction) != End && offset <= 3; offset++)
			{
				if (*(agentRoom + offset * direction) == Wumpus)
				{
					*(agentRoom + offset * direction) = Empty;
				}
			}
		}
		else
		{
			printf("I don't know what you are talking about\n");
		}
	}
}

int DifferenceByDirection(int dir)
{
	if (dir == Left)
	{
		return -1;
	}
	return 1;
}

bool DisplayStatus(int cave[], int *agent)
{
	if (*agent == Wumpus)
	{
		printf("You have been eaten by the Wumpus\n");
		return true;
	}
	if (*(agent - 1) == Wumpus || *(agent + 1) == Wumpus)
	{
		printf("I smell a Wumpus\n");
		return false;
	}

	for (int index = 0; index < ArraySize; index++)
	{
		if (cave[index] == Wumpus)
		{
			return false;
		}
	}

	printf("You are win!\n");
	return true;
}

void CreateWorld(int cave[])
{
	for (int i = 0; i < ArraySize; i++)
	{
		cave[i] = Empty;
	}

	cave[0] = End;
	cave[ArraySize - 1] = End;

	int *room;
	room = GetEmptyRoot(cave);
	*room = Wumpus;
}

int *GetEmptyRoot(int cave[])
{
	int room;
	do
	{
		room = rand() % ArraySize;
	} while (cave[room] != Empty);

	return &cave[room];
}

void DisplayWorld(int cave[], int *agent, int agentDir)
{
	int *room;
	for (room = cave + 1; *room != End; room++)
	{
		if (room == agent)
		{
			switch (agentDir)
			{
			case Left:
				printf("<A  ");
				break;

			case Right:
				printf(" A> ");
				break;
			}
			continue;
		}
		switch (*room)
		{
		case Wumpus:
			printf("-W- ");
			break;
		default:
			printf(" .  ");
			break;
		}
	}
	printf("\n");
}