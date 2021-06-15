#include <stdio.h>
#include <math.h>

int main()
{
	double interval;
	double a, b;
	double result = 0;
	printf("Enter a value for a:");
	scanf("%lf", &a);
	printf("Enter a value for b:");
	scanf("%lf", &b);
	double old = 0;

	for (int subintervals = 10; subintervals <= 100000; subintervals++)
	{
		interval = (b - a) / (double)subintervals;
		if (interval < 1e-10)
			break;
		for (double pcs = 1; pcs <= subintervals; pcs++)
		{
			double x = a + (pcs - 0.5) * interval;
			if (x != 0)
				result += sin(x * M_PI) / (M_PI * x);
			else
				result += 1;
		}
		result *= interval;
		printf("%d: %.9lf %.9e\n", subintervals, result, result);
		if (old != 0 && fabs(fabs(old) - fabs(result)) < 1e-10)
		{
			break;
		}
		old = result;
		result = 0;
	}
	return 0;
}
