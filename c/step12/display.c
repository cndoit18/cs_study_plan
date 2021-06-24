#include <stdio.h>

#include "hanoi.h"

/*
 * Display a disk of width p
 */
void DisplayDisk(int p)
{
	int i;
	int spaceAround = (MaxDiskSize - p) / 2;

	for (i = 0; i < spaceAround; i++)
		printf(" ");

	if (p == 0)
	{
		printf("|");
	}
	else
	{
		for (i = 0; i < p; i++)
			printf("O");
	}

	for (i = 0; i < spaceAround; i++)
		printf(" ");
}