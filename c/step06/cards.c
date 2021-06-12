#include <stdio.h>
#include <stdlib.h>
#include <time.h>

/* Function declaration */
void PrintCard(int card, int suit);
void PrintResult(int player1card, int player1suit, int player2card, int player2suit);

/* 
 * Name : cndoit18
 * Program to draw playing cards
 */

int main()
{
	int player1suit, player1card;
	int player2suit, player2card;
	/* 
     . This seeds the random number 
     . generator with the current time 
     */
	srand(time(NULL));
	player1suit = rand() % 4;
	player1card = rand() % 13 + 1;

	do
	{
		player2suit = rand() % 4;
		player2card = rand() % 13 + 1;
	} while (player1suit != player2suit && player1card != player2card);
	printf("Player1: ");
	PrintCard(player1card, player1suit);
	printf("\n");

	printf("Player2: ");
	PrintCard(player2card, player2suit);
	printf("\n");

	PrintResult(player1card, player1suit, player2card, player2suit);
	return 0;
}

void PrintCard(int card, int suit)
{
	switch (card)
	{
	case 10:
		printf("Ace");
		break;
	case 11:
		printf("Jack");
		break;
	case 12:
		printf("Queen");
		break;
	case 13:
		printf("King");
		break;
	default:
		printf("%d", card);
		break;
	}
	printf(" of ");

	switch (suit)
	{
	case 0:
		printf("Hearts");
		break;

	case 1:
		printf("Diamonds");
		break;

	case 2:
		printf("Spades");
		break;

	case 3:
		printf("Clubs");
		break;
	}
}

void PrintResult(int player1card, int player1suit, int player2card, int player2suit)
{
	int wins = 0;
	if (player1card != player2card)
	{
		if (player1card == 2)
		{
			player1card -= 2;
		}
		if (player2card == 2)
		{
			player2card -= 2;
		}
		if (player1card == 10)
		{
			player1card += 4;
		}
		if (player2card == 10)
		{
			player2card += 4;
		}
		if (player1card > player2card)
		{
			wins = 1;
		}
		else if (player1card < player2card)
		{
			wins = 2;
		}
	}

	if (player1suit < player2suit || wins == 1)
	{
		printf("Player 1 wins\n");
		return;
	}
	else if (player1suit > player2suit || wins == 2)
	{
		printf("Player 2 wins\n");
		return;
	}

	printf("There is a tie\n");
}