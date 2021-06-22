#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

/*
 * Name :  cndoit18
 * 
 * Simple Wumpus game in 2D
*/

/* Id's for things in our Cave */
#define Empty 0
#define Wumpus 1
#define End 2
#define Pit 3

/* Number of rooms in our Cave in each dimension */
#define CaveSize 10
#define ArraySize (CaveSize + 2)

/* Directions I can face */
#define Left 0
#define Up 1
#define Right 2
#define Down 3

void CreateWorld(int cave[ArraySize][ArraySize]);
int *GetEmptyRoot(int cave[ArraySize][ArraySize]);
void DisplayWorld(int cave[ArraySize][ArraySize], int *agent, int agentDir);
int DifferenceByDirection(int dir);
bool DisplayStatus(int cave[ArraySize][ArraySize], int *agent);

int main()
{

	int cave[ArraySize][ArraySize];
	int *agentRoom;
	int agentDirection;
	char command[20];

	/* Seed the random number generator */
	srand(time(NULL));

	CreateWorld(cave);

	agentRoom = GetEmptyRoot(cave);
	agentDirection = rand() % 4;

	/* The game loop */
	while (true)
	{

		if (DisplayStatus(cave, agentRoom))
		{
			break;
		}

		DisplayWorld(cave, agentRoom, agentDirection);
		/* Get the command */
		printf("Command: ");
		scanf("%20s", command);
		if (strcmp(command, "quit") == 0)
		{
			/* Exit, we are doing */
			break;
		}
		else if (strcmp(command, "move") == 0)
		{
			int d = DifferenceByDirection(agentDirection);

			if (*(agentRoom + d) != End)
				agentRoom += d;
		}
		else if (strcmp(command, "turn") == 0)
		{
			agentDirection++;
			if (agentDirection > Down)
				agentDirection = Left;
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

void CreateWorld(int cave[ArraySize][ArraySize])
{
	int *room;
	for (int col = 0; col < ArraySize; col++)
	{
		for (int row = 0; row < ArraySize; row++)
		{
			if (row == 0 || row == ArraySize - 1 || col == 0 || col == ArraySize - 1)
			{
				cave[col][row] = End;
				continue;
			}
			cave[col][row] = Empty;
		}
	}

	for (int pits = 0; pits < 10; pits++)
	{
		*GetEmptyRoot(cave) = Pit;
	}

	room = GetEmptyRoot(cave);
	*room = Wumpus;
}

int *GetEmptyRoot(int cave[][ArraySize])
{
	int col, row;
	do
	{
		col = rand() % ArraySize, row = rand() % ArraySize;
	} while (cave[col][row] != Empty);

	return &cave[col][row];
}

void DisplayWorld(int cave[ArraySize][ArraySize], int *agent, int agentDir)
{
	int row, col;
	int *room;

	/* Loop over the rows of the cave */
	for (row = 1; row <= CaveSize + 1; row++)
	{
		/* 
         * This loop lets us print an up direction 
         * above the agent or a v below the agent
         */

		for (col = 1; col <= CaveSize; col++)
		{
			if (&cave[row][col] == agent && agentDir == Up)
			{
				printf(" ^  ");
			}
			else if (&cave[row - 1][col] == agent && agentDir == Down)
			{
				printf(" v  ");
			}
			else
			{
				printf("    ");
			}
		}

		printf("\n");
		if (row > CaveSize)
			break;

		/*
         * This loop prints the agent or the room contents
         */

		for (col = 1; col <= CaveSize; col++)
		{
			room = &cave[row][col];
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

				default:
					printf(" A  ");
					break;
				}
				continue;
			}

			switch (*room)
			{
			case Wumpus:
				printf("-W- ");
				break;
			case Pit:
				printf(" X  ");
				break;
			default:
				printf(" .  ");
				break;
			}
		}

		printf("\n");
	}
}

int DifferenceByDirection(int dir)
{
	switch (dir)
	{
	case Up:
		return -ArraySize;

	case Down:
		return ArraySize;

	case Left:
		return -1;

	case Right:
		return 1;
	}

	return 0;
}

bool DisplayStatus(int cave[ArraySize][ArraySize], int *agent)
{
	if (*agent == Wumpus)
	{
		printf("You have been eaten by the Wumpus\n");
		return true;
	}

	if (*agent == Pit)
	{
		printf("You fell into a pit\n");
		return true;
	}

	if (*(agent - 1) == Wumpus || *(agent + 1) == Wumpus || *(agent + ArraySize) == Wumpus || *(agent - ArraySize) == Wumpus)
	{
		printf("I smell a Wumpus\n");
		return false;
	}

	if (*(agent - 1) == Pit || *(agent + 1) == Pit || *(agent + ArraySize) == Pit || *(agent - ArraySize) == Pit)
	{
		printf("I feel a draft\n");
		return false;
	}

	for (int col = 0; col < ArraySize; col++)
	{
		for (int row = 0; row < ArraySize; row++)
		{
			if (cave[col][row] == Wumpus)
			{
				return false;
			}
		}
	}

	printf("You are win!\n");
	return true;
}