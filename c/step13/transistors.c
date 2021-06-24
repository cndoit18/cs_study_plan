#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>

#include "transistors.h"

/*
 * Name :  cndoit18
 * Description : Simple transistor description example program
*/

int main()
{
	int i;
	Tran *trans;
	int numTrans = 0;
	char again[2];

	printf("transistors!\n");
	/* Input the transistor */
	trans = malloc(sizeof(Tran) * (++numTrans + 1));
	trans[0] = InputTransistor();

	while (1)
	{
		do
		{
			printf("Would you like to enter another transistor (Y/N)?");
			InputString(again, 2);
			if (again[0] == 'N' || again[0] == 'n')
			{
				goto transistor;
			}
		} while (again[0] != 'Y' && again[0] != 'y');

		/* Allocate space for one transistor */
		trans = realloc(trans, sizeof(Tran) * (numTrans + 1));
		numTrans++;

		trans[numTrans - 1] = InputTransistor();
	}
transistor:
	/* Output the transistors */
	printf("\nThe transistors:\n");
	for (i = 0; i < numTrans; i++)
	{
		DisplayTransistor(trans[i]);
	}

	/* Free the memory */
	free(trans);
}

void DisplayTransistor(Tran tran)
{
	printf("Number: %s\n", tran.number);
	switch (tran.type)
	{
	case NPN:
		printf("Type: NPN\n");
		break;

	case PNP:
		printf("Type: PNP\n");
		break;
	}

	printf("pMax: %.3f\n", tran.pmax);
	printf("icMax: %.3f\n", tran.icmax);
}

Tran InputTransistor()
{
	Tran t1;

	InputString(t1.number, sizeof(t1.number));
	t1.type = InputTransistorType("Input type: ");
	t1.caseStyle = TO39;
	t1.pmax = InputPositiveValue("Input pMax: ");
	t1.icmax = InputPositiveValue("Input icMax: ");

	return t1;
}

void InputString(char *str, int max)
{
	char buffer[100];

	/* Get a line of up to 100 characters */
	fgets(buffer, sizeof(buffer), stdin);

	/* Remove any stray newlines from the buffer */
	while (buffer[0] == '\n')
		fgets(buffer, sizeof(buffer), stdin);

	/* Remove any \n we may have input */
	if (strlen(buffer) > 0)
		buffer[strlen(buffer) - 1] = '\0';

	/* Copy up to max characters to our string */
	strncpy(str, buffer, max);
	str[max - 1] = '\0';
}

double InputPositiveValue(char *prompt)
{
	char buffer[100];
	double value = 0;
	do
	{
		printf("%s", prompt);

		/* Get a line of up to 100 characters */
		fgets(buffer, sizeof(buffer), stdin);

		/* Remove any \n we may have input */
		if (strlen(buffer) > 0)
			buffer[strlen(buffer) - 1] = '\0';

		sscanf(buffer, "%lf", &value);
	} while (value <= 0);
	return value;
}

int InputTransistorType(char *prompt)
{
	char buffer[100];
	int value = 0;
	do
	{
		printf("%s", prompt);
		fgets(buffer, sizeof(buffer), stdin);
		if (strlen(buffer) > 0)
			buffer[strlen(buffer) - 1] = '\0';

		sscanf(buffer, "%d", &value);
	} while (value != NPN && value != PNP);
	return value;
}