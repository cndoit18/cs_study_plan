#include <stdio.h>
#include <math.h>

int main()
{
	double capacitance, inductance, frequency;
	printf("Input Capacitance (microfarads): ");
	scanf("%lf", &capacitance);
	printf("Input Inductance (millihenrys): ");
	scanf("%lf", &inductance);
	frequency = (1.0 / sqrt(capacitance / 1000000000 * inductance)) / (2 * M_PI);
	printf("Resonant Frequency is %.3lf\n", frequency);
	return 0;
}