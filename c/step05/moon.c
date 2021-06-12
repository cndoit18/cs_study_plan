#include <stdio.h>
#include <math.h>
#include <stdbool.h>
#include <stdlib.h>

/* 
 * Simple lunar lander program.
 * By: cndoit18
 * Best landing: Time = 12 seconds, Fuel = 88.00, Velocity = -1.56
 */

int main()
{
	printf("Lunar Lander - (c) 2021, by cndoit18\n");

	double altitude = 100; /* Meters */
	double velocity = 0;   /* Meters per second */
	double fuel = 100;	   /* Kilograms */
	double power = 1.5;	   /* Acceleration per pound of fuel */
	double g = -1.63;	   /* Moon gravity in m/s^2 */
	double burn;		   /* Amount of fuel to burn */
	bool valid;			   /* Valid data entry flag */
	int seconds = 0;	   /* Time spent */
	printf("Altitude: %f\n", altitude);
	printf("Velocity: %f\n", velocity);
	printf("You have %f kilograms of fuel\n", fuel);

	while (altitude > 0)
	{
		printf("Altitude: %.2f\n", altitude);
		printf("Velocity: %.2f\n", velocity);
		printf("You have %.1f kilograms of fuel\n", fuel);
		do
		{
			valid = false;
			printf("How much fuel would you like to burn: ");
			scanf("%lf", &burn);

			if (burn < 0)
			{
				printf("You can't burn negative fuel\n");
			}
			else if (burn > fuel)
			{
				printf("You can't burn fuel you don't have\n");
			}
			else
			{
				valid = true;
				printf("Burning %.1f kilograms of fuel\n", burn);
			}
		} while (!valid);
		seconds++;
		velocity = velocity + g + power * burn;
		altitude += velocity;
		fuel -= burn;
	}

	printf("You landed with a velocity of %.2f\n", velocity);
	if (fabs(velocity) > 3)
	{
		printf("Your next of kin have been notified\n");
		exit(0);
	}
	printf("landing: Time = %d seconds, Fuel = %.2f, Velocity = %+.2f\n", seconds, fuel, velocity);

	return 0;
}