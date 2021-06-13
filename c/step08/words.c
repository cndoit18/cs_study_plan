#include <stdio.h>
#include <string.h>

/* 
 * Name : cndoit18
 * Program to experiment with character arrays
 */

#define MaxWord 20

int main()
{
	char c;
	char str[MaxWord + 1];
	char max[MaxWord + 1] = {0};
	int len = 0;
	double sum = 0, words = 0;

	puts("Enter text. Include a dot ('.') to end a sentence to exit:");
	do
	{
		c = getchar();
		if (c != '.' && c != ' ')
		{
			if (len < MaxWord)
			{
				str[len++] = c;
			}
		}
		else
		{
			if (len == 0)
			{
				continue;
			}
			str[len] = 0;
			if (len > strlen(max))
			{
				for (int i = 0; i < len; i++)
				{
					max[i] = str[i];
				}
				max[len] = 0;
			}
			sum += len;
			words++;
			printf("%s\n", str);
			len = 0;
		}
	} while (c != '.');
	printf("%.2lf, %s\n", sum / words, max);
}