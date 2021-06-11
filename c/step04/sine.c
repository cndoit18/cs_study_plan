#include <stdio.h>
#include <math.h>

int main()
{
	double angle, sinVal, cosVal;
	int numSteps = 40;
	double maxAngle = M_PI * 2;
	int i, wide;
	char plot = '*';
	do {
		printf("Input the number of steps: ");
		scanf("%d", &numSteps);
	} while(numSteps < 2);

	for (i = 0; i <= numSteps; i++)
	{
		angle = (double)i / (double)numSteps * maxAngle;
		sinVal = sin(angle);
		cosVal = cos(angle);
		wide =  (int)(60 * (sinVal + 1));
		if (fabs(cosVal) < 0.1) {
			plot = '*';
		} else if (cosVal > 0) {
			plot = '\\';
		} else {
			plot = '/';
		}
		printf("%3d: %5.2f %*c\n", i, angle, wide, plot);
	}

	return 0;
}