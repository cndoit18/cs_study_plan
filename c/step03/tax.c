#include <stdio.h>
#include <stdlib.h>

int main()
{
	double income;
	int select, dependents = 0, children = 0, deduction = 0;
	double tax = 0;
	printf("Enter your annual income: ");
	scanf("%lf", &income);
	if (income < 9350)
	{
		printf("0\n");
		exit(0);
	}
	printf("What is your filing status? \n1) single\n2) married filing jointly\n3) married filing separately\nPlease enter a number: ");
	scanf("%d", &select);
	switch (select)
	{
	case 1:
	case 3:
		dependents = 1;
		deduction = 5700;
		break;
	case 2:
		deduction = 11400;
		dependents = 2;
		break;
	default:
		printf("The option you entered is incorrect\n");
		exit(1);
	}
	if (dependents > 1)
	{
		printf("How many children do you have? ");
		scanf("%d", &children);
	}
	deduction += (children + dependents) * 3650;
	income -= deduction;
	if (income >= 137300)
	{
		printf("%.2lf\n", (income - 137300) * 0.28 + 26687.5);
	}
	else if (income >= 68000)
	{
		printf("%.2lf\n", (income - 68000) * 0.25 + 9362.5);
	}
	else if (income >= 16750)
	{
		printf("%.2lf\n", (income - 16750) * 0.15 + 1675);
	}
	else if (income >= 0)
	{
		printf("%.2lf\n", income * 0.1);
	}
	return 0;
}